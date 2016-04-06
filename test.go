package main

import (
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/base64"
    "net/http"
)

const debug bool = true

func getHash(password string) (string) {
    if debug {
        fmt.Println("in getHash()")
    }
    hash := sha512.New()
    io.WriteString( hash, password )
    return base64.URLEncoding.EncodeToString(hash.Sum(nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
    if debug {
        fmt.Println("in handler()");
    }
    result := getHash("angryMonkey")
    fmt.Println(result)
}

func main() {
    fmt.Println("Starting server...")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}