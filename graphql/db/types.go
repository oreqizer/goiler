package db

import (
	"github.com/volatiletech/null"
	"strconv"
)

type Average struct {
	Average float64
}

type Interval struct {
	Seconds null.Int
}

func (i *Interval) Format() null.String {
	if !i.Seconds.Valid {
		return null.String{String: "", Valid: false}
	}
	return null.StringFrom(strconv.Itoa(i.Seconds.Int) + " seconds")
}
