package login

import (
	"context"
	"time"
)

func Exchange(code string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
