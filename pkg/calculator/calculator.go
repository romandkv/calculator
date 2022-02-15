package calculator

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/romandkv/calculator/pkg/stack"
)

// Оболочка для оператора
type operator rune

// Структура калькулятора
type calculator struct {
	tempStack   stack.Stack
	resultQueue stack.Queue
}

// Возвращение калькулятора в исходное состояние, чтобы запускать несколько раз
func (calc *calculator) reset() {
	calc.resultQueue.Reset()
	calc.tempStack.Reset()
}

// Инициализация калькулятора
func (calc *calculator) init() {
	calc.resultQueue = stack.GetQueue()
	calc.tempStack = stack.GetStack()
}

// Получение экземпляра калькулятора
func GetCalculator() calculator {
	calc := calculator{}
	calc.init()
	return calc
}

// Запуск решения выражения
func (calc calculator) Run(expression string) (float64, error) {
	defer calc.reset()
	calc.init()
	expression = clearWhitespaces(expression)
	if !validateBrackets(expression) {
		return 0, errors.New("Invalid brackets")
	}
	err := calc.makeNotation(clearWhitespaces(expression))
	if err != nil {
		return 0, err
	}
	return calc.calculate()
}

// Проверка выражения на количество скобочек и их порядок
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

// Расчет результата с готовой обратной польской нотацией
func (calc *calculator) calculate() (float64, error) {
	// Стэк для хранения операндов
	operands := stack.GetStack()
	// Случай когда пользователь привел просто константу
	if calc.resultQueue.Length() == 1 {
		return calc.resultQueue.Pop().(float64), nil
	}
	// Если выражение не является константой, то в очереди
	// должны ледатлежать как минимум два операнда и один знак
	if calc.resultQueue.Length() < 3 {
		return 0, errors.New("Error in expression")
	}
	for {
		value := calc.resultQueue.Pop()
		// Если полученный элемент не является оператором,
		// то является операндом,  пушим его в стек операндов 
		if _, ok := value.(operator); !ok {
			operands.Push(value)
			continue
		}
		// Если элемент является оператором, применяем его
		// на двух последних операндах в стеке
		operands.Push(makeOperation(
			operands.Pop().(float64),
			operands.Pop().(float64),
			value.(operator),
		))
		// Заканчиваем работу, когда очередь в польской нотации исчерпана
		if calc.resultQueue.Length() == 0 {
			return operands.Pop().(float64), nil
		}
	}
}

// Применение операции на двух операндах
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

// Создание обратной польской нотации
// https://habr.com/ru/post/100869/?ysclid=kzo91k0d9f
func (calc *calculator) makeNotation(expression string) error {
	operand := make([]rune, 0, 20)
	for _, char := range expression {
		// Если символ не оператор, то накапливаем символы в операнд
		if !isOperator(char) {
			operand = append(operand, char)
			continue
		}
		// Пушим операнд в очередь нотации
		err := calc.pushOperand(string(operand))
		if err != nil {
			return err
		}
		// Обнуляем операнд
		operand = make([]rune, 0, 20)
		// Обрабатываем оператор
		calc.handleOperator(operator(char))
	}
	// Не забываем пушить послений операнд
	calc.pushOperand(string(operand))
	// Очищаем временный стек операторов в очередь нотации
	for calc.tempStack.Head() != nil {
		calc.resultQueue.Push(calc.tempStack.Pop())
	}
	return nil
}

// Преобразование строки операнда в float и
// загрузка в очередь нотации
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

// Обработка операторв
func (calc *calculator) handleOperator(op operator) {
	currentPriority := getPriority(op)
	head := calc.tempStack.Head()
	headPriority := -1
	if head != nil {
		headPriority = getPriority(head.(operator))
	}
	// Отдельная обработка скобок
	if op == '(' {
		calc.tempStack.Push(op)
		return
	}
	if op == ')' {
		// Очищение стека до открывающей скобочки
		for calc.tempStack.Head().(operator) != '(' {
			calc.resultQueue.Push(calc.tempStack.Pop())
		}
		calc.tempStack.Pop()
		return
	}
	// Если приоритет данной операции ниже лежащей в стеке,
	// то лежащий в стеке отправляется в очередь, иначе на вершине стека
	// оказывается данный оператор
	if currentPriority < headPriority {
		calc.resultQueue.Push(calc.tempStack.Pop())
	}
	calc.tempStack.Push(op)
}

// Получение приоритета операции
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

// Проверка является ли символ оператором
func isOperator(char rune) bool {
	return char == '+' ||
		char == '-' ||
		char == '*' ||
		char == '/' ||
		char == '(' ||
		char == ')'
}

// Очищение все вайтспейсов
func clearWhitespaces(str string) string {
	runes := make([]rune, 0, len(str))

	for _, char := range str {
		if !unicode.IsSpace(char) {
			runes = append(runes, char)
		}
	}
	return string(runes)
}
