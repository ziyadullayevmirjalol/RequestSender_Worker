package main

import (
    "fmt"
    "net/http"
    "time"
)

func sendGetRequest(url string) {
    for {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Printf("An error occurred: %v\n", err)
        } else {
            fmt.Printf("Response status code: %d\n", resp.StatusCode)
            resp.Body.Close()
        }
        time.Sleep(1 * time.Second)
    }
}

func main() {
    // example request endpoint
    url := "https://booky-b5bu.onrender.com/api/book"
    sendGetRequest(url)
}
