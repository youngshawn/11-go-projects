package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/youngshawn/11-go-projects/go-bookstore/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	//http.Handle("/", r)
	fmt.Println("Starting Server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
