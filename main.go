package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %q\n", err)
	}
}

// Build with >> go build -o bin/sci-version-proxy -v .
func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World\n")
	})

	log.Printf("Golang App running...\n")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
