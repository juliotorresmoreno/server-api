package app

import (
	"github.com/juliotorresmoreno/server-api/schemas"
	mem "github.com/rs/rest-layer-mem"
	"github.com/rs/rest-layer/resource"
)

func bind() resource.Index {
	index := resource.NewIndex()

	// Add a resource on /users[/:user_id]
	index.Bind("users", schemas.User, mem.NewHandler(), resource.Conf{
		// We allow all REST methods
		// (rest.ReadWrite is a shortcut for []resource.Mode{resource.Create, resource.Read, resource.Update, resource.Delete, resource,List})
		AllowedModes: resource.ReadWrite,
	})
	return index
}
