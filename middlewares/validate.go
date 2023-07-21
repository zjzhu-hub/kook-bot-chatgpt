package middlewares

import (
	"kook-bot-chatgpt/utils"
	"net/http"
	"strings"
)

// 校验请求
func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 必须是POST请求
		if r.Method != "POST" {
			utils.ErrorLogger(w, "Method not allowed")
			return
		}

		// 必须是JOSN
		contentType := r.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			utils.ErrorLogger(w, "Invalid Content-Type. Expected 'application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}
