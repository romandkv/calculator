package stack

type Node struct {
	Data interface{}
	next *Node
}

type Stack struct {
	head *Node
}

func GetStack() Stack {
	return Stack{
		head: nil,
	}
}

func (list *Stack) Push(data interface{}) {
	node := Node{
		Data: data,
		next: list.head,
	}
	list.head = &node
}

func (list *Stack) Pop() interface{} {
	if list.head == nil {
		return nil
	}
	defer func() {
		list.head = list.head.next
	}()
	return list.head.Data
}

func (list *Stack) Head() interface{} {
	if list.head == nil {
		return nil
	}
	return list.head.Data
}

func (list *Stack) Reset() {
	for list.head != nil {
		list.Pop()
	}
}

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
