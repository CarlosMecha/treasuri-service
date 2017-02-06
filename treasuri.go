package main

import (
    "fmt"
    "net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, Caca %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", HandleRequest)
    http.ListenAndServe(":8080", nil)
}
