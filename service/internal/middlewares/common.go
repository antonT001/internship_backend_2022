package middlewares

import (
	"fmt"
	"net/http"
)

type CommonMiddleware interface {
	Handle(next http.Handler) http.Handler
}

type commonMiddleware struct {
}

func NewCommonMiddleware() CommonMiddleware {
	return &commonMiddleware{}
}

func (m commonMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(m.a(next.ServeHTTP))
}

func (m commonMiddleware) a(t func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover panic")
				fmt.Printf("r: %v\n", r)
			}
		}()
		t(w, r)
	}
}
