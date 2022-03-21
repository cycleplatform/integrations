package types

import (
	locationObj "github.com/cycleplatform/integrations/types/objects/location"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// input to POST /v1/infrastructure/ip/assign
	AssignIPRequest struct {
		Kind     IPKind                `json:"kind"`
		ServerId string                `json:"server_id"`
		Location *locationObj.Location `json:"location"`
		Cycle    CycleIPMeta           `json:"cycle"`
	}

	// output from POST /v1/infrastructure/ip/assign
	AssignIPResponse struct {
		IpId           string `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentID string `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		CIDR           string `bson:"cidr" json:"cidr"`
		Gateway        string `bson:"gateway" json:"gateway"`
		Netmask        string `bson:"netmask" json:"netmask"`
		Network        string `bson:"network" json:"network"`
	}

	// input to POST /v1/infrastructure/ip/release
	ReleaseIPRequest struct {
		Kind           IPKind                `json:"kind"`
		IpId           string                `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentID string                `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		ServerId       string                `json:"server_id"`
		Location       *locationObj.Location `json:"location"`
		Cycle          CycleIPMeta           `json:"cycle"`
	}

	CycleIPMeta struct {
		PoolID     *primitive.ObjectID `json:"pool_id"`
		ServerID   *primitive.ObjectID `json:"server_id"`
		LocationID *primitive.ObjectID `json:"location_id"`
	}

	IPKind string
)

const (
	IP_V6 IPKind = "v6"
	IP_V4 IPKind = "v4"
)
