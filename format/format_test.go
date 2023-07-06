// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package format_test

import (
	"fmt"
	"github.com/my10c/packages-go/format"
	"github.com/my10c/packages-go/print"
	"testing"
)

func TestPrint(t *testing.T) {
	var isBig int64 = 1000000
	p := print.New()
	fmt.Printf("%s", p.ClearScreen())
	fmt.Printf("\t%s\n", p.PrintHeader(print.Blue, print.Purple, "hello world", 20, false))
	fmt.Printf("\tValue is %s\n", format.Format(isBig))
	fmt.Printf("\t%s\n", p.PrintDangerZone())
	fmt.Printf("\t%s\n", p.PrintColorMessage(print.Cyan, "Hello World"))
	p.TheEnd()
}
