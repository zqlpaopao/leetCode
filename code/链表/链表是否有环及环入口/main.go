package main

import (
	"fmt"
)

/**
  给定链表，判断是否有环及环入口位置
*/

type nodeList struct {
	val int
	Next *nodeList
}

//-- ----------------------------
//--> @Description 链表是否有环 及环的入口
//--> @Param
//--> @return
//-- ----------------------------
func isRingSet(head *nodeList)*nodeList{
	var Map = make(map[string]struct{})
	cur := head
	for cur != nil && cur.Next != nil{
		if _ ,ok := Map[fmt.Sprintf("%p",cur)];ok{
			return cur
		}
		Map[fmt.Sprintf("%p",cur)] = struct{}{}
		cur = cur.Next
	}
	return nil
}

//-- ----------------------------
//--> @Description 链表是否有环
//--> @Param
//--> @return
//-- ----------------------------
func isRingCycle(head *nodeList)*nodeList{
	slow , fast := head,head
	for  slow != nil  && fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			return slow
		}
	}
	return nil
}

/**
快慢指针判断环的入口

fast = 2,slow = 1 当第一次相遇时候，快指针从起点走，
慢指针正常前行，当第二次相遇的时候就是环的入口

*/
func loopCycle(head *nodeList)*nodeList{
	slow,fast := head,head
	for slow != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}

	fast = head
	for  slow != nil && fast != nil && fast.Next != nil{
		fast = fast.Next
		slow = slow.Next
		if fast == slow{
			return slow
		}
	}
	return nil
}




func main(){

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
	n6.val = 7

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n4

	//fmt.Println(*isRingSet(n1))
	//fmt.Println(*isRingCycle(n1))
	fmt.Println(*loopCycle(n1))

}