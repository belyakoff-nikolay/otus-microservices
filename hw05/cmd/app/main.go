package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	addr := fmt.Sprintf(":%v", port)
	fmt.Println("addr", addr)

	healthCheck := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"status": "OK"}`))
	}

	mux := &http.ServeMux{}
	mux.HandleFunc("/health/", healthCheck)
	mux.HandleFunc("/health", healthCheck)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("homework 5"))
	})

	server := &http.Server{Addr: addr, Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
