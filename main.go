package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env only on local (if exists)
	_ = godotenv.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mongoURI := os.Getenv("MONGO_DB_URI")
		fmt.Fprintf(w, "This is the Golang deployed server on Railway. MONGO_DB_URI: %s", mongoURI)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
