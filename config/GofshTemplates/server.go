package main

import (
    "log"
    "fmt"
    "net/http"
    "html/template"
)

type submitResponse struct {
    Name      string
    DataInput bool
}

// Starts web application generated for pond
func main() {
    // Setting static folder in route, then stripping it from url
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Setting functions to handle routes users may go to
    http.HandleFunc("/", HomePage)
    http.HandleFunc("/form", SubmitPage)

    // Alerting console servers starting
    fmt.Println("Starting server on localhost:3000")

    // Starting server & logging errors to console
    log.Fatal(http.ListenAndServe(":3000", nil))
}

// "/" route handler 
func HomePage (w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    http.ServeFile(w, r, "./templates/index.html")
}

// "/form" route handler 
func SubmitPage (w http.ResponseWriter, r *http.Request) {
    
    if r.URL.Path != "/form" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    resp := submitResponse{"", false}
    tmpl := template.Must(template.ParseFiles("templates/form.html"))

    // handles post request, if method is not "POST" generic template is used
    if r.Method == "POST" {
        resp.Name      = r.FormValue("name")
        resp.DataInput = true

        // Do more stuff with response here


        tmpl.Execute(w, resp)
    }
       
    tmpl.Execute(w, resp)
}

// Handles 404 error page
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    
    http.ServeFile(w, r, "./templates/404.html")
}

// Checks for errors throughout code to save space
func checkErr(e error) {
    if e != nil {
        fmt.Println(e)
    }
}