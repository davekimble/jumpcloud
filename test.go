package main

import (
    "fmt"
    "crypto/sha512"
    "io"
    "encoding/base64"
)

func main() {
    var pw = "angryMonkey"
    hash := sha512.New()
    io.WriteString( hash, pw )
    fmt.Println(base64.URLEncoding.EncodeToString(hash.Sum(nil)))
}