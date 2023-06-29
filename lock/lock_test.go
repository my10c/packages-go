// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package lock_test

import (
	"fmt"
	"github.com/my10c/packages-go/lock"
	"os"
	"testing"
)

func TestLock(t *testing.T) {
	lockPid := os.Getpid()
	lockFile := "/tmp/lock-test.lock"

	l := lock.New(lockFile)
	err := l.LockIt(lockPid)
	if err != nil {
		t.Fatalf("Errored with %s\n", err.Error())
	}

	pid, err := l.LockGetPid()
	if err != nil {
		t.Fatalf("Errored with %s\n", err.Error())
	}
	fmt.Printf("Pid %d\n", pid)

	err = l.LockRelease()
	if err != nil {
		t.Fatalf("Errored with %s\n", err.Error())
	}
}
