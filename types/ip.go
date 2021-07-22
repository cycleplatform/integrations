package types

import (
	locationObj "github.com/cycleplatform/integrations/types/objects/location"
)

type (
	NewIPRequest struct {
		Type     IPKind                `json:"type"`
		ServerID string                `json:"server_id"`
		Location *locationObj.Location `json:"location"`
		Cycle    CycleIPMeta           `json:"cycle"`
	}

	NewIPResponse struct {
		IpID    string `json:"ip_id"` // what is the ID of the ip at the provider?
		CIDR    string `bson:"cidr" json:"cidr"`
		Gateway string `bson:"gateway" json:"gateway"`
		Netmask string `bson:"netmask" json:"netmask"`
		Network string `bson:"network" json:"network"`
	}

	DeleteIPRequest struct {
		Type     IPKind                `json:"type"`
		IpID     string                `json:"ip_id"` // what is the ID of the ip at the provider?
		ServerID string                `json:"server_id"`
		Location *locationObj.Location `json:"location"`
		Cycle    CycleIPMeta           `json:"cycle"`
	}

	AssignIPRequest struct {
		Type           IPKind                `json:"type"`
		IpID           string                `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentID string                `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		ServerID       string                `json:"server_id"`
		Location       *locationObj.Location `json:"location"`
		Cycle          CycleIPMeta           `json:"cycle"`
	}

	AssignIPResponse struct {
		IpID           string `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentID string `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
	}

	UnassignIPRequest struct {
		Type           IPKind                `json:"type"`
		IpID           string                `json:"ip_id"`            // what is the ID of the ip at the provider?
		IpAssignmentID string                `json:"ip_assignment_id"` // what is the ID of the ip assignment at the provider?
		ServerID       string                `json:"server_id"`
		Location       *locationObj.Location `json:"location"`
		Cycle          CycleIPMeta           `json:"cycle"`
	}

	CycleIPMeta struct {
		PoolID     *string `json:"pool_id"`
		ServerID   *string `json:"server_id"`
		LocationID *string `json:"location_id"`
	}

	IPKind string
)

const (
	IP_V6 IPKind = "v6"
	IP_V4 IPKind = "v4"
)
