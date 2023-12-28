package structure

import "fmt"

type LinkedList struct {
	head *LinkedListNode
}

func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) LinkedListFromArray(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		node := InitLinkedListNode(array[i])
		l.InsertNodeAtHead(node)
	}
}

func (l *LinkedList) GetNode(head *LinkedListNode, pos int) *LinkedListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.next
			p++
		}
		return ptr
	}
	return nil
}

func (l *LinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func Length(head *LinkedListNode) int {
	temp := head
	length := 0
	for temp != nil {
		length++
		temp = temp.next
	}
	return length
}

func DisplayLinkedListWithForwardArrow(l *LinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
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
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" → ")
		}
	}
}
