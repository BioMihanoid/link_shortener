package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"web-app/pkg/res"
)

type AuthHandler struct {
}

func NewAuthHandler(router *http.ServeMux) {
	handler := &AuthHandler{}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//TODO: do check regexp email
		match, _ := regexp.MatchString("^[a-zA-Z0-9_-]{5,15}$", req.Email)
		if match {
			http.Error(w, "Email bad", http.StatusForbidden)
			return
		}
		if req.Email == "" || req.Password == "" {
			http.Error(w, "Email and password fields are required", http.StatusBadRequest)
			return
		}
		data := LoginResponse{
			Token: "123",
		}
		res.JsonAnswer(w, 200, data)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
