package clean

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	tt := []struct {
		date   time.Time
		output string
	}{
		{
			date:   time.Date(2020, 10, 22, 14, 59, 37, 0, time.UTC),
			output: "October 22nd, 2020 at 14:59PM",
		},
		{
			date:   time.Date(2021, 1, 20, 20, 20, 20, 20, time.UTC),
			output: "January 20th, 2021 at 20:20PM",
		},
	}

	for _, test := range tt {
		assert.Equal(t, formatDate(test.date), test.output)
	}
}
