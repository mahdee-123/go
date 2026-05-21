package main

import (
	"github.com/arnob/go/lucky-number/internal/random"
	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen)
	green.Printf("Your lucky number is %d\n", random.Number())
}