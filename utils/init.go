package utils

import (
	"fmt"
	"strings"
)

// This clears the terminal
func Clear() {
	print("\033[H\033[2J")
}

func TerminalTitle(symbol string, price float64) {
	fmt.Printf("\033]0;%s: %.2f \007", strings.ToUpper(symbol), price)
}
