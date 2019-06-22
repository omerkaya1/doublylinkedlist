package doublylinkedlist

/**
 * Doubly Linked List
 */
type ListImpl struct {
	head *Node
	tail *Node
	len  int
}

func NewList() *ListImpl {
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
	l.push(data, true)
}

func (l *ListImpl) PushBack(data interface{}) {
	l.push(data, false)
}

func (l *ListImpl) push(data interface{}, front bool) {
	if data == nil {
		return
	}
	node := &Node{Data: data, list: l}
	if l.tail == nil || l.head == nil {
		l.head, l.tail = node, node
		l.len++
		return
	}
	if front {
		l.head.prev, node.next, l.head = node, l.head, node
	} else {
		l.tail.next, node.prev, l.tail = node, l.tail, node
	}
	l.len++
	return
}

func (l *ListImpl) RemoveNode(n *Node) {
	if n.list != l {
		return
	}
	l.remove(n)
}

func (l *ListImpl) remove(n *Node) {
	if n.prev == nil {
		l.head = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
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
	Data interface{}
	prev *Node
	next *Node
	list *ListImpl
}

func (n Node) Value() interface{} {
	return n.Data
}

func (n Node) Next() *Node {
	return n.next
}

func (n Node) Prev() *Node {
	return n.prev
}
