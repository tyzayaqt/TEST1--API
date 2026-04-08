package main

import (
	"database/sql"
	"net/http"
)

// This file is responsible for defining the routes and linking them to their respective handlers.
func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/users", usersRouter(db))
	http.HandleFunc("/projects", projectsRouter(db))
	http.HandleFunc("/tasks", tasksRouter(db))
}
