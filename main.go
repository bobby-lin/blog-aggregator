package main

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

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"OPTION", "GET", "POST", "PUT", "DELETE"},
		AllowCredentials: false,
	}))

	r.Mount("/v1", v1Router())

	fmt.Println("The server is starting on port", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func v1Router() http.Handler {
	r := chi.NewRouter()
	return r
}
