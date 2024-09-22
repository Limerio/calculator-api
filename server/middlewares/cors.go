package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedMethods: []string{"POST"},
	}).Handler(h)
}
