package main

import "fmt"

type nodeList struct {
	val int
	Next *nodeList
}

/**
	hashSet的解法
	1、遍历链表A，将链表地址存储在hashSet中
	2、遍历链表B，如果链表B存在hashSet中，则此时就是交点

	时间复杂度O(n)，空间复杂度O(n)
*/
func isCycleSet(a,b *nodeList)*nodeList{
	var pMap = make(map[string]struct{})
	for a.Next != nil{
		pMap[fmt.Sprintf("%p",a)] = struct{}{}
		a = a.Next
	}

	for b.Next != nil{
		if _ ,ok := pMap[fmt.Sprintf("%p",b)];ok{
			return b
		}
		b= b.Next
	}

	return nil
}


/**
	快慢指针
	1、指针 pA 指向 A 链表，指针 pB 指向 B 链表，依次往后遍历
	2、如果 pA 到了末尾，则 pA = headB 继续遍历
	3、如果 pB 到了末尾，则 pB = headA 继续遍历
	4、比较长的链表指针指向较短链表head时，长度差就消除了
	5、如此，只需要将最短链表遍历两次即可找到位置
	总的思想就是 我吹过你吹过的晚风
	时间复杂度o(m+n) 也就是o(n)
	空间复杂度O(1)
*/
func isCycleRun(a,b *nodeList)*nodeList{
	if a==nil || b == nil{
		return nil
	}
	headA,headB  := a,b
	for a != b{
		if a == nil{
			a = headB
		}else{
			a=a.Next
		}

		if b == nil{
			b = headA
		}else{
			b = b.Next
		}
	}

	return a
}






func main(){
	//构造链表
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


	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6

	//链表2
	e1 := new(nodeList)
	e1.val = 9

	e2 := new(nodeList)
	e2.val = 5

	e3 := new(nodeList)
	e3.val = 6


	e1.Next = e2
	e2.Next = e3
	e3.Next = n5
	//fmt.Println(e3.Next == n4.Next)

	fmt.Println(*isCycleSet(n1,e1))
	fmt.Println(*isCycleRun(n1,e1))
}