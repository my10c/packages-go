// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package lock

import (
	"bufio"
	"os"
	"strconv"
)

type (
	Lock struct {
		LockFile string
		LockPid  int
	}
)

func New(lockFile string) *Lock {
	return &Lock{
		LockFile: lockFile,
	}
}

func (l *Lock) LockIt(lockPid int) error {
	l.LockPid = lockPid

	fp, err := os.Create(l.LockFile)
	if err != nil {
		return err
	}

	defer fp.Close()
	_, err = fp.WriteString(strconv.Itoa(l.LockPid))
	if err != nil {
		return err
	}
	return nil
}

func (l *Lock) LockRelease() error {
	return os.Remove(l.LockFile)
}

func (l *Lock) LockGetPid() (int, error) {
	var currLine string
	fp, err := os.Open(l.LockFile)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		currLine = scanner.Text()
		break
	}
	return strconv.Atoi(currLine)
}
