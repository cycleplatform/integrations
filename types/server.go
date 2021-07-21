package types

import (
	locationObj "github.com/cycleplatform/integrations/types/objects/location"
	serverModelObj "github.com/cycleplatform/integrations/types/objects/server-model"
)

type (
	// input to POST /v1/infrastructure/server/create
	NewServerRequest struct {
		Hostname         string                      `json:"hostname"`
		Model            *serverModelObj.ServerModel `json:"model"`
		Location         *locationObj.Location       `json:"location"`
		ProvisionOptions map[string]interface{}      `json:"provision_options"` // any extra provision options that might be needed for certain providers
		Cycle            CycleServerMeta             `json:"cycle"`
	}

	// output from POST /v1/infrastructure/server/create
	NewServerResponse struct {
		ProviderId string `json:"provider_id"` // what is the ID of the server at the provider?
		Auth       struct {
			InitialIPs []string `json:"initial_ips"`
			MacAddrs   []string `json:"mac_addrs"`
		} `json:"auth"`
		Extra map[string]interface{} `json:"extra"`
	}

	// input to POST /v1/infrastructure/server/restart
	RestartServerRequest struct {
		ProviderId string          `json:"provider_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`    // send hostname in case provider requires it
		Cycle      CycleServerMeta `json:"cycle"`
	}

	// input to POST /v1/infrastructure/server/delete
	DeleteServerRequest struct {
		ProviderId string          `json:"provider_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`    // send hostname in case provider requires it
		Cycle      CycleServerMeta `json:"cycle"`
	}

	CycleServerMeta struct {
		ServerID   *string `json:"server_id"`   // the id of the server @ cycle
		ModelID    *string `json:"model_id"`    // the id of the model @ cycle
		LocationID *string `json:"location_id"` // the id of the location @ cycle
	}
)
