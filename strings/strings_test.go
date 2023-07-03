// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package strings_test

import (
	"fmt"
	"github.com/my10c/packages-go/strings"
	"testing"
)

func printvalue() {
	fmt.Printf("Hello World\n\n")
	fmt.Printf("strings function is %s\n", strings.stringsFunc())
}

func Teststrings(t *testing.T) {
	fmt.Printf("Calling printvalue\n")
	printvalue()
}
