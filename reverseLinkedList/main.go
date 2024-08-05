package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(curr, headNew, currNew **ListNode) {
	if (*curr).Next == nil {
		*headNew = &ListNode{
			Val:  (*curr).Val,
			Next: nil,
		}
		*currNew = *headNew
		return
	}

	reverse(&(*curr).Next, headNew, currNew)

	(*currNew).Next = &ListNode{
		Val:  (*curr).Val,
		Next: nil,
	}
	*currNew = (*currNew).Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var curr *ListNode = head
	var headNew *ListNode = &ListNode{}
	var currNew *ListNode = &ListNode{}

	reverse(&curr, &headNew, &currNew)

	return headNew
}
