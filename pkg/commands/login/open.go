package login

import "github.com/pkg/browser"

// Open the auth page in the user's browser
func OpenAuthPage() error {
	// We're not setting a state at the moment, thus the blank string
	return browser.OpenURL(AuthPageURL())
}

// Get the auth page URL
func AuthPageURL() string {
	return cfg.AuthCodeURL("")
}
