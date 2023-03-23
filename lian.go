package main

import "fmt"

//字节面试题：给定奇偶链表，排序
//如1->8->3->6->5->4->7->2
//奇数位是递增，偶数位递减，排序

//定义链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortEvenAndOddLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//翻转奇偶链表
	Odd := head       //奇数节点的起点
	Even := head.Next //偶数节点的起点
	oddHead := Odd
	evenHead := Even
	for Even != nil && Even.Next != nil {
		Odd.Next = Even.Next
		Odd = Odd.Next
		Even.Next = Odd.Next
		Even = Even.Next
	}
	Odd.Next = nil
	//将偶数的链表进行翻转
	p := reverseLinkedList(evenHead)
	//合并两个排序的链表
	res := mergeSortEvenAndOdd(oddHead, p)
	return res
}

//翻转链表
func reverseLinkedList(p *ListNode) *ListNode {
	var pre *ListNode
	for p != nil {
		tmp := p.Next
		p.Next = pre
		pre = p
		p = tmp
	}
	return pre
}

//合并两个有序链表
func mergeSortEvenAndOdd(p *ListNode, q *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	cur := dummy
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			cur = cur.Next
			p = p.Next
		} else {
			cur.Next = q
			cur = cur.Next
			q = q.Next
		}
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return dummy.Next
}

//创建链表
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{-1, nil}
	head := &ListNode{nums[0], nil}
	dummy.Next = head
	for i := 1; i < len(nums); i++ {
		head.Next = &ListNode{nums[i], nil}
		head = head.Next
	}
	return dummy.Next
}

//打印链表
func printLinkedList(head *ListNode) {
	fmt.Println("------------------------------------------------------------")
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	nums := []int{1, 8, 3, 6, 5, 4, 7, 2}
	head := createLinkedList(nums)
	printLinkedList(head)
	head2 := sortEvenAndOddLinkedList(head)
	printLinkedList(head2)
}