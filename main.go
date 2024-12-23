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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	url := "https://booky-b5bu.onrender.com/api/book"

	go sendGetRequest(url)

    http.HandleFunc("/", helloHandler)
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
