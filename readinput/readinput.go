// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package readinput

import (
	"bufio"
	"fmt"
	exit "github.com/my10c/packages-go/exit"
	print "github.com/my10c/packages-go/print"
	"os"
	"strings"
)

// function to read input from keyboard with a message before reading the input
func ReadInput(messageColor string, messsage string) string {

	p := print.New()
	e := exit.New("ReadInput error", 1)

	fmt.Printf("%s", p.PrintColorMessage(messageColor, messsage))

	reader := bufio.NewReader(os.Stdin)
	valueEntered, err := reader.ReadString('\n')
	valueEntered = strings.TrimSuffix(valueEntered, "\n")
	e.ExitNil(err)
	return valueEntered
}

// function to check for given y/n
func ReadYN(keyboardInput string, defaultReturn bool) bool {
	yesSelection := []string{"yes", "y", "YES", "Y"}

	if keyboardInput == "" {
		return defaultReturn
	}
	for _, selected := range yesSelection {
		if strings.EqualFold(string(selected), keyboardInput) {
			return true
		}
	}
	return defaultReturn
}
