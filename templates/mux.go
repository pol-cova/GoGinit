package templates

const MuxTemplate = `package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/about", AboutHandler).Methods("GET")
    
    // Start the server
    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// HomeHandler handles requests to the root URL
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to Mux!"))
}

// AboutHandler handles requests to the /about URL
func AboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About Page"))
}
`
