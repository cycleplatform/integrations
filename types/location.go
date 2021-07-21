package types

import (
	locationObj "github.com/cycleplatform/integrations/types/objects/location"
)

type (
	ListLocationsResponse struct {
		Locations []locationObj.Location `json:"locations"`
		Extra     map[string]interface{} `json:"extra"`
	}
)
