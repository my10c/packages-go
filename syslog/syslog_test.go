// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package syslog_test

import (
	"fmt"
	"github.com/my10c/packages-go/syslog"
	"testing"
)

func printvalue() {
	fmt.Printf("Hello World\n\n")
	fmt.Printf("syslog function is %s\n", syslog.syslogFunc())
}

func Testsyslog(t *testing.T) {
	fmt.Printf("Calling printvalue\n")
	printvalue()
}
