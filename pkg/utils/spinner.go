package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

var (
	// Character set used for all spinners
	SpinnerCharSet = spinner.CharSets[14]
	// Speed user for all spinners
	SpinnerSpeed = 100 * time.Millisecond
)
