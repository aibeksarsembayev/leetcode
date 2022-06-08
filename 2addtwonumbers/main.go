package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Convert linked list to array
	numArr1 := converListToIntArr(l1)
	numArr2 := converListToIntArr(l2)
	// Sum both arrays
	sumArr := sumArrays(numArr1, numArr2)
	// convert array of sum ints to linked list
	result := convertArrToList(sumArr)
	return result
}

func convertArrToList(sumArr []int) *ListNode {
	var result *ListNode
	for i := len(sumArr) - 1; i > -1; i-- {
		result = addNode(result, sumArr[i])
	}
	return result
}

func addNode(head *ListNode, v int) *ListNode {
	l := &ListNode{Val: v}
	if head == nil {
		return l
	} else {
		l.Next = head
		return l
	}
}

func sumArrays(n1, n2 []int) []int {
	var sumArr []int
	if len(n1) < len(n2) {
		n1, n2 = n2, n1
	}
	var upTen int
	for index, value := range n1 {
		if len(n2) > index {
			v := n2[index] + value + upTen
			upTen = 0
			if v > 9 {
				upTen = v / 10
				v %= 10
			}
			sumArr = append(sumArr, v)
		} else {
			v := value + upTen
			upTen = 0
			if v > 9 {
				upTen = v / 10
				v %= 10
			}
			sumArr = append(sumArr, v)
		}
		if upTen != 0 && index == len(n1)-1 {
			sumArr = append(sumArr, 1)
		}
	}
	return sumArr
}

func converListToIntArr(l *ListNode) []int {
	var intArr []int
	counter := 0
	for l != nil {
		intArr = append(intArr, l.Val)
		l = l.Next
		counter++
	}
	return intArr
}
