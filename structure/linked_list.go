package structure

import "fmt"

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func (l *LinkedList) LinkedListFromArray(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		node := InitLinkedListNode(array[i])
		l.InsertNodeAtHead(node)
	}
}

func (l *LinkedList) GetNode(Head *LinkedListNode, pos int) *LinkedListNode {
	if pos != -1 {
		p := 0
		ptr := Head
		for p < pos {
			ptr = ptr.Next
			p++
		}
		return ptr
	}
	return nil
}

func (l *LinkedList) DisplayLinkedList() {
	temp := l.Head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func Length(Head *LinkedListNode) int {
	temp := Head
	length := 0
	for temp != nil {
		length++
		temp = temp.Next
	}
	return length
}

func DisplayLinkedListWithForwardArrow(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
}
func DisplayLinkedListWithForwardArrowLoop(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		}
	}
}
