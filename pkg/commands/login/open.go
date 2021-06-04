package login

import (
	"github.com/gleich/fgh/pkg/utils"
	"github.com/pkg/browser"
)

// Open the auth page in the user's browser
func OpenAuthPage() utils.CtxErr {
	var (
		url = AuthPageURL()
		err = browser.OpenURL(url)
	)
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to open auth page in browser. Please open it manually:\n" + url,
			Error:   err,
		}
	}
	return utils.CtxErr{}
}

// Get the auth page URL
func AuthPageURL() string {
	return cfg.AuthCodeURL("")
}
