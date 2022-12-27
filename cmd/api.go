// ----------------------------------------------------------------------------
// Copyright (c) Ben Coleman, 2020
// Licensed under the MIT License.
//
// Define a simple 'Things' API for the service
// ----------------------------------------------------------------------------

package main

import (
	"github.com/benc-uk/go-rest-api/pkg/api"
	"github.com/go-chi/chi/v5"
)

// ThingAPI is a wrap of the common base API with local implementation
type ThingAPI struct {
	*api.Base
	// Add extra fields here: database connections, SDK clients
}

func (api ThingAPI) addRoutes(r chi.Router) {
	r.Get("/things", api.getThings)
	r.Get("/things/{id}", api.getThingByID)
	r.Post("/things", api.createThing)
}

func NewThingAPI() ThingAPI {
	return ThingAPI{
		api.NewBase(serviceName, version, buildInfo, healthy),
		// Database connections, SDK clients, etc can be added here
	}
}
