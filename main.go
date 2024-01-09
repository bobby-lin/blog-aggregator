package main

import (
	"context"
	"database/sql"
	"github.com/bobby-lin/blog-aggregator/internal/database"
	"github.com/bobby-lin/blog-aggregator/internal/utils"
	_ "github.com/lib/pq"
	"strings"
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
	r.Get("/users", cfg.middlewareAuth(cfg.GetUserHandler))

	r.Post("/feeds", cfg.middlewareAuth(cfg.CreateFeedHandler))
	r.Get("/feeds", cfg.GetFeedHandler)

	r.Post("/feed_follows", cfg.middlewareAuth(cfg.CreateFeedFollowerHandler))
	r.Get("/feed_follows", cfg.middlewareAuth(cfg.GetFeedFollowersHandler))
	r.Delete("/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.DeleteFeedFollowsHandler))
	return r
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authenticates request
		apiKey := strings.Replace(r.Header.Get("Authorization"), "ApiKey ", "", 1)
		ctx := context.Background()
		user, err := cfg.DB.SelectUser(ctx, apiKey)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "invalid api key")
			return
		}

		// Call the next handler
		handler(w, r, user)
	}
}
