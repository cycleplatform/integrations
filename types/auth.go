package types

type (
	// Providers have many different things they might want during auth, we need to be flexibile.
	Auth struct {
		Namespace    *string        `json:"namespace,omitempty"` // can also double as 'project'
		Region       *string        `json:"region,omitempty"`
		APIKey       *string        `json:"api_key,omitempty"` // can also double as 'token'
		Secret       *string        `json:"secret,omitempty"`
		Base64Config *string        `json:"config,omitempty"` // if a provider needs a var more complex config, sent via base64
		Cycle        *CycleAuthMeta `json:"cycle,omitempty"`
	}

	CycleAuthMeta struct {
		HubName string `json:"hub_name"`
		HubID   string `json:"hub_id"`
	}
)

func (a *Auth) Sanitize() {
	if a.APIKey != nil {
		*a.APIKey = (*a.APIKey)[:5] + "..."
	}

	if a.Secret != nil {
		*a.Secret = (*a.Secret)[:5] + "..."
	}
}
