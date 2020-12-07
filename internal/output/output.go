package output

import (
	"fmt"

	"github.com/fatih/color"
)

// Info prints an info message to screen
func Info(message string) {
	blue := color.New(color.FgBlue)
	blue.Print(">")
	fmt.Print(" ")
	fmt.Print(message)
	fmt.Print("\n")
}

// Success prints a success message to screen
func Success(message string) {
	green := color.New(color.FgGreen)
	green.Print("✓")
	fmt.Print(" ")
	whiteBold.Print(message)
	fmt.Print("\n")
}

// Warning prints a warning message to screen
func Warning(message string) {
	yellow := color.New(color.FgYellow)
	yellow.Print("!")
	fmt.Print(" ")
	whiteBold.Print(message)
	fmt.Print("\n")
}

// Error prints an error message to screen
func Error(message string) {
	red := color.New(color.FgRed)
	red.Print("✗")
	fmt.Print(" ")
	whiteBold.Print(message)
	fmt.Print("\n")
}
