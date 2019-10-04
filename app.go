// Entrypoint for API
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"gitlab.com/euthinkia/api-ayuntamientos/store"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "Port which will be operative for the API")
	flag.Parse()
	if port == nil {
		log.Fatal("--port must be set")
	}

	router := store.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	p := fmt.Sprintf(":%d", *port)
	fmt.Println("Listening on port:", *port)
	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(p, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
