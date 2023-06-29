// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package readinput_test

import (
	"fmt"
	"github.com/my10c/packages-go/print"
	"github.com/my10c/packages-go/readinput"
	"testing"
)

func TestReadInput(t *testing.T) {
	data := "Colorado"
	readinput.ReadInput(print.Blue, "Hello World: ")
	ok := readinput.ReadYN("y", false)
	fmt.Printf("\ndata %s : ok %t\n", data, ok)
}
