package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GoogleTokenInfo contains the verified claims from a Google ID token.
type GoogleTokenInfo struct {
	Sub           string `json:"sub"`            // Google user ID (stable, unique)
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Aud           string `json:"aud"`            // Must match our client ID
}

var googleHTTPClient = &http.Client{Timeout: 5 * time.Second}

// VerifyGoogleIDToken validates a Google ID token by calling Google's tokeninfo endpoint.
// Returns the token claims if valid, or an error.
func VerifyGoogleIDToken(idToken string, expectedClientIDs []string) (*GoogleTokenInfo, error) {
	client := googleHTTPClient

	resp, err := client.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken)
	if err != nil {
		return nil, fmt.Errorf("google token verification request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid google token (status %d)", resp.StatusCode)
	}

	var info GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode google token info: %w", err)
	}

	// Verify audience matches one of our client IDs
	audValid := false
	for _, cid := range expectedClientIDs {
		if cid != "" && info.Aud == cid {
			audValid = true
			break
		}
	}
	if !audValid {
		return nil, fmt.Errorf("google token audience mismatch: got %q", info.Aud)
	}

	if info.Sub == "" {
		return nil, fmt.Errorf("google token missing sub claim")
	}

	return &info, nil
}
