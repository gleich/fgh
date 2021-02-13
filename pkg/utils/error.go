package utils

import "github.com/Matt-Gleich/statuser/v2"

// An error with specific context to said error
type CtxErr struct {
	Error   error
	Context string
}

// Log errors
func LogErr(err CtxErr) {
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
