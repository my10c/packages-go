// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package signalhandler_test

import (
	"fmt"
	s "github.com/my10c/packages-go/signalhandler"
	"testing"
	"time"
)

func HelloWorld(world string) {
	fmt.Printf("Hello %s world\n", world)
}

func TestSignalhandler(t *testing.T) {
	s := s.New(HelloWorld)
	s.SignalHandler("Colorado")
	fmt.Printf("Hit Control-C to test within 10 seconds, test should return failed!\n")
	time.Sleep(10 * time.Second)
}
