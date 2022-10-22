package middlewares

import (
	"context"
	"fmt"
	"net"
	"net/http"
	c "user_balance/service/internal/constants"
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
		host, port, _ := net.SplitHostPort(r.Host)
		scheme := "https"
		if r.TLS == nil {
			scheme = "http"
		}

		ctx := context.WithValue(r.Context(), c.BASE_PATH, scheme+"://"+host+":"+port)

		r = r.WithContext(ctx)
		t(w, r)
	}
}
