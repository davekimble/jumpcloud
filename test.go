package main

import (
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/base64"
    "net/http"
    "time"
    "flag"
)

var debug bool
var pause int

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
    }

    password := r.URL.Query().Get("password")
    if len(password) == 0 {
        fmt.Println("ERROR - No password specified!")
	return
    }

    if debug {
        fmt.Printf("password = %s\n", password)
        fmt.Printf("pausing %d seconds...\n", pause)
    }

    time.Sleep(time.Duration(pause) * time.Second)

    result := getHash(password)
    fmt.Printf("result = %s\n", result)

    io.WriteString(w, result)
}

func main() {
    fmt.Println("Starting server...")
    debugPtr := flag.Bool("debug", false, "a bool")
    portPtr := flag.Int("port", 8080, "an int")
    pausePtr := flag.Int("pause", 5, "an int")
    flag.Parse()

    portNumberString := fmt.Sprintf(":%d",*portPtr)
    pause = *pausePtr
    debug = *debugPtr

    if debug {
        fmt.Println("portNumberString:", portNumberString)
        fmt.Println("pause:", pause)
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(portNumberString, nil)
}