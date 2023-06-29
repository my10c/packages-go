// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package print_test

import (
	"fmt"
	"github.com/my10c/packages-go/print"
	"testing"
)

func TestPrint(t *testing.T) {
	p := print.New()
	fmt.Printf("%s", p.ClearScreen())
	fmt.Printf("\t%s\n", p.PrintHeader(print.Blue, print.Purple, "hello world", 20, false))
	p.PrintRed("\tHello World\n")
	fmt.Printf("\t%s\n", p.PrintDangerZone())
	fmt.Printf("\t%s\n", p.PrintColorMessage(print.Cyan, "Hello World"))
	p.TheEnd()
}
