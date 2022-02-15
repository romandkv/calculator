package main

import (
	"fmt"
	"log"
	"os"

	"github.com/romandkv/calculator/pkg/calculator"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("expression isnt defined")
	}
	calc := calculator.GetCalculator()

	fmt.Println(calc.Run(os.Args[1]))
}
