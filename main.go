package main

import (
    "net/http"
    "io"
    "fmt"
)

func main() {
    res, _ := http.Get("https://ifconfig.me/ip")
    defer res.Body.Close()
    body, _ := io.ReadAll(res.Body)
    fmt.Println(string(body))
}
