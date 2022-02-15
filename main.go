package main

import (
	"fmt"

	"github.com/romandkv/calc/pkg/calculator"
)

func main() {
	calc := calculator.GetCalculator()

	fmt.Println(calc.Run("(8+2*5)/(1+3*(2-4))"))
	fmt.Println(calc.Run("(8+2*5()/(1+3*(2-4))"))
}
