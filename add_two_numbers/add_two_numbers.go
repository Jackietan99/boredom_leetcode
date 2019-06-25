package add_two_number

import (
	"fmt"
)

/*
	You are given two non-empty linked lists
	representing two non-negative integers.
	The digits are stored in reverse order and
	each of their nodes contain a single digit.
	Add the two numbers and return it as a linked list.
	You may assume the two numbers do not contain
	any leading zero, except the number 0 itself.

	Example:

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	node := new(ListNode)
	sum := 0
	v1 := 0
	v2 := 0
	if l1 != nil {
		v1 = l1.Val
	}
	if l2 != nil {
		v2 = l2.Val
	}
	sum += v1
	sum += v2
	next := new(ListNode)
	if sum/10 > 0 {
		if l1.Next == nil && l2.Next == nil {
			return &ListNode{sum % 10, nil}
		}
		if l1.Next == nil {
			l2.Next.Val++
			next = addTwoNumbers(nil, l2.Next)
		} else if l2.Next == nil {
			l1.Next.Val++
			next = addTwoNumbers(l1.Next, nil)
		} else {
			l1.Next.Val++
			next = addTwoNumbers(l1.Next, l2.Next)
		}
		node.Val = sum % 10
	} else {
		if l1.Next == nil && l2.Next != nil {
			next = addTwoNumbers(nil, l2.Next)
		} else if l2.Next == nil && l1.Next != nil {
			next = addTwoNumbers(l1.Next, nil)
		} else if l1.Next != nil && l2.Next != nil {
			next = addTwoNumbers(l1.Next, l2.Next)
		} else {
		}
		node.Val = sum
		//next = addTwoNumbers(l1.Next, l2.Next)
		//node.Val = sum
	}
	node.Next = next
	return node
}

func main() {
	////l1
	//l1 := new(ListNode)
	//l1.Val=1
	//next1 := new(ListNode)
	//next1.Val = 4
	//next2 := new(ListNode)
	//next2.Val =3
	//l1.Next = next1
	//next1.Next = next2
	//
	////PrintNodes(l1)
	//
	//l2 := new(ListNode)
	//l2.Val=5
	//next3 := new(ListNode)
	//next3.Val = 6
	//next4 := new(ListNode)
	//next4.Val =4
	//l2.Next = next3
	//next3.Next = next4
	////PrintNodes(l2)

	//l1
	l1 := new(ListNode)
	l1.Val = 1
	next1 := new(ListNode)
	next1.Val = 4
	next2 := new(ListNode)
	next2.Val = 3
	l1.Next = next1
	next1.Next = next2

	//PrintNodes(l1)

	l2 := new(ListNode)
	l2.Val = 5
	next3 := new(ListNode)
	next3.Val = 6
	next4 := new(ListNode)
	next4.Val = 4
	l2.Next = next3
	next3.Next = next4

	numbers := addTwoNumbers(l1, l2)

	PrintNodes(numbers)

}

func PrintNodes(n *ListNode) {
	fmt.Println(n.Val)
	if n.Next != nil {
		PrintNodes(n.Next)
	}
}
