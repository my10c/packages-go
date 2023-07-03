// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package md5_test

import (
	"fmt"
	"github.com/my10c/packages-go/md5"
	"testing"
)

func printvalue() {
	fmt.Printf("Hello World\n\n")
	fmt.Printf("md5 function is %s\n", md5.md5Func())
}

func Testmd5(t *testing.T) {
	fmt.Printf("Calling printvalue\n")
	printvalue()
}
