package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	// called when a request is made to a new region for the first time
	// input to POST /v1/location/configure
	ConfigureLocationRequest struct {
		LocationId     string `json:"location_id"`
		CurrentVersion string `json:"version"`
	}

	// output from POST /v1/location/configure
	ConfigureLocationResponse struct {
		Configured    bool           `json:"configured"`
		Version       *string        `json:"version"`
		Configuration *Configuration `json:"configuration"`
	}

	Configuration map[string]interface{}

	CycleLocationMeta struct {
		LocationID primitive.ObjectID `json:"location_id"` // the id of the location @ cycle
	}
)
