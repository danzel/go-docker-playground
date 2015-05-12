package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Printf("start.\n")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/products/{id}", ProductsHandler)

	server := &http.Server{
		Addr:    ":35235",
		Handler: r,
	}
	server.ListenAndServe()
	//http.Handle("/", r)
	fmt.Printf("end.\n")
}

func HomeHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, "Home request handled. <a href=\"/products/123\">Link</a>\n")
}

func ProductsHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	args := mux.Vars(request)

	fmt.Fprintf(w, "Products request handled %v.\n", args["id"])
}
