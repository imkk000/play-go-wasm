package main

import (
	"net/http"
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ignore asset not found warning in browser console
		if r.URL.Path == "/favicon.ico" ||
			r.URL.Path == "/css/style.css" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	handlers := middleware(http.FileServer(http.Dir("./public")))
	http.ListenAndServe("127.0.0.1:8080", handlers)
}
