// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package exit_test

import (
	"github.com/my10c/packages-go/exit"
	"testing"
)

func TestPrint(t *testing.T) {
	e := exit.New("\nExit test\n", 2)
	e.ExitNil(nil)
	e.Exit("exiting\n")
}
