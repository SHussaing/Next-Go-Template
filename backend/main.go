
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	Handlers "backend/handlers"
)

func main() {
	http.HandleFunc("/hello", Handlers.HelloHandler)

	// Enable CORS for all origins
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Println("Starting backend server on http://localhost:8080")
	// Start the server with CORS enabled
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(http.DefaultServeMux)))
}
