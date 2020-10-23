package login

import "github.com/pkg/browser"

func OpenAuthPage() error {
	// We're not setting a state at the moment, thus the blank string
	return browser.OpenURL(cfg.AuthCodeURL(""))
}
