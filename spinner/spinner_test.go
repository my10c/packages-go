// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package spinner_test

import (
	"fmt"
	"github.com/my10c/packages-go/spinner"
	"testing"
	"time"
)

func TestSpinne(t *testing.T) {
	s := spinner.New(100)
	fmt.Printf("start\n")
	go s.Run()
	time.Sleep(10 * time.Second)
	s.Stop()
	fmt.Printf("\nstoped\n")
}
