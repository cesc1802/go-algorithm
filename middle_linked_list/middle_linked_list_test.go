package middle_linked_list

import (
	"fmt"
	"strings"
	"testing"

	"algorithm/structure"
)

func TestLookingForMiddleNode(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}
	for i, input := range inputs {
		linkedList := new(structure.LinkedList)
		linkedList.LinkedListFromArray(input)
		fmt.Printf("%d.\tLinked list: ", i+1)
		structure.DisplayLinkedListWithForwardArrow(linkedList.Head)
		fmt.Printf("\n\tMiddle of linked list is: %d\n", LookingForMiddleNode(linkedList.Head).Data)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
