package auth

import (
	"fmt"
	"net/http"
	"web-app/pkg/req"
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
		result, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(result)
		data := LoginResponse{
			Token: "123",
		}
		res.JsonAnswer(w, 400, data)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		fmt.Println(result)
		data := RegisterResponse{
			Token: "123",
		}
		res.JsonAnswer(w, 400, data)
	}
}
