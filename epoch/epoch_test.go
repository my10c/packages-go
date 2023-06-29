// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package epoch_test

import (
	"fmt"
	"github.com/my10c/packages-go/epoch"
	"testing"
)

func TestPrint(t *testing.T) {
	e := epoch.New()
	fmt.Printf("weeks   : %v \n", e.Weeks())
	fmt.Printf("days    : %v \n", e.Days())
	fmt.Printf("hours   : %v \n", e.Hours())
	fmt.Printf("minutes : %v \n", e.Minutes())
	fmt.Printf("seconds : %v \n", e.Seconds())
	fmt.Printf("current : %v \n", e.Curent())
	fmt.Printf("date    : %v \n", e.Date())
	newEpoch, newDate := e.NotToday(180)
	fmt.Printf("+180 => epoch: %v date: %v \n", newEpoch, newDate)
	newEpoch, newDate = e.NotToday(-180)
	fmt.Printf("-180 => epoch: %v date: %v \n", newEpoch, newDate)
}
