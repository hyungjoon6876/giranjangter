package oauth

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

// GoogleTokenInfo contains the verified claims from a Google ID token.
type GoogleTokenInfo struct {
	Sub   string
	Email string
	Name  string
}

// VerifyGoogleIDToken validates a Google ID token locally using Google's public keys.
func VerifyGoogleIDToken(token string, expectedClientIDs []string) (*GoogleTokenInfo, error) {
	ctx := context.Background()

	for _, clientID := range expectedClientIDs {
		if clientID == "" {
			continue
		}
		payload, err := idtoken.Validate(ctx, token, clientID)
		if err != nil {
			continue
		}

		sub, _ := payload.Claims["sub"].(string)
		email, _ := payload.Claims["email"].(string)
		name, _ := payload.Claims["name"].(string)

		if sub == "" {
			return nil, fmt.Errorf("google token missing sub claim")
		}

		return &GoogleTokenInfo{Sub: sub, Email: email, Name: name}, nil
	}

	return nil, fmt.Errorf("google token validation failed for all client IDs")
}
