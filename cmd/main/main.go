package main

import (
	"bookstore/pkg/config"
	"bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	config.ConnectDb()
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8080",r))
}