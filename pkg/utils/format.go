package utils

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	tf "github.com/hepsiburada/time-formatter"
)

// Format date in the following format:
// December 25th, 2020 at 12:00PM
func FormatDate(date time.Time) string {
	var (
		formatter     = tf.New(tf.EN)
		hour          = date.Hour()
		formattedHour string
	)

	if hour > 12 {
		formattedHour = fmt.Sprint(hour - 12)
	} else {
		formattedHour = fmt.Sprint(hour)
	}

	return formatter.To(date, fmt.Sprintf(
		"%s %s, %s at %v:%02v%s",
		tf.TIME_MMMM,
		humanize.Ordinal(date.Day()),
		tf.TIME_YYYY,
		formattedHour,
		date.Minute(),
		tf.TIME_A,
	))
}
