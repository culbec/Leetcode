/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    var curr *ListNode = head
    var prev *ListNode = nil

    // Iterating the list on each node.
    for curr != nil {
        // Computing the double of the current node value
        twice := curr.Val * 2

        // Verifying if the double value is a digit
        if twice < 10 {
            curr.Val = twice
        } else if prev != nil {
            // If there exists a previous node, we carry the remaining digit to it
            curr.Val = twice % 10
            prev.Val = prev.Val + 1
        } else {
            // If there is no existing node, we create a new head with the remaining digit
            head = &ListNode{
                Val: 1,
                Next: curr,
            }
            curr.Val = twice % 10
        }

        // Updating the pointers
        prev = curr
        curr = curr.Next
    }

    return head
}
