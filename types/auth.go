package types

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type (
	// Providers have many different things they might want during auth, we need to be flexibile.
	Auth struct {
		ClientId       *string        `json:"client_id,omitempty"`
		SubscriptionId *string        `json:"subscription_id,omitempty"`
		Namespace      *string        `json:"namespace,omitempty"` // can also double as 'project or subscription id'
		Region         *string        `json:"region,omitempty"`
		APIKey         *string        `json:"api_key,omitempty"` // can also double as 'token'
		Secret         *string        `json:"secret,omitempty"`
		Base64Config   *string        `json:"config,omitempty"` // if a provider needs a var more complex config, sent via base64
		Cycle          *CycleAuthMeta `json:"cycle,omitempty"`
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

func (a *Auth) Dirty(orig Auth) {
	if a.APIKey != nil && strings.HasSuffix(*a.APIKey, "...") {
		*a.APIKey = *orig.APIKey
	}

	if a.Secret != nil && strings.HasSuffix(*a.Secret, "...") {
		*a.Secret = *orig.Secret
	}
}

func (a *Auth) Validate() error {
	if a.SubscriptionId != nil && strings.TrimSpace(*a.SubscriptionId) == "" {
		return fmt.Errorf("a subscription id cannot be an empty string. if no subscription id is required, use null.")
	}

	if a.ClientId != nil && strings.TrimSpace(*a.ClientId) == "" {
		return fmt.Errorf("a client id cannot be an empty string. if no client id is required, use null.")
	}

	if a.APIKey != nil && strings.TrimSpace(*a.APIKey) == "" {
		return fmt.Errorf("a api key cannot be an empty string. if no api key is required, use null.")
	}

	if a.Region != nil && strings.TrimSpace(*a.Region) == "" {
		return fmt.Errorf("a region cannot be an empty string. if no region is required, use null.")
	}

	if a.Namespace != nil && strings.TrimSpace(*a.Namespace) == "" {
		return fmt.Errorf("a namespace cannot be an empty string. if no namespace is required, use null.")
	}

	if a.Secret != nil && strings.TrimSpace(*a.Secret) == "" {
		return fmt.Errorf("a secret cannot be an empty string. if no secret is required, use null.")
	}

	if a.Base64Config != nil {
		if strings.TrimSpace(*a.Base64Config) == "" {
			return fmt.Errorf("a config cannot be an empty string. if no config is required, use null.")
		}

		if _, err := base64.StdEncoding.DecodeString(*a.Base64Config); err != nil {
			return fmt.Errorf("config is not a valid base64 string")
		}
	}

	return nil
}
