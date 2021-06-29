package main

import "fmt"



/**
	输入 1->2->3->4->null
	输出 5->4->3->2->1_.null

	思路 把指针指向前一个元素，后边的不断向前移动，并改变方向
	位置调换次数	pre					cur					whole
	0			nil					1->2->3->4->5		1->2->3->4->5
	1			1->nil				2->-3>->4->5		2->3->4->5->1->nil
	2			2->1->nil			3->4->5				3->4->5->2->1->nil
	3			3->2->1->nil		4->5				4->5->3->2->1->nil
	4			4->3->2->1->nil		5					5->4->3->2->1->nil
	pre 就是cur的最前边的那位（pre=cur）
	cur就是当前元素后面链表元素（cur=cur.Next）
	cur.Next肯定是接pre（curl.Next=pre）原因是cur赋值给了pre

*/
//反转链表双指针迭代one
func reversList(head *ListNode)*ListNode{
	cur := head
	var pre *ListNode = nil
	for cur != nil{
		pre ,cur ,cur.Next = cur,cur.Next,pre//关键点
	}
	return pre
}

/**************************************************************************/
//反转链表递归实现
func reversListTwo(head *ListNode)*ListNode{
	 if head == nil || head.Next == nil{
	 	return head
	 }
	 pre := reversListTwo(head.Next)
	 head.Next.Next = head
	 head.Next = nil
	 return pre
}


type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reversList(head)
	fmt.Printf("%#v\n",pre)
	fmt.Printf("%#v\n",pre.Next)
	fmt.Printf("%#v\n",pre.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next.Next.Next)

	fmt.Println("递归实现")
	pre = reversListTwo(head)
	fmt.Printf("%#v\n",pre)
	fmt.Printf("%#v\n",pre.Next)
	fmt.Printf("%#v\n",pre.Next.Val)
	//fmt.Printf("%#v\n",pre.Next.Next.Next)
	//fmt.Printf("%#v\n",pre.Next.Next.Next.Next)
	//fmt.Printf("%#v\n",pre.Next.Next.Next.Next.Next)
}

