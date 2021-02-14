package utils

// An error with specific context to said error
type CtxErr struct {
	Error   error
	Context string
}
