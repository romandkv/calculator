package stack


// Структура очереди
type Queue struct {
	head *Node
	tail *Node
}

// Получение экземпляра очереди
func GetQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
	}
}


// Добавление в очередь элемента
func (list *Queue) Push(data interface{}) {
	node := Node{
		Data: data,
		next: nil,
	}
	if list.head == nil && list.tail == nil {
		list.head = &node
		list.tail = &node
		return
	}
	list.tail.next = &node
	list.tail = &node
}

// Выталкивание из очереди элемента
func (list *Queue) Pop() interface{} {
	if list.head == nil {
		return nil
	}
	defer func() {
		if list.head == list.tail {
			list.Reset()
			return
		}
		list.head = list.head.next
	}()
	return list.head.Data
}

// Обнуление очереди
func (list *Queue) Reset() {
	list.head = nil
	list.tail = nil
}

// Длинна очереди
func (list *Queue) Length() uint {
	var i uint
	temp := list.head

	i = 0
	for temp != nil {
		temp = temp.next
		i++
	}
	return i
}

// Получение первого элемента очереди,
// без выталкивания
func (list *Queue) Head() interface{} {
	if list.tail == nil {
		return nil
	}
	return list.tail.Data
}
