package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println(port)
}
