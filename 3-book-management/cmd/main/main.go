package main

import (
	"book-management/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/junzhu/dialectes/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
