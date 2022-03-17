package types

type (
	// input to POST /v1/infrastructure/server/create
	NewServerRequest struct {
		Hostname         string                 `json:"hostname"`
		ModelId          string                 `json:"model_id"`
		LocationID       string                 `json:"location_id"`
		ProvisionOptions map[string]interface{} `json:"provision_options"` // any extra provision options that might be needed for certain providers
		Cycle            CycleServerMeta        `json:"cycle"`
	}

	// output from POST /v1/infrastructure/server/create
	NewServerResponse struct {
		ServerId string `json:"server_id"` // what is the ID of the server at the provider?
		Auth     struct {
			UUID       *string  `json:"uuid"`
			InitialIPs []string `json:"initial_ips"`
			MacAddrs   []string `json:"mac_addrs"`
		} `json:"auth"`
		Extra map[string]interface{} `json:"extra"`
	}

	// input to POST /v1/infrastructure/server/restart
	RestartServerRequest struct {
		ServerId   string          `json:"server_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`  // send hostname in case provider requires it
		ModelId    string          `json:"model_id"`
		LocationID string          `json:"location_id"`
		Cycle      CycleServerMeta `json:"cycle"`
	}

	// input to POST /v1/infrastructure/server/delete
	DeleteServerRequest struct {
		ServerId   string          `json:"server_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`  // send hostname in case provider requires it
		ModelId    string          `json:"model_id"`
		LocationID string          `json:"location_id"`
		Cycle      CycleServerMeta `json:"cycle"`
	}

	CycleServerMeta struct {
		ServerID      *string `json:"server_id"`      // the id of the server @ cycle
		ModelClass    *string `json:"model_class"`    // the class of the server @ cycle
		ModelCategory *string `json:"model_category"` // the category of the server at Cycle
		ModelID       *string `json:"model_id"`       // the id of the model @ cycle
		LocationID    *string `json:"location_id"`    // the id of the location @ cycle
	}
)
