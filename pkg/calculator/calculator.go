package calculator

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/romandkv/calc/pkg/stack"
)

type operator rune

type calculator struct {
	tempStack   stack.Stack
	resultQueue stack.Queue
}

func (calc *calculator) reset() {
	calc.resultQueue.Reset()
	calc.tempStack.Reset()
}

func (calc *calculator) init() {
	calc.resultQueue = stack.GetQueue()
	calc.tempStack = stack.GetStack()
}

func GetCalculator() calculator {
	calc := calculator{}
	calc.init()
	return calc
}

func (calc calculator) Run(expression string) (float64, error) {
	defer calc.reset()
	calc.init()
	expression = clearWhitespaces(expression)
	if !validateBrackets(expression) {
		return 0, errors.New("Invalid brackets")
	}
	calc.makeNotation(clearWhitespaces(expression))
	return calc.calculate()
}

func validateBrackets(exp string) bool {
	open := 0

	for _, char := range exp {
		if char == '(' {
			open++
		}
		if char == ')' {
			open--
		}
		if open < 0 {
			return false
		}
	}
	return open == 0
}

func (calc *calculator) calculate() (float64, error) {
	operands := stack.GetStack()
	if calc.resultQueue.Length() == 1 {
		return calc.resultQueue.Pop().(float64), nil
	}
	if calc.resultQueue.Length() < 3 {
		return 0, errors.New("Error in expression")
	}
	for {
		value := calc.resultQueue.Pop()
		if _, ok := value.(operator); !ok {
			operands.Push(value)
			continue
		}
		operands.Push(makeOperation(
			operands.Pop().(float64),
			operands.Pop().(float64),
			value.(operator),
		))
		if calc.resultQueue.Length() == 0 {
			return operands.Pop().(float64), nil
		}
	}
}

func makeOperation(right, left float64, operator operator) float64 {
	switch operator {
	case '+':
		return left + right
	case '-':
		return left - right
	case '/':
		return left / right
	case '*':
		return left * right
	}
	return 0
}

func (calc *calculator) makeNotation(expression string) error {
	var err error

	operand := make([]rune, 0, 20)
	for _, char := range expression {
		if !isOperator(char) {
			operand = append(operand, char)
			continue
		}
		calc.pushOperand(string(operand))
		operand = make([]rune, 0, 20)
		calc.handleOperator(operator(char))
		if err != nil {
			return err
		}
	}
	calc.pushOperand(string(operand))
	for calc.tempStack.Head() != nil {
		calc.resultQueue.Push(calc.tempStack.Pop())
	}
	return nil
}

func (calc *calculator) pushOperand(operand string) error {
	if len(operand) == 0 {
		return nil
	}
	parsedOperand, err := strconv.ParseFloat(string(operand), 32)
	if err != nil {
		return err
	}
	calc.resultQueue.Push(parsedOperand)
	return nil
}

func (calc *calculator) handleOperator(op operator) {
	currentPriority := getPriority(op)
	head := calc.tempStack.Head()
	headPriority := -1
	if head != nil {
		headPriority = getPriority(head.(operator))
	}
	if op == '(' {
		calc.tempStack.Push(op)
		return
	}
	if op == ')' {
		for calc.tempStack.Head().(operator) != '(' {
			calc.resultQueue.Push(calc.tempStack.Pop())
		}
		calc.tempStack.Pop()
		return
	}
	if currentPriority < headPriority {
		calc.resultQueue.Push(calc.tempStack.Pop())
	}
	calc.tempStack.Push(op)
}

func getPriority(operator operator) int {
	switch operator {
	case '-':
		return 1
	case '+':
		return 2
	case '/':
		return 3
	case '*':
		return 4
	}
	return -1
}

func isOperator(char rune) bool {
	return char == '+' ||
		char == '-' ||
		char == '*' ||
		char == '/' ||
		char == '(' ||
		char == ')'
}

func clearWhitespaces(str string) string {
	runes := make([]rune, 0, len(str))

	for _, char := range str {
		if !unicode.IsSpace(char) {
			runes = append(runes, char)
		}
	}
	return string(runes)
}
