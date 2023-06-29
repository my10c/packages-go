// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package caller_test

import (
	"fmt"
	"github.com/my10c/packages-go/caller"
	"testing"
)

func printvalue() {
	fmt.Printf("Hello World\n\n")
	fmt.Printf("Caller function is %s\n", caller.CallerFunc())
}

func TestCaller(t *testing.T) {
	fmt.Printf("Calling printvalue\n")
	printvalue()
}
