package main

import (
	. "leetCode/basic"
)

func reverseList(head *ListNode) *ListNode {
	var fristList *ListNode
	secondList := head
	for secondList != nil {
		temp := secondList.Next
		secondList.Next = fristList
		fristList = secondList
		secondList = temp
	}
	return fristList
}
