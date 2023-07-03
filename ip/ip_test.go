// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package ip_test

import (
	"fmt"
	"github.com/my10c/packages-go/ip"
	"testing"
)

func printvalue() {
	fmt.Printf("Hello World\n\n")
	fmt.Printf("ip function is %s\n", ip.ipFunc())
}

func Testip(t *testing.T) {
	fmt.Printf("Calling printvalue\n")
	printvalue()
}
