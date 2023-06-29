// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package epoch

import (
	"time"
)

type (
	Epoch struct {
		CurrentEpoch int64
		CurrentDate  string
	}
)

const (
	USformat = "Jan 2, 2006 15:04:05"
)

func New() *Epoch {
	return &Epoch{
		CurrentEpoch: time.Now().Unix(),
		CurrentDate:  time.Now().Format(USformat),
	}
}

func (e *Epoch) Weeks() int64 {
	return (e.CurrentEpoch / (86400 * 7))
}

func (e *Epoch) Days() int64 {
	return (e.CurrentEpoch / 86400)
}

func (e *Epoch) Hours() int64 {
	return ((e.CurrentEpoch / 86400) * 24)
}

func (e *Epoch) Minutes() int64 {
	return ((e.CurrentEpoch / 86400) * (60 * 24))
}

func (e *Epoch) Seconds() int64 {
	return e.CurrentEpoch
}

func (e *Epoch) Curent() int64 {
	return e.CurrentEpoch
}

func (e *Epoch) Date() string {
	return e.CurrentDate
}

func (e *Epoch) ReadableEpoch(epoch int64) string {
	value := time.Unix(epoch, 0)
	return value.Format(USformat)
}

func (e *Epoch) NotToday(days int) (int64, string) {
	delta := 86400 * int64(days)
	newEpoch := e.CurrentEpoch + delta
	newDate := time.Unix(newEpoch, 0)
	return newEpoch, newDate.Format(USformat)
}
