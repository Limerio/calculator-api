package middlewares

import "net/http"

func Headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "The Content-Type header needs to be set to \"application/json\"", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		next.ServeHTTP(w, r)
	})
}
