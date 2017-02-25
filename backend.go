package main

import (
    "fmt"
    "log"
    "net/http"
)
var str="hello"
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        fmt.Fprintln(w, "Hello, world!")
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.Method, r.URL.Path)
        fmt.Fprintln(w, str)
    })

    log.Println("Listening on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalln(err)
    }
}