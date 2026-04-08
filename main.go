package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	DSN := "postgres://tyzaya:password@localhost:5432/taskmanager?sslmode=disable"

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal("Could not open database:", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Could not connect to database (timed out):", err)
	}

	fmt.Println("Database connected successfully!")
	fmt.Println("Server running on http://localhost:4000")

	// Call the function from routes.go
	RegisterRoutes(db)

	// Start the server
	log.Fatal(http.ListenAndServe(":4000", nil))
}
