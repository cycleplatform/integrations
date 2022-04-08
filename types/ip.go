package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// input to POST /v1/infrastructure/ip/allocate
	AllocateIPRequest struct {
		Kind                  IPKind        `json:"kind"`
		ServerId              string        `json:"server_id"`
		LocationId            string        `json:"location_id"`
		LocationConfiguration Configuration `json:"location_configuration"`
		Cycle                 CycleIPMeta   `json:"cycle"`
	}

	// output from POST /v1/infrastructure/ip/allocate
	AllocateIPResponse struct {
		IpId           string  `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentId string  `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		CIDR           string  `bson:"cidr" json:"cidr"`
		Gateway        string  `bson:"gateway" json:"gateway"`
		Netmask        string  `bson:"netmask" json:"netmask"`
		Network        string  `bson:"network" json:"network"`
		NATPrivateIP   *string `bson:"nat_private_ip" json:"nat_private_ip"`
	}

	// input to POST /v1/infrastructure/ip/release
	ReleaseIPRequest struct {
		Kind                  IPKind        `json:"kind"`
		IpId                  string        `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentId        string        `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		CIDR                  string        `bson:"cidr" json:"cidr"`
		NATPrivateIP          *string       `bson:"nat_private_ip" json:"nat_private_ip"`
		ServerId              string        `json:"server_id"`
		LocationId            string        `json:"location_id"`
		LocationConfiguration Configuration `json:"location_configuration"`
		Cycle                 CycleIPMeta   `json:"cycle"`
	}

	CycleIPMeta struct {
		PoolID              *primitive.ObjectID `json:"pool_id"`
		ServerID            primitive.ObjectID  `json:"server_id"`
		ServerModelClass    string              `json:"server_model_class"`    // the class of the server @ cycle
		ServerModelCategory string              `json:"server_model_category"` // the category of the server at Cycle
		LocationID          primitive.ObjectID  `json:"location_id"`
	}

	IPKind string
)

const (
	IP_V6 IPKind = "ipv6"
	IP_V4 IPKind = "ipv4"
)
