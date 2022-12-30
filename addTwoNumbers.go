package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head *ListNode
	var temp *ListNode
	carry, sum := 0, 0

	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		node := ListNode{Val: sum % 10}
		carry = sum / 10

		if head == nil {
			temp = &node
			head = temp
		} else {
			temp.Next = &node
			temp = temp.Next
		}

	}

	if carry > 0 {
		temp.Next = &ListNode{Val: carry}
	}

	return head
}
