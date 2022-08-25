package http

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              true,
	})

	return c
}
