package types

type (
	NewIPRequest struct {
		Type     IPKind      `json:"type"`
		Location string      `json:"location"`
		Cycle    CycleIPMeta `json:"cycle"`
	}

	DeleteIPRequest struct {
		ProviderId string      `json:"provider_id"` // what is the ID of the ip at the provider?
		Type       IPKind      `json:"type"`
		Cycle      CycleIPMeta `json:"cycle"`
	}

	AssignIPRequest struct {
		ProviderId           string      `json:"provider_id"`            // what is the ID of the ip at the provider?
		ProviderAssignmentId string      `json:"provider_assignment_id"` // if the provider requires a different ID for assignments
		Type                 IPKind      `json:"type"`
		Location             string      `json:"location"`
		Cycle                CycleIPMeta `json:"cycle"`
	}

	UnassignIPRequest struct {
		ProviderId           string      `json:"provider_id"`            // what is the ID of the ip at the provider?
		ProviderAssignmentId string      `json:"provider_assignment_id"` // if the provider requires a different ID for assignments
		Type                 IPKind      `json:"type"`
		Location             string      `json:"location"`
		Cycle                CycleIPMeta `json:"cycle"`
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
