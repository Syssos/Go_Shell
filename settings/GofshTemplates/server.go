package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/index.html")
    })

    fmt.Println("Starting server on localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}