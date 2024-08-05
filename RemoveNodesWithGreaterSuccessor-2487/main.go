package main

type ListNode struct {
  Val int
  Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
  if head == nil {
    return nil
  }

  curr := head
  next := removeNodes(curr.Next)

  curr.Next = next
  if next == nil || curr.Val >= next.Val {
    return curr
  }

  return next
}
