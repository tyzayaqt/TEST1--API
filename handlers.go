package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

// Helper function to send error messages in JSON format
func sendError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// USERS
func usersRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			rows, _ := db.Query("SELECT id, name, email, created_at FROM users")
			defer rows.Close()
			var users []User = []User{}
			for rows.Next() {
				var u User
				rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
				users = append(users, u)
			}
			json.NewEncoder(w).Encode(users)
		} else if r.Method == "POST" {
			var req CreateUserRequest
			json.NewDecoder(r.Body).Decode(&req)
			var u User
			db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at",
				req.Name, req.Email).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
			json.NewEncoder(w).Encode(u)
		}
	}
}

// PROJECTS
func projectsRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			rows, _ := db.Query("SELECT id, owner_id, name, description, created_at FROM projects")
			defer rows.Close()
			var projects []Project = []Project{}
			for rows.Next() {
				var p Project
				rows.Scan(&p.ID, &p.OwnerID, &p.Name, &p.Description, &p.CreatedAt)
				projects = append(projects, p)
			}
			json.NewEncoder(w).Encode(projects)
		} else if r.Method == "POST" {
			var req CreateProjectRequest
			json.NewDecoder(r.Body).Decode(&req)
			var p Project
			db.QueryRow("INSERT INTO projects (owner_id, name, description) VALUES ($1, $2, $3) RETURNING id, owner_id, name, description, created_at",
				req.OwnerID, req.Name, req.Description).Scan(&p.ID, &p.OwnerID, &p.Name, &p.Description, &p.CreatedAt)
			json.NewEncoder(w).Encode(p)
		}
	}
}

func tasksRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//GET ALL TASKS
		if r.Method == "GET" {
			rows, _ := db.Query("SELECT id, project_id, created_by, title, description, status, due_date, created_at FROM tasks")
			defer rows.Close()
			var tasks []Task = []Task{}
			for rows.Next() {
				var t Task
				rows.Scan(&t.ID, &t.ProjectID, &t.CreatedBy, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt)
				tasks = append(tasks, t)
			}
			json.NewEncoder(w).Encode(tasks)

			//CREATE A TASK
		} else if r.Method == "POST" {
			var req CreateTaskRequest
			json.NewDecoder(r.Body).Decode(&req)
			var dueDate *time.Time
			if req.DueDate != "" {
				t, _ := time.Parse("2006-01-02", req.DueDate)
				dueDate = &t
			}
			var t Task
			db.QueryRow(`INSERT INTO tasks (project_id, created_by, title, description, due_date) 
				VALUES ($1, $2, $3, $4, $5) RETURNING id, project_id, created_by, title, description, status, due_date, created_at`,
				req.ProjectID, req.CreatedBy, req.Title, req.Description, dueDate).
				Scan(&t.ID, &t.ProjectID, &t.CreatedBy, &t.Title, &t.Description, &t.Status, &t.DueDate, &t.CreatedAt)
			json.NewEncoder(w).Encode(t)

			// 3. UPDATE TASK STATUS (e.g., /tasks?id=1)
		} else if r.Method == "PUT" {
			id := r.URL.Query().Get("id")
			var req struct {
				Status string `json:"status"`
			}
			json.NewDecoder(r.Body).Decode(&req)
			_, err := db.Exec("UPDATE tasks SET status = $1 WHERE id = $2", req.Status, id)
			if err != nil {
				sendError(w, http.StatusInternalServerError, "Update failed")
				return
			}
			w.WriteHeader(http.StatusNoContent)

			// 4. DELETE A TASK (e.g., /tasks?id=1)
		} else if r.Method == "DELETE" {
			id := r.URL.Query().Get("id")
			_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
			if err != nil {
				sendError(w, http.StatusInternalServerError, "Delete failed")
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
