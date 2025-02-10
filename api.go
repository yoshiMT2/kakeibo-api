package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIServer struct {
	listenAddr string
	db         *sql.DB
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/item", makeHttpHandleFunc(s.handleItem))

	log.Println("Server started running on port %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleItem(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetItems(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateItem(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteItem(w, r)
	}
	return WriteJson(w, http.StatusMethodNotAllowed, ApiError{Error: "method not allowed"})
}

func (s *APIServer) handleGetItems(w http.ResponseWriter, r *http.Request) error {
	items := []Item{}
	userId := r.URL.Query().Get("user-id")
	rows, err := s.db.Query("SELECT id, name, created_at FROM items WHERE user_id = $1", userId)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, ApiError{Error: "bad request"})
	}
	for rows.Next() {
		var i Item
		err = rows.Scan(&i.ID, &i.Name, &i.CreatedAt)
		if err != nil {
			return WriteJson(w, http.StatusInternalServerError, ApiError{Error: "internal server error"})
		}
		items = append(items, i)
		fmt.Println(i)
	}
	fmt.Println(items)
	return WriteJson(w, http.StatusOK, items)
}

func (s *APIServer) handleCreateItem(w http.ResponseWriter, r *http.Request) error {
	var i *ItemForm
	err := json.NewDecoder(r.Body).Decode(&i)

	// body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return WriteJson(w, http.StatusBadRequest, ApiError{Error: "BR"})
	}
	result, dbErr := s.db.Exec("INSERT INTO items (user_id, name) VALUES ($1, $2)", i.UserID, i.Name)
	if dbErr != nil {
		return fmt.Errorf("db error: %v", dbErr)
	}
	fmt.Println(result)
	return nil
}

func (s *APIServer) handleDeleteItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}
