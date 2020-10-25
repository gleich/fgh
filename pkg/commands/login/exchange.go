package login

import (
	"context"
	"time"
)

// Exchange an OAuth authorization code for an access token
func Exchange(code string) (string, error) {
	// Time out after 20 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
