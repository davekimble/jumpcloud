package main

import (
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/base64"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    var pw = "angryMonkey"
    hash := sha512.New()
    io.WriteString( hash, pw )
    fmt.Println(base64.URLEncoding.EncodeToString(hash.Sum(nil)))

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}