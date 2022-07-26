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

	// output from POST /v1/auth/verify
	AuthVerifyResponse struct {
		Valid bool `json:"valid"`
	}

	CycleAuthMeta struct {
		HubName string `json:"hub_name"`
		HubID   string `json:"hub_id"`
	}
)

func (a *Auth) Sanitize() {
	if a.APIKey != nil && len(*a.APIKey) > 5 {
		*a.APIKey = (*a.APIKey)[:5] + "..."
	}

	if a.Secret != nil && len(*a.Secret) > 5 {
		*a.Secret = (*a.Secret)[:5] + "..."
	}
}
