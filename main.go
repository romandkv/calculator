package main

import (
	"fmt"

	"github.com/romandkv/calc/pkg/calculator"
)

func main() {
	calc := calculator.GetCalculator()

	fmt.Println(calc.Run("1 + 3 * 6 / 2"))
}
