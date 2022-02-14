package stack

type Queue struct {
	head *Node
	tail *Node
}

func GetQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
	}
}

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

func (list *Queue) Reset() {
	list.head = nil
	list.tail = nil
}

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

func (list *Queue) Head() interface{} {
	if list.tail == nil {
		return nil
	}
	return list.tail.Data
}
