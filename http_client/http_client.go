package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// GET request
	resp, err := http.Get("http://localhost:8080/echo")
	if err != nil {
		fmt.Println("<Error> GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("<Error> GET responce read:", err)
		return
	}
	fmt.Println("GET responce:", string(body))

	// POST request
	postData := []byte(`{"message":"Hello, Server!"}`)
	resp, err = http.Post("http://localhost:8080/echo", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("<Error> POST request:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("<Error> GET responce read:", err)
		return
	}
	fmt.Println("POST responce:", string(body))
}
