package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rec(head **ListNode, tail *ListNode) bool {
	if tail == nil {
		return true
	}

	var r bool = rec(head, (*tail).Next)
	var ok bool = (**head).Val == (*tail).Val

	if *head == tail && (*head).Next == tail {
		return ok
	}

	*head = (*head).Next
	return ok && r
}

func palindromeSLL(head *ListNode) bool {
	return rec(&head, head.Next)
}

func createSLL(arr []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = nil

	for _, e := range arr {
		if head == nil {
			head = &ListNode{
				Val:  e,
				Next: nil,
			}

			curr = head
		} else {
			node := &ListNode{
				Val:  e,
				Next: nil,
			}

			curr.Next = node
			curr = node
		}
	}
	return head
}

func main() {
	head := createSLL([]int{1, 0, 0})
	fmt.Println(palindromeSLL(head))
}
