// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package is_test

import (
	"fmt"
	"github.com/my10c/packages-go/is"
	"testing"
)

func TestIs(t *testing.T) {
	i := is.New()
	ok, _ := i.IsUser("luc")

	fmt.Printf("is root %t\n", i.IsRoot())
	fmt.Printf("check is luc  %t\n", ok)
	fmt.Printf("OS is %s\n", i.IsOS())

	msg, ok := i.IsFileOwner("./is_test.go", []string{"root", "nobody", "luc"})
	fmt.Printf("Msg %s, stat %t\n", msg, ok)

	msg, ok = i.IsFileOwner("./is_test.go", []string{"root", "nobody", "bin"})
	fmt.Printf("Msg %s, stat %t\n", msg, ok)

	msg, ok = i.IsFilePermission("./is_test.go", []string{"0600", "0644"})
	fmt.Printf("Msg %s, stat %t\n", msg, ok)

	msg, ok, err := i.IsExist("/var/doesnotexiast", "directory")
	if err != nil {
		fmt.Printf("Msg %s, stat %t, %v\n", msg, ok, err.Error())
	} else {
		fmt.Printf("Msg %s, stat %tn", msg, ok)
	}

	msg, ok, err = i.IsExist("/etc/passwd", "file")
	if err != nil {
		fmt.Printf("Msg %s, stat %t, error %v\n", msg, ok, err.Error())
	} else {
		fmt.Printf("Msg %s, stat %t\n", msg, ok)
	}

	msg, err = i.IsRunningUser()
	if err != nil {
		fmt.Printf("Msg %s, error %v\n", msg, err.Error())
	} else {
		fmt.Printf("Msg %s, \n", msg)
	}
}
