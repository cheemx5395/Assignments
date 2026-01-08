package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/cheemx5395/assignments/backendfundamentals/go-crud/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const port = "8080"

type cfg struct {
	DB *database.Queries
}

func main() {

	// Loading db url for connection from .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment file: %v\n", err)
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening postgres: %v", err)
	}

	dbQueries := database.New(db)
	dbcfg := &cfg{
		DB: dbQueries,
	}

	mux := &http.ServeMux{}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Health Checkup handler
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		res := struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		}

		body, _ := json.Marshal(res)
		w.WriteHeader(200)
		w.Write(body)
	})

	// Create User handler
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		req := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		res := struct {
			Message string `json:"message"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("Error in creating the user: %v", err)
			return
		}
		if req.Email == "" || req.Name == "" {
			res.Message = "Both fields are supposed to be filled."
			body, _ := json.Marshal(res)
			w.WriteHeader(400)
			w.Write(body)
			return
		}

		_, err = dbcfg.DB.CreateUser(context.Background(), database.CreateUserParams{
			Email: req.Email,
			Name:  req.Name,
		})
		if err != nil {
			log.Printf("Error creating user: %v", err)
			res.Message = "Error creating user."
			body, _ := json.Marshal(res)
			w.WriteHeader(500)
			w.Write(body)
			return
		}

		res.Message = "User created successfully."
		body, _ := json.Marshal(res)
		w.WriteHeader(201)
		w.Write(body)
	})

	// List Users handler
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := dbcfg.DB.GetUsers(context.Background())

		res := struct {
			Message string `json:"message"`
		}{}

		if err != nil {
			log.Printf("Error retrieving users: %v", err)
			res.Message = "Error retrieving users."
			body, _ := json.Marshal(res)
			w.WriteHeader(500)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(users)
		w.WriteHeader(200)
		w.Write(body)
	})

	// Get user based on path parameter
	mux.HandleFunc("GET /users/{email}", func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")

		res := struct {
			Message string `json:"message"`
		}{}

		if email == "" {
			log.Printf("Email not specified in path parameter: %v", err)
			res.Message = "Email not specified in path parameter."
			body, _ := json.Marshal(res)
			w.WriteHeader(400)
			w.Write(body)
			return
		}

		user, err := dbcfg.DB.GetUser(context.Background(), email)
		if err != nil {
			log.Printf("Error retrieving user: %v", err)
			res.Message = "Error retrieving user."
			body, _ := json.Marshal(res)
			w.WriteHeader(500)
			w.Write(body)
			return
		}

		body, _ := json.Marshal(user)
		w.WriteHeader(200)
		w.Write(body)
	})

	// Delete user based on path parameter
	// Get user based on path parameter
	mux.HandleFunc("DELETE /users/{email}", func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")

		res := struct {
			Message string `json:"message"`
		}{}

		if email == "" {
			log.Printf("Email not specified in path parameter: %v", err)
			res.Message = "Email not specified in path parameter."
			body, _ := json.Marshal(res)
			w.WriteHeader(400)
			w.Write(body)
			return
		}

		err := dbcfg.DB.DeleteUser(context.Background(), email)
		if err != nil {
			log.Printf("Error retrieving user: %v", err)
			res.Message = "Error retrieving user."
			body, _ := json.Marshal(res)
			w.WriteHeader(500)
			w.Write(body)
			return
		}

		res.Message = "User deleted successfully."
		body, _ := json.Marshal(res)
		w.WriteHeader(200)
		w.Write(body)
	})

	log.Printf("Server running on: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
