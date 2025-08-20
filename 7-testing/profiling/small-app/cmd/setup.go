package main

import (
	"net/http"
	"small-app/handlers"
)

func setupRoutes(s *http.ServeMux) {
	s.HandleFunc("/find", handlers.FindUser)
}
