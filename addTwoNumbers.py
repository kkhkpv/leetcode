class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head: ListNode | None = None
        temp: ListNode | None = None
        carry = 0
        sum = 0
        while l1 or l2:
            sum = carry
            if l1:
                sum += l1.val
                l1 = l1.next
            if l2:
                sum += l2.val
                l2 = l2.next
            node = ListNode(sum%10)
            carry = sum // 10
            if temp:
                temp.next = node
                temp = temp.next
            else:
                temp = head = node
        if carry > 0:
            temp.next = ListNode(carry)
        return head
