package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList
// @param pHead ListNode类
// @return ListNode类
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	cur := pHead
	var pre *ListNode
	var temp *ListNode
	for cur.Next != nil {
		temp = cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func ReverseListSimple(pHead *ListNode) *ListNode {
	// write code here
	var newHead *ListNode
	for pHead != nil {
		pHead.Next, newHead, pHead = newHead, pHead, pHead.Next
	}
	return newHead
}
