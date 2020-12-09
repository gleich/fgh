package utils

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/v6/progress"
)

var (
	// Character set used for all spinners
	SpinnerCharSet = spinner.CharSets[14]
	// Speed user for all spinners
	SpinnerSpeed = 30 * time.Millisecond
)

// Generate a progress writer
func GenerateProgressWriter() progress.Writer {
	pw := progress.NewWriter()
	pw.SetTrackerLength(30)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.SetUpdateFrequency(SpinnerSpeed)
	return pw
}
