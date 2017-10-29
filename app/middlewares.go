package app

import (
	"github.com/juliotorresmoreno/server-api/routes"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func middlewares() alice.Chain {
	c := alice.New()
	c = c.Append(routes.NewRouter)

	// Add CORS support with passthrough option on so rest-layer can still
	// handle OPTIONS method
	c = c.Append(cors.New(cors.Options{OptionsPassthrough: true}).Handler)
	return c
}
