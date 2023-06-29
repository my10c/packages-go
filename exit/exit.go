// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package exit

import (
	"fmt"
	"os"
)

const (
	off = "\x1b[0m"    // Text Reset
	red = "\x1b[1;31m" // red
)

type (
	Exit struct {
		ExitMsg   string
		ExitValue int
	}
)

func New(exitMsg string, exitValue int) *Exit {
	return &Exit{
		ExitMsg:   exitMsg,
		ExitValue: exitValue,
	}
}

func (e *Exit) ExitError(err error) {
	if err != nil {
		if len(e.ExitMsg) > 0 {
			fmt.Fprintln(os.Stderr, e.ExitMsg)
		}
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(e.ExitValue)
	}
}

// function to exit if pointer is nill
func (e *Exit) ExitNil(ptr interface{}) {
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "Error: got a nil pointer."+e.ExitMsg)
		os.Exit(e.ExitValue)
	}
}

// function to exit wih message
func (e *Exit) Exit(msg string) {
	fmt.Printf("%s%s%s", red, msg, off)
	os.Exit(e.ExitValue)
}
