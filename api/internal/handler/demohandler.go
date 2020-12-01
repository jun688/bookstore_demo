package handler

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

func MiddlewareDemoFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("request ...")
		next(w, r)
		logx.Info("response ... ")
	}
}
