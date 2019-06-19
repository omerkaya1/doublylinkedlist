package doublylinkedlist

type DoublyLinkedList interface {
	Len() int
	First() *Node
	Last() *Node
	PushFront(interface{})
	PushBack(interface{})
	RemoveNode(*Node)
}

/**
 * Doubly Linked List
 */
type ListImpl struct {
	Head *Node
	Tail *Node
}

/**
 * Node
 */
type Node struct {
	Data   interface{}
	Parent *Node
	Child  *Node
}

func NewList() DoublyLinkedList {
	return &ListImpl{}
}

func (l ListImpl) Len() int {
	length := 0
	for temp := l.Head; temp != nil; length++ {
		temp = temp.Next()
	}
	return length
}

func (l ListImpl) First() *Node {
	return l.Head
}

func (l ListImpl) Last() *Node {
	return l.Tail
}

func (l *ListImpl) PushFront(data interface{}) {
	l.push(data)
}

func (l *ListImpl) PushBack(data interface{}) {
	l.pop(data)
}

func (l *ListImpl) push(data interface{}) {
	if data == nil {
		return
	}
	node := &Node{Data: data}
	if l.Head == nil {
		l.Head, l.Tail = node, node
	} else {
		l.Head.Parent, node.Child, l.Head = node, l.Head, node
	}
	return
}

func (l *ListImpl) pop(data interface{}) {
	if data == nil {
		return
	}
	node := &Node{Data: data}
	if l.Tail == nil {
		l.Head, l.Tail = node, node
	} else {
		l.Tail.Child, node.Parent, l.Tail = node, l.Tail, node
	}
	return
}

func (l *ListImpl) RemoveNode(n *Node) {
	switch n {
	case l.Head:
		l.Head, l.Head.Parent = n.Child, nil
		break
	case l.Tail:
		l.Tail, l.Tail.Child = n.Parent, nil
		break
	default:
		n.Parent.Child, n.Child.Parent = n.Child, n.Parent
		break
	}
}

func (n Node) Value() interface{} {
	return n.Data
}

func (n Node) Next() *Node {
	return n.Child
}

func (n Node) Prev() *Node {
	return n.Parent
}
