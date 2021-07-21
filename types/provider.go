package types

import (
	locationObj "github.com/cycleplatform/integrations/types/objects/location"
	serverModelObj "github.com/cycleplatform/integrations/types/objects/server-model"
)

type (
	// output from GET /v1/infrastructure/provider/locations
	ListLocationsResponse struct {
		Locations []locationObj.Location `json:"locations"`
		Extra     map[string]interface{} `json:"extra"`
	}

	// output from GET /v1/infrastructure/provider/server-models
	ListServerModelsResponse struct {
		Models []serverModelObj.ServerModel `json:"models"`
		Extra  map[string]interface{}       `json:"extra"`
	}
)
