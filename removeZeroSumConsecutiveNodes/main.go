package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func deleteNodes(fromNode *ListNode, toNode *ListNode) {
	if fromNode == toNode {
		return
	}

	deleteNodes(fromNode.Next, toNode)
	fromNode.Next = nil
}

func addNewNode(curr *ListNode, elem int) {
	toAdd := createNewNode(elem)
	curr.Next = toAdd
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	// Dummy node for storing the new SLL.
	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	newHead.Next = head

	sum := 0

	// Map that contains all the partial sums of the list
	partialSums := map[int]*ListNode{}
	partialSums[0] = newHead

	for curr := head; curr != nil; curr = curr.Next {
		sum += curr.Val

		// Checking if we already calculated this sum => we reached a case where we have a 0.
		// We need to remove the nodes in (node_of_sum, curr_node]
		if _, ok := partialSums[sum]; ok {
			prevNode := partialSums[sum]
			node := prevNode
			sumTemp := sum

			// Deleting all the nodes.
			for node = node.Next; node != curr; node = node.Next {
				sumTemp += node.Val
				delete(partialSums, sumTemp)
			}

			// Remaking the links.
			prevNode.Next = curr.Next
		} else {
			partialSums[sum] = curr
		}
	}
	return newHead.Next
}

func createSLL(arr []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = nil

	for _, elem := range arr {
		if head == nil {
			head = createNewNode(elem)
			curr = head
		} else {
			addNewNode(curr, elem)
			curr = curr.Next
		}
	}

	return head
}

func passSLL(head *ListNode) {
	curr := head

	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}

	fmt.Println()
}

func main() {
	fmt.Println(removeZeroSumSublists(nil))
	passSLL(removeZeroSumSublists(createSLL([]int{1, 2, -3, 3, 1})))
	passSLL(removeZeroSumSublists(createSLL([]int{1, 2, 3, -3, 4})))
	passSLL(removeZeroSumSublists(createSLL([]int{1, 2, 3, -3, -2})))
	passSLL(removeZeroSumSublists(createSLL([]int{-40, 40, 9, -2, 4})))
	passSLL(removeZeroSumSublists(createSLL([]int{24, -38, -38, -6, 19})))
	passSLL(removeZeroSumSublists(createSLL([]int{-11, -14, -40, -1, 30})))
	passSLL(removeZeroSumSublists(createSLL([]int{1, -36, 14, -27, -29})))
	passSLL(removeZeroSumSublists(createSLL([]int{30, -13, 32, -36, -1})))
	passSLL(removeZeroSumSublists(createSLL([]int{1, 3, 2, -3, -2, 5, 5, -5, 1})))
	passSLL(removeZeroSumSublists(createSLL([]int{5, -3, -4, 1, 6, -2, -5})))
}
