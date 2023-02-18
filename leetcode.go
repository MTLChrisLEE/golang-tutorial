package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	// //Finding Sum
	// l1Value := 0
	// l1Unit := 1
	// for l1 != nil {
	// 	l1Value += l1.Val * l1Unit
	// 	l1Unit *= 10
	// 	l1 = l1.Next
	// }

	// l2Value := 0
	// l2Unit := 1
	// for l2 != nil {
	// 	l2Value += l2.Val * l2Unit
	// 	l2Unit *= 10
	// 	l2 = l2.Next
	// }
	// sum := l1Value + l2Value

	// //Building new linked list
	// currentIndex := head
	// for sum >= 0 {
	// 	currentIndex.Next = &ListNode{Val: sum % 10}
	// 	sum /= 10
	// 	currentIndex = currentIndex.Next
	// }

	// return head.Next

	sum := 0
	multiplier := 1
	for l1 != nil && l2 != nil {
		sum += l1.Val*multiplier + l2.Val*multiplier
		l1 = l1.Next
		l2 = l2.Next
		multiplier *= 10
	}

	if sum == 0 {
		return &ListNode{Val: 0}
	}

	//Building new linked list
	currentIndex := head
	for sum > 0 {
		currentIndex.Next = new(ListNode)
		currentIndex.Next.Val = sum % 10
		sum /= 10
		currentIndex = currentIndex.Next
	}

	if l1 == nil {
		currentIndex.Next = l2
	}

	if l2 == nil {
		currentIndex.Next = l1

	}

	return head.Next
}

func main() {
	list1 := new(ListNode)
	list1.Val = 1
	list1.Next = &ListNode{Val: 0}
	list1.Next.Next = &ListNode{Val: 0}

	list1.Next.Next.Next = &ListNode{Val: 0}
	list1.Next.Next.Next.Next = &ListNode{Val: 0}
	list1.Next.Next.Next.Next = &ListNode{Val: 0}
	list1.Next.Next.Next.Next.Next = &ListNode{Val: 1}

	list2 := new(ListNode)
	list2.Val = 5
	list2.Next = &ListNode{Val: 6}
	list2.Next.Next = &ListNode{Val: 4}

	addTwoNumbers(list1, list2)
}
