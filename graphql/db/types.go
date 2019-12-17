package db

import (
	"github.com/volatiletech/null"
	"strconv"
)

// Interval holds information about an interval
type Interval struct {
	Seconds null.Int
}

// Format formats interval to a string
func (i *Interval) Format() null.String {
	if !i.Seconds.Valid {
		return null.String{String: "", Valid: false}
	}
	return null.StringFrom(strconv.Itoa(i.Seconds.Int) + " seconds")
}
