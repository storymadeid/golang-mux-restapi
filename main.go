package main

import (
	"log"
	"net/http"

	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/controllers/product"
	"github.com/Sepctrevuln-Sketch/golang-restapi-mux/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabases()
	router := mux.NewRouter()

	router.HandleFunc("/products", product.Index).Methods("GET")
	router.HandleFunc("/products/{id}", product.Show).Methods("GET")
	router.HandleFunc("/products", product.Create).Methods("POST")
	router.HandleFunc("/products/update/{id}", product.Update).Methods("PUT")
	router.HandleFunc("/products/delete", product.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5544", router))
}
