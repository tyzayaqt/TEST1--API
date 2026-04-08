package main

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Project struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Simplified Request objects for decoding JSON
type CreateProjectRequest struct {
	OwnerID     int    `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Task struct {
	ID          int        `json:"id"`
	ProjectID   int        `json:"project_id"`
	CreatedBy   int        `json:"created_by"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date"` // Pointer allows this to be null
	CreatedAt   time.Time  `json:"created_at"`
}

// Simplified Request objects for decoding JSON
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateTaskRequest struct {
	ProjectID   int    `json:"project_id"`
	CreatedBy   int    `json:"created_by"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}
