package utils

import (
	"log"
	"net/http"
	"strings"
)

func ErrorLogger(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusBadRequest)
	log.Println(errMsg)
}

func ChainMiddleware(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}

func ParseCommand(content string) (string, string) {
	if !strings.HasPrefix(content, "/") {
		return "", ""
	}
	parts := strings.SplitN(content, " ", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
