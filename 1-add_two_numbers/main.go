package main

import (
	"fmt"
)


type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1, node2 := l1, l2
	var root_node *ListNode = &ListNode{
		Val: 0,
		Next: nil,
	}
	current_node := root_node
	borrow := 0
    for node1 != nil || node2 != nil || borrow != 0{
		val1 := 0
		if (node1 != nil) {
			val1 = node1.Val
		}
		val2 := 0
		if (node2 != nil) {
			val2 = node2.Val
		}
		s := val1 + val2 + borrow
		borrow = s / 10
		new_node := &ListNode{
			Val: s%10,
			Next: nil,
		}
		current_node.Next = new_node
		current_node = new_node
		
		if (node1 != nil) {
			node1 = node1.Next
		}
		if (node2 != nil) {
			node2 = node2.Next
		}
    }
	return root_node.Next
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: nil,
		},
	}
	res_nodes := addTwoNumbers(&l1, &l2)
	fmt.Println(res_nodes)
	var current *ListNode = res_nodes
	for current != nil {
		fmt.Print(current.Val)
		current = current.Next
	}
}