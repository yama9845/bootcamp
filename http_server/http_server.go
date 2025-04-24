package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	/*
	   w: クライアントにレスポンスを返すためのもの（http.ResponseWriter）
	   r: クライアントから送られてきたリクエスト情報（*http.Request）
	   fmt.Fprint: 指定した出力先（ここでは w）に文字列を書き込む
	*/
	fmt.Fprint(w, "Hello World!")

	// リクエストのHTTPメソッド（GET, POST）を取得
	log.Println("Method:", r.Method)
	// リクエストされたURLのパスを取得
	log.Println("Path:", r.URL.Path)
	// ヘッダー全体を取得する
	log.Println("Content-Type:", r.Header)
	// 特定のヘッダーを取得する
	log.Println("Content-Type:", r.Header.Get("Content-type"))
}

func greet(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("name")

	if queryParam == "" {
		fmt.Fprint(w, "Hello, Guest!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", queryParam)
}

type jsonData struct {
	Message string `json:"message"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var jsdata jsonData

	err = json.Unmarshal(requestBody, &jsdata)
	jsdata.Message = "bug message"
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.Marshal(jsdata)
	if err != nil {
		http.Error(w, "422 Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/echo", echo)

	server.ListenAndServe()
}
