package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found, relying on environment variables")
	// }

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting database: %s", err.Error())
	}
	defer db.Close()
	fmt.Println("connected to db...")

	server := NewAPIServer(":3000", db)
	server.Run()
}
