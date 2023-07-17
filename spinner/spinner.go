// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package spinner

import (
	"fmt"
	"os"
	"time"
)

type (
	Wheel struct {
		StopChannel chan int
		Speed       time.Duration
	}
)

const (
	off       = "\x1b[0m"
	purple    = "\x1b[1;35m"
	clearLine = "\x1b[0G\x1b[2K\x1b[0m\r"
	clearCounter = "\x1b[0G\x1b[2K\r"
)

// the spinner class
func New(speed time.Duration) *Wheel {
	return &Wheel{
		StopChannel: make(chan int),
		Speed:       speed,
	}
}

// run the spinner until a StopChannel was reveived
func (w *Wheel) Run() {
	var spinstr = "| / - \\"
	fmt.Fprintf(os.Stdout, "\x1b[?25l")
	for {
		select {
		case <-w.StopChannel:
			w.StopChannel <- 1
			return

		default:
			break

		}
		for _, c := range spinstr {
			fmt.Printf("[%c]", c)
			time.Sleep(w.Speed * time.Millisecond)
			fmt.Printf("\b\b\b\b\b\b%s", purple)
		}
	}
}

// stop the spinner by sending 1 to the StopChannel
func (w *Wheel) Stop() {
	// shutdown the spinner
	w.StopChannel <- 1
	<-w.StopChannel
	close(w.StopChannel)

	// show cursor
	fmt.Fprint(os.Stdout, off)
	fmt.Fprint(os.Stdout, "\x1b[?25h")
	fmt.Fprint(os.Stdout, clearLine)
}

// counter
func (w *Wheel) Counter(count int, msg string) {
	cnt := count
	// hide cursore
	fmt.Fprintf(os.Stdout, "\x1b[?25l")
	fmt.Printf(purple)
	for cnt > 0 {
		fmt.Printf("%s [%d]", msg, cnt)
		time.Sleep(w.Speed * time.Millisecond)
		fmt.Printf(clearCounter)
		cnt--
	}
	// show cursor
	fmt.Fprint(os.Stdout, off)
	fmt.Fprint(os.Stdout, "\x1b[?25h")
	fmt.Fprint(os.Stdout, clearLine)
}
