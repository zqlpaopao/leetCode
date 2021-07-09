package main

import (
	"fmt"
	"os"
	"time"
)



/**
	tempHead = &ListNode{0,nil}
	tempHead.Next = head

	pre,cur,nex = tempHead,head,head.Next

	pre = 0-1-2-3-4-5-6-7
	cur = 1-2-3-4-5-6-7
	nex = 2-3-4-5-6-7

	for cur != nil && cur.Next != nil{

		cur.Next = nex.Next
		cur = 1-3-4-5-6-7

		nex.Next = cur
		nex = 2-1-3-4-5-6-7

		pre.Nex = nex
		pre = 0-2-1-3-4-5-6-7


		此时pre是取址，已经赋值给tempHead
		tempHead = 0-2-1-3-4-5-6-7

		为下次循环做准备
		pre = cur
		pre = 1-3-4-5-6-7
		cur = cur.next
		cur = 3-4-5-6-7
		if cur.Next == nil{
			break
		}
		nex = cur.Next
		nex = 4-5-6-7

		再次循环
		cur = 1-4-5-6-7
		nex = 4-3-5-6-7
		pre = 2-1-4-3-5-6-7
	}
*/

func recoverList(head *nodeList)*nodeList{
	if (head == nil) || (head.Next == nil) {
		return head
	}
	tmpHead := &nodeList{0, nil}
	tmpHead.Next = head
	pre, cur, nex := tmpHead, head, head.Next
	fmt.Println("pre",pre)
	fmt.Println("cur",cur)
	fmt.Println("nex",nex)
	for cur != nil && nex != nil {
		cur.Next = nex.Next
		fmt.Println()
		//fmt.Println("pre",pre)
		fmt.Println("cur",cur)
		fmt.Println("cur-next",cur.Next)
		//fmt.Println("nex",nex)
		nex.Next = cur
		fmt.Println()
		fmt.Println("nex",nex)
		fmt.Println("nex-next",nex.Next)
		pre.Next = nex
		fmt.Println()
		fmt.Println("pre",pre)
		fmt.Println("pre-next",pre.Next)
		fmt.Println("tmpHead",tmpHead)
		fmt.Println("tmpHead-next",tmpHead.Next)

		pre = cur
		fmt.Println()
		fmt.Println("pre",pre)
		fmt.Println("pre-next",pre.Next)
		fmt.Println("tmpHead",tmpHead)
		fmt.Println("tmpHead-next",tmpHead.Next)

		cur = cur.Next
		fmt.Println()
		fmt.Println("cur",cur)
		fmt.Println("cur-next",cur.Next)
		if cur == nil {
			break
		}
		nex = cur.Next
		fmt.Println()
		fmt.Println("nex",nex)
		fmt.Println("nex-next",nex.Next)
		fmt.Println()
		fmt.Println(tmpHead)
		fmt.Println(tmpHead.Next)
		os.Exit(3)
	}
	return tmpHead.Next
}

type nodeList struct {
	val int
	Next *nodeList
}

func main(){
	var t1 ,t2  time.Time

	t1 = time.Now().Add(60 * time.Second)
	t2 = time.Now()
	fmt.Println(t1.Sub(t2))

	if t1.Sub(t2)< 61 * time.Second{
		fmt.Println(2333)
	}

	fmt.Println( t1.Sub(t2))
	fmt.Println( t2.Sub(t1))
	if t2.Sub(t1)> 3 *time.Second{
		fmt.Println(23333)
	}

os.Exit(3)



	n1 := new(nodeList)
	n1.val = 1

	n2 := new(nodeList)
	n2.val = 2

	n3 := new(nodeList)
	n3.val = 3
	n4 := new(nodeList)
	n4.val = 4

	n5 := new(nodeList)
	n5.val = 5

	n6 := new(nodeList)
	n6.val = 6

	n7 := new(nodeList)
	n7.val = 7

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	pre := recoverList(n1)

	fmt.Printf("%#v\n",pre)
	fmt.Printf("%#v\n",pre.Next)
	fmt.Printf("%#v\n",pre.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next.Next.Next)
	fmt.Printf("%#v\n",pre.Next.Next.Next.Next.Next.Next)
	//fmt.Printf("%#v\n",recoverList(n1).next.next.next)
	//fmt.Printf("%#v\n",recoverList(n1).next.next.next.next)
	//fmt.Printf("%#v\n",recoverList(n1).next.next.next.next.next)
	//fmt.Printf("%#v\n",recoverList(n1).next.next.next.next.next.next)
}