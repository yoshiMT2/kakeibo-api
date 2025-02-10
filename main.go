package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env vars")
	}

	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting database: %s", err.Error())
	}
	defer db.Close()
	fmt.Println("connected to db...")

	server := NewAPIServer(":3000", db)
	server.Run()
}
