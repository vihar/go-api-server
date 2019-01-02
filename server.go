package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Used to add routes to the application.
    router := mux.NewRouter()

    // Report and Start Server on Localhost 8000.
    fmt.Println("Running on Port 8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}