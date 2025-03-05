package main

import (
	"fmt"
	"net/http"
	"web-app/config"
	"web-app/internal/auth"
	"web-app/pkg/db"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	conf := config.LoadConfig()
	_ = db.NewDb(conf)
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
