package req

import (
	"net/http"
	"web-app/pkg/res"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JsonAnswer(w, 402, err.Error())
		return nil, err
	}
	err = IsValid[T](body)
	if err != nil {
		res.JsonAnswer(w, 402, err.Error())
		return nil, err
	}
	return &body, nil
}
