package types

type (
	Auth struct {
		Project      *string       `json:"project,omitempty"` // can also double as 'namespace'
		APIKey       *string       `json:"api_key,omitempty"` // can also double as 'token'
		Secret       *string       `json:"secret,omitempty"`
		Base64Config *string       `json:"config,omitempty"` // if a provider needs a var more complex config, sent via base64
		Cycle        CycleAuthMeta `json:"cycle"`
	}

	CycleAuthMeta struct {
		HubName string `json:"hub_name"`
		HubID   string `json:"hub_id"`
	}
)
