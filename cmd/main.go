package main

import (
	"fmt"
	"net/http"
	"web-app/config"
	"web-app/internal/auth"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	_ = config.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running and listing on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
