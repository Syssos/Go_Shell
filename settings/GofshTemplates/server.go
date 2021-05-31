package main

import (
    "fmt"
    "log"
    // "os"
    "io/ioutil"
    "strings"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            errorHandler(w, r, http.StatusNotFound)
            return
        }
        http.ServeFile(w, r, "./templates/index.html")
    })

    fmt.Println("Starting server on localhost:3000")
    log.Fatal(http.ListenAndServe(":3000", nil))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    fmt.Println(r.URL)
    dat, err := ioutil.ReadFile("./templates/404.html")
    checkErr(err)

    pageData := strings.Replace(string(dat), "{{url}}", fmt.Sprintf("%v", r.URL), -1)
    if status == http.StatusNotFound {
        fmt.Fprint(w, pageData)
    }
}

func checkErr(e error) {
    if e != nil {
        fmt.Println(e)
    }
}