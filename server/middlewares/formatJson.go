package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Limerio/calculator-api/server/utils/constants"
)

func FormatJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var values []float64

		err := json.NewDecoder(r.Body).Decode(&values)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), constants.BODY_JSON, values)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
