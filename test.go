package main

import (
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/base64"
    "net/http"
    "time"
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
        fmt.Println("in handler()")
        fmt.Println("pausing 5 seconds...")
    }

    password := r.URL.Query().Get("password")
    if len(password) == 0 {
        fmt.Println("ERROR - No password specified!")
	return
    }

    io.WriteString(w, password)

    if debug {
        fmt.Printf("password = %s\n", password)
    }

    time.Sleep(time.Duration(5) * time.Second)

    result := getHash(password)
    fmt.Printf("result = %s\n", result)
}

func main() {
    fmt.Println("Starting server...")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}