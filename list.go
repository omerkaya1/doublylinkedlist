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
	head *Node
	tail *Node
	len  int
}

func NewList() DoublyLinkedList {
	return &ListImpl{}
}

func (l *ListImpl) Len() int {
	return l.len
}

func (l *ListImpl) First() *Node {
	return l.head
}

func (l *ListImpl) Last() *Node {
	return l.tail
}

func (l *ListImpl) PushFront(data interface{}) {
	l.push(data)
}

func (l *ListImpl) PushBack(data interface{}) {
	l.pop(data)
}

func (l *ListImpl) RemoveNode(n *Node) {
	if n.list != l {
		return
	}
	l.remove(n)
}

func (l *ListImpl) push(data interface{}) {
	if data == nil {
		return
	}
	node := &Node{Data: data, list: l}
	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		l.head.parent, node.child, l.head = node, l.head, node
	}
	l.len++
	return
}

func (l *ListImpl) pop(data interface{}) {
	if data == nil {
		return
	}
	node := &Node{Data: data, list: l}
	if l.tail == nil {
		l.head, l.tail = node, node
	} else {
		l.tail.child, node.parent, l.tail = node, l.tail, node
	}
	l.len++
	return
}

func (l *ListImpl) remove(n *Node) {
	switch {
	case n == l.head && n != l.tail:
		n.child.parent, l.head = nil, n.child
		break

	case n != l.head && n == l.tail:
		n.parent.child, l.tail = nil, n.parent
		break

	case n == l.head && n == l.tail:
		break

	case n != l.head && n != l.tail:
		n.parent.child, n.child.parent = n.child, n.parent
		break

	default:
		panic("literally impossible situation")
	}
	n = nil
	// Decrement list length and ensure that when the last element
	// of the list gets removed, we go back to where we started
	if l.len--; l.len == 0 {
		l.head, l.tail = nil, nil
	}
	return
}

/**
 * Node
 */
type Node struct {
	Data   interface{}
	parent *Node
	child  *Node
	list   *ListImpl
}

func (n Node) Value() interface{} {
	return n.Data
}

func (n Node) Next() *Node {
	return n.child
}

func (n Node) Prev() *Node {
	return n.parent
}
