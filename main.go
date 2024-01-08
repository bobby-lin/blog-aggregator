package main

import (
	"database/sql"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	dbURL := os.Getenv("DB_STRING")
	db, err := sql.Open("postgres", dbURL)
	cfg := apiConfig{DB: database.New(db)}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"OPTION", "GET", "POST", "PUT", "DELETE"},
		AllowCredentials: false,
	}))

	r.Mount("/v1", cfg.v1Router())

	fmt.Println("The server is starting on port", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (cfg *apiConfig) v1Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/readiness", ReadinessHandler)
	r.Get("/err", ErrorHandler)
	r.Post("/users", cfg.CreateUserHandler)
	r.Get("/users", cfg.GetUserHandler)
	r.Post("/feeds", cfg.CreateFeedHandler)
	return r
}
