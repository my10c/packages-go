// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package random_test

import (
	"fmt"
	"github.com/my10c/packages-go/random"
	"testing"
)

func TestPrint(t *testing.T) {
	r := random.New(true, 25)
	fmt.Printf("\t%s\n", r.Generate())
	fmt.Printf("\t%s\n", r.Alpha())
	fmt.Printf("\t%s\n", r.Numeric())
	r.SetLength(16)
	fmt.Printf("\t%s\n", r.Generate())
	fmt.Printf("\t%s\n", r.Alpha())
	fmt.Printf("\t%s\n", r.Numeric())
}
