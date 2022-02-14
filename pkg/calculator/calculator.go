package calculator

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/romandkv/calc/pkg/stack"
)

type calculator struct {
	operatorQueue stack.Queue
	tempStack     stack.Stack
	resultStack   stack.Stack
}

func (calc *calculator) reset() {
	calc.operatorQueue.Reset()
	calc.resultStack.Reset()
	calc.tempStack.Reset()
}

func (calc *calculator) init() {
	calc.operatorQueue = stack.GetQueue()
	calc.resultStack = stack.GetStack()
	calc.tempStack = stack.GetStack()
}

func GetCalculator() calculator {
	calc := calculator{}
	calc.init()
	return calc
}

func (calc calculator) Run(expression string) (float64, error) {
	calc.init()
	calc.makeNotation(clearWhitespaces(expression))
	return calc.calculate()
}

func (calc *calculator) calculate() (float64, error) {
	var result float64 = 0
	var tempNum *float64

	if calc.operatorQueue.Length()+1 != calc.resultStack.Length() {
		return 0, errors.New("Error in expression")
	}
	if calc.resultStack.Length() == 0 {
		return calc.resultStack.Pop().(float64), nil
	}
	for {
		if calc.operatorQueue.Head() == nil && calc.resultStack.Head() == nil {
			return result, nil
		}
		if tempNum == nil {
			tempNum = new(float64)
			*tempNum = calc.resultStack.Pop().(float64)
		}
		result = makeOperation(
			*tempNum,
			calc.resultStack.Pop().(float64),
			calc.operatorQueue.Pop().(rune),
		)
		*tempNum = result
	}
}

func makeOperation(left, right float64, operator rune) float64 {
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
		calc.handleOperator(char)
		if err != nil {
			return err
		}
	}
	calc.pushOperand(string(operand))
	for calc.tempStack.Head() != nil {
		calc.operatorQueue.Push(calc.tempStack.Pop())
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
	calc.resultStack.Push(parsedOperand)
	return nil
}

func (calc *calculator) handleOperator(operator rune) {
	currentPriority := getPriority(operator)
	head := calc.tempStack.Head()
	headPriority := -1
	if head != nil {
		headPriority = getPriority(head.(rune))
	}
	if currentPriority < headPriority {
		calc.operatorQueue.Push(calc.tempStack.Pop())
	}
	calc.tempStack.Push(operator)
}

func getPriority(operator rune) int {
	switch operator {
	case '-':
		return 0
	case '+':
		return 1
	case '*':
		return 2
	case '/':
		return 3
	case ')':
		return 4
	case '(':
		return 5
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
