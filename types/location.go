package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	// called when a request is made to a new region for the first time
	// input to POST /v1/location/configure
	ConfigureLocationRequest struct {
		LocationId           string            `json:"location_id"`
		CurrentVersion       string            `json:"version"`
		CurrentConfiguration Configuration     `json:"configuration_current"`
		Cycle                CycleLocationMeta `json:"cycle"`
	}

	// output from POST /v1/location/configure
	ConfigureLocationResponse struct {
		Configured          bool          `json:"configured"`
		Version             *string       `json:"version"`
		LatestConfiguration Configuration `json:"configuration_latest"`
	}

	Configuration map[string]interface{}

	CycleLocationMeta struct {
		LocationID primitive.ObjectID `json:"location_id"` // the id of the location @ cycle
	}
)
