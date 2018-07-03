package leetcode

import (
	"fmt"
)

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

 // Definition for singly-linked list.
 type ListNode struct {
	 Val int
	 Next *ListNode
 }

 func genListNode(array []int) *ListNode {	
	var l, curr *ListNode
	for _, v := range array {
		n := &ListNode{Val: v}
		if l == nil {
			l = n
			curr = n
		} else {
			curr.Next = n
			curr = n
		}
	}
	return l
 }

 func printListNode(l *ListNode) string {
	str := "["
	for l != nil {
		str = fmt.Sprintf("%s %d", str, l.Val)
		l = l.Next
	}
	str = str + "]"
	return str
 }

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var output *ListNode
    var curr *ListNode
    var overflow int = 0
    for l1 != nil && l2 != nil {
		// fmt.Printf("/ %d \n", (l1.Val + l2.Val + overflow) / 10)
        node := &ListNode{Val:(l1.Val + l2.Val + overflow) % 10}
		overflow = (l1.Val + l2.Val + overflow) / 10
		// fmt.Printf("mod %d \n", (l1.Val + l2.Val + overflow) % 10)
		if output == nil {
            output = node
            curr = node
        } else {
            curr.Next = node
            curr = node
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    for l1 != nil {
		node := &ListNode{Val:(l1.Val + overflow) % 10}
		overflow = (l1.Val + overflow) / 10
        curr.Next = node
        curr = node
        l1 = l1.Next
    }
    for l2 != nil {
        node := &ListNode{Val:(l2.Val + overflow) % 10}
		overflow = (l2.Val + overflow) / 10
		curr.Next = node
        curr = node
        l2 = l2.Next
    }
    if overflow > 0 {
        node := &ListNode{Val: overflow}
        curr.Next = node
    }
    return output
}

func TestAddTwo () {
	fmt.Println("TestLongestSubstring")
	// l1 := genListNode([]int{2, 4, 3})
	// l2 := genListNode([]int{5, 6, 4})
	l1 := genListNode([]int{0})
	l2 := genListNode([]int{7, 3})
	fmt.Println(printListNode(l1))
	fmt.Println(printListNode(l2))
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(printListNode(l3))
}
