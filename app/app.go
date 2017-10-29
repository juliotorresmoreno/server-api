// +build go1.7

package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/rs/rest-layer/rest"
	"github.com/rs/zerolog/log"
)

func NewServer() *http.Server {
	server := mux.NewRouter().StrictSlash(true)
	// Create a REST API resource index
	index := bind()

	// Create API HTTP handler for the resource graph
	api, err := rest.NewHandler(index)
	if err != nil {
		log.Fatal().Msgf("Invalid API configuration: %s", err)
	}
	c := middlewares()

	router := c.Then(http.StripPrefix("/api/", api))
	server.PathPrefix("/api").Handler(router)

	var addr = ":8080"
	return &http.Server{
		Addr:           addr,
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
