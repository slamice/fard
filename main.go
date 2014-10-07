package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"
    "html/template"
)

var (
    port = flag.String("p", "8000", "Port number (default 8000)")
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.New("index.html").ParseFiles("templates/index.html", "templates/base.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // Render the template
    err = t.ExecuteTemplate(w, "base", map[string]interface{}{"Page": "home"})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/", HomeHandler)
    fmt.Println("Running on localhost:" + *port)

    log.Fatal(http.ListenAndServe(":"+*port, nil))
}