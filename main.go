package main

import (
    "fmt"
    "log"
    "http/net"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Hello! This is my first Golang Server!\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
