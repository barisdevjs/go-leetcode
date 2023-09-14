package main

import "fmt"

func DeleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		fmt.Println("C", current)
		if current.Next != nil && current.Val == current.Next.Val {
			fmt.Println(current.Val, current.Next.Val)
			current.Next = current.Next.Next
		} else {
			current = current.Next
			fmt.Println("SSS", current, current.Next)

		}
	}
	return head
}

func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = DeleteDuplicates2(head.Next)
		return head
	} else {
		head.Next = DeleteDuplicates2(head.Next)
		return head
	}
}

func DeleteDuplicates3(head *ListNode) *ListNode {

	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return head
}
