package doublylinkedlist

import (
	//"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	//assert.NotNil(t, NewList())
	list := NewList()
	list.PushFront("first")
	list.PushFront("second")
	list.PushFront("third")
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	fmt.Printf("length: %d\n", list.Len())
}
