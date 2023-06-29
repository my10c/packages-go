// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package signalhandler

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type (
	Handler struct {
		handlerFunc func(string)
	}
)

func New(fn func(string)) *Handler {
	return &Handler{
		handlerFunc: fn,
	}
}

// function to capture reveived signal
// the handler function must use string as argument!
func (s *Handler) SignalHandler(funcData string) {
	interrupt := make(chan os.Signal, 1)

	// we handle only these signal: SIGINT(2) - SIGTRAP(5) - SIGKILL(9) - SIGTERM(15), SIGHUP(1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTRAP, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sigId := <-interrupt
		fmt.Fprintln(os.Stderr, "Received signal "+fmt.Sprintf("%v %d", sigId, sigId))
		exit, _ := strconv.Atoi(fmt.Sprintf("%d", sigId))
		s.handlerFunc(funcData)
		os.Exit(exit)
	}()
}
