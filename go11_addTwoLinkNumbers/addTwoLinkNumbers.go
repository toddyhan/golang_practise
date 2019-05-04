package main

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
	pHead := ListNode{}
	point := &pHead
	var carry int

	for l1 != nil && l2 != nil {
		sum := carry + l1.Val + l2.Val
		point.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		point = point.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := carry + l1.Val
		point.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		point = point.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := carry + l2.Val
		point.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		point = point.Next
		l2 = l2.Next
	}

	if carry != 0 {
		point.Next = &ListNode{carry, nil}
	}
	return pHead.Next
}

//击败100%的答案，区别在于在处理l1或l2长出来的部分时，直接用l1或l2节点而不再分配新节点。
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var head ListNode
//	now := &head
//	var m int
//	for l1 != nil && l2 != nil {
//		var tmp *ListNode
//		tmp = new(ListNode)
//		//val = (l1.Val + l2.Val)
//		//m, v = val/10, val%10 //m为进位,v为当前数字
//		tmp.Val = (l1.Val + l2.Val + m) % 10
//		m = (l1.Val + l2.Val + m) / 10
//
//		now.Next = tmp
//
//		now = now.Next
//		l1 = l1.Next
//		l2 = l2.Next
//	}
//	//
//	if l1 != nil {
//		now.Next = l1
//		tmpVal := (l1.Val + m)/10
//		l1.Val = (l1.Val + m)%10
//		m = tmpVal
//		for m > 0 {
//			if l1.Next == nil {
//				var tmp *ListNode
//				tmp = new(ListNode)
//				l1.Next = tmp
//			}
//			l1 = l1.Next
//
//			tmpVal := (l1.Val + m)/10
//			l1.Val = (l1.Val + m)%10
//			m = tmpVal
//		}
//	}
//	if l2 != nil {
//		now.Next = l2
//		tmpVal := (l2.Val + m)/10
//		l2.Val = (l2.Val + m)%10
//		m = tmpVal
//		for m > 0 {
//			if l2.Next == nil {
//				var tmp *ListNode
//				tmp = new(ListNode)
//				l2.Next = tmp
//			}
//			l2 = l2.Next
//
//			tmpVal := (l2.Val + m)/10
//			l2.Val = (l2.Val + m)%10
//			m = tmpVal
//		}
//	}
//	if m > 0 {
//		var tmp *ListNode
//		tmp = new(ListNode)
//		tmp.Val = 1
//		now.Next = tmp
//	}
//	return head.Next
//}
func main() {

}
