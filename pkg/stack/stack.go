package stack

// Структура элемента односвязного списка
type Node struct {
	Data interface{}
	next *Node
}

// Структура стека
type Stack struct {
	head *Node
}

// Получение экземпляра стека
func GetStack() Stack {
	return Stack{
		head: nil,
	}
}

// Добавление в стек
func (list *Stack) Push(data interface{}) {
	node := Node{
		Data: data,
		next: list.head,
	}
	list.head = &node
}

// Выталкивание из стека
func (list *Stack) Pop() interface{} {
	if list.head == nil {
		return nil
	}
	defer func() {
		list.head = list.head.next
	}()
	return list.head.Data
}

// Получение верхнего элемента стека без выталкивания
func (list *Stack) Head() interface{} {
	if list.head == nil {
		return nil
	}
	return list.head.Data
}

// Обнуление стека
func (list *Stack) Reset() {
	list.head = nil
}

// Длинна стека
func (list *Stack) Length() uint {
	var i uint
	temp := list.head

	i = 0
	for temp != nil {
		temp = temp.next
		i++
	}
	return i
}
