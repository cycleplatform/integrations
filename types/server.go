package types

import (
	serverModelObj "github.com/cycleplatform/integrations/types/objects/server-model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// input to POST /v1/infrastructure/server/provision
	ProvisionServerRequest struct {
		Hostname         string                  `json:"hostname"`
		ModelId          string                  `json:"model_id"`
		LocationId       string                  `json:"location_id"`
		ModelFeatures    serverModelObj.Features `json:"model_features"`    // the features associated with this model
		ProvisionOptions ProvisionOptions        `json:"provision_options"` // any extra provision options that might be needed for certain providers
		Cycle            CycleServerMeta         `json:"cycle"`
	}

	ProvisionOptions struct {
		AttachedStorageSize *int    `bson:"attached_storage_size,omitempty" json:"storage_size,omitempty"`
		ReservationId       *string `bson:"reservation_id,omitempty" json:"reservation_id,omitempty"`
	}

	// output from POST /v1/infrastructure/server/provision
	ProvisionServerResponse struct {
		ServerId string `json:"server_id"` // what is the ID of the server at the provider?
		Auth     struct {
			UUID       *string  `json:"uuid"`
			InitialIPs []string `json:"initial_ips"`
			MacAddr    *string  `json:"mac_addr"`
		} `json:"auth"`
		Extra map[string]interface{} `json:"extra"`
	}

	// input to POST /v1/infrastructure/server/restart
	RestartServerRequest struct {
		ServerId   string          `json:"server_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`  // send hostname in case provider requires it
		ModelId    string          `json:"model_id"`
		LocationId string          `json:"location_id"`
		Cycle      CycleServerMeta `json:"cycle"`
	}

	// input to POST /v1/infrastructure/server/decommission
	DecommissionServerRequest struct {
		ServerId   string          `json:"server_id"` // what is the ID of the server at the provider?
		Hostname   string          `json:"hostname"`  // send hostname in case provider requires it
		ModelId    string          `json:"model_id"`
		LocationId string          `json:"location_id"`
		Cycle      CycleServerMeta `json:"cycle"`
	}

	CycleServerMeta struct {
		ServerID      primitive.ObjectID `json:"server_id"`      // the id of the server @ cycle
		ModelClass    string             `json:"model_class"`    // the class of the server @ cycle
		ModelCategory string             `json:"model_category"` // the category of the server at Cycle
		ModelID       primitive.ObjectID `json:"model_id"`       // the id of the model @ cycle
		LocationID    primitive.ObjectID `json:"location_id"`    // the id of the location @ cycle
	}
)
