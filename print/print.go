// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package print

import (
	"fmt"
	"strings"
)

type (
	Print struct {
		Color string
	}
)

var (
	Off    = "\x1b[0m"    // Text Reset
	Black  = "\x1b[1;30m" // Black
	Red    = "\x1b[1;31m" // Red
	Green  = "\x1b[1;32m" // Green
	Yellow = "\x1b[1;33m" // Yellow
	Blue   = "\x1b[1;34m" // Blue
	Purple = "\x1b[1;35m" // Purple
	Cyan   = "\x1b[1;36m" // Cyan
	White  = "\x1b[1;37m" // White

	RedBase    = "\x1b[0;31m" // Red no highlighted
	Greenbase  = "\x1b[0;32m" // Green no highlighted
	YellowBase = "\x1b[0;33m" // Yellow no highlighted
	BlueBase   = "\x1b[0;34m" // Blue no highlighted
	PurpleBase = "\x1b[0;35m" // Purple no highlighted
	CyanBase   = "\x1b[0;36m" // Cyan no highlighted
	WhiteBase  = "\x1b[0;37m" // White no highlighted

	RedUnderline = "\x1b[4;31m" // Red underline
	OneLineUP    = "\x1b[A"

	ClearLine   = "\x1b[0G\x1b[2K\x1b[0m\r"
	ClearScreen = "\x1b[H\x1b[2J"
	HEADER      = "---------------"
	LINE        = "_________________________________________________"

	dangerZone = fmt.Sprintf("%sDanger Zone%s, be sure you understand the implication!",
		RedUnderline, Off)
)

func New() *Print {
	return &Print{}
}

// functions to print messsage in the given color
func (p *Print) PrintRed(message string) {
	fmt.Printf("%s%s%s", Red, message, Off)
}

func (p *Print) PrintGreen(message string) {
	fmt.Printf("%s%s%s", Green, message, Off)
}

func (p *Print) PrintYellow(message string) {
	fmt.Printf("%s%s%s", Yellow, message, Off)
}

func (p *Print) PrintBlue(message string) {
	fmt.Printf("%s%s%s", Blue, message, Off)
}

func (p *Print) PrintPurple(message string) {
	fmt.Printf("%s%s%s", Purple, message, Off)
}

func (p *Print) PrintCyan(message string) {
	fmt.Printf("%s%s%s", Cyan, message, Off)
}

func (p *Print) MessageRed(message string) string {
	return fmt.Sprintf("%s%s%s", Red, message, Off)
}

func (p *Print) MessageGreen(message string) string {
	return fmt.Sprintf("%s%s%s", Green, message, Off)
}

func (p *Print) MessageYellow(message string) string {
	return fmt.Sprintf("%s%s%s", Yellow, message, Off)
}

func (p *Print) MessageBlue(message string) string {
	return fmt.Sprintf("%s%s%s", Blue, message, Off)
}

func (p *Print) MessagePurple(message string) string {
	return fmt.Sprintf("%s%s%s", Purple, message, Off)
}

func (p *Print) MessageCyan(message string) string {
	return fmt.Sprintf("%s%s%s", Cyan, message, Off)
}

// function Message line
func (p *Print) PrintLine(lineColor string, count int) string {
	line := strings.Repeat("⎻", count)
	return fmt.Sprintf("%s%s%s", lineColor, line, Off)
}

// function to clear the screen
func (p *Print) ClearScreen() string {
	return fmt.Sprintf("%s", ClearScreen)
}

// function to move one line up
func (p *Print) OneLineUP(clearTheLine bool) string {
	out := fmt.Sprintf("%s", OneLineUP)
	if clearTheLine {
		out = out + fmt.Sprintf("%s", ClearLine)
	}
	return out
}

func (p *Print) PrintDangerZone() string {
	return fmt.Sprintf("%s", dangerZone)
}

func (p *Print) PrintColorMessage(messageColor, message string) string {
	return fmt.Sprintf("%s%s%s", messageColor, message, Off)
}

func (p *Print) PrintHeader(edgeColor, messageColor, message string, edgeCount int, clearScreen bool) string {
	if clearScreen {
		fmt.Printf("%s", p.ClearScreen())
	}
	edges := p.PrintLine(edgeColor, edgeCount)
	msg := p.PrintColorMessage(messageColor, message)
	return fmt.Sprintf("%s %s %s", edges, msg, edges)
}

// The End Message
func (p *Print) TheEnd() {
	p.PrintGreen("\tEnjoy a cuppa of hot cOffee ☕️/🥃 \n")
	p.PrintGreen("\tThe End\n")
}
