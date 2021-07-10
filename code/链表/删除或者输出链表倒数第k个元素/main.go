package main

type nodeList struct {
	val int
	Next *nodeList
}


/**
	思路1
	遍历链表，获取长度l
	再次遍历链表，第k个坐标为l-1-k
	时间复杂度O(3n),也就是O(n)
	空间复杂度O(3n),也就是O(n)
*/
func delK(head *nodeList,n int)*nodeList{
	var (
		l int
		sl []*nodeList
		slC []*nodeList
	)
	for head != nil{
		sl = append(sl,head)
		head = head.Next
		l++
	}

	//单独处理一个节点和删除1的问题
	if l ==1 {
		return nil
	}

	for i,v := range sl{
		if i == l-n{
			continue
			//如果需要输出第k个元素就可以输出了
		}
		slC = append(slC,v)

	}

	for i,v:= range slC{
		if i < len(slC)-1{
			v.Next = slC[i+1]

		}
	}
	if	slC != nil{
		return slC[0]
	}
	return nil
}

/**
	快慢指针想起始位置一起出发，快指针移动到快k的位置
	快指针比慢指针快k+1步，当快指针到的尾部的时候，慢指针正好是要删除的或者输出的元素、
	为什么是k+1步，因为只有这样，当快指针到达尾部的时候，慢指针才是指向要删除元素的上一个元素
	时间复杂度O(n)
	空间复杂度O(1)
*/
func delFRun(head *nodeList,n int)*nodeList{
	//构造虚拟头节点，避免头节点单独处理
	dummyNode := new(nodeList)
	dummyNode.Next = head

	var slow *nodeList
	cur := dummyNode
	//fmt.Printf("%p\n",cur) //0xc000010260
	//fmt.Printf("%p\n",dummyNode)//0xc000010260

	i:=1
	for head != nil{
		//slow先不走，让快指针到达k+1的位置
		if i>=n{
			slow = cur//要删除元素的上一个元素
			cur = cur.Next//慢指针移动
		}
		i++
		head = head.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
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
	//fmt.Println(*delK(n1,1))
	delFRun(n1,2)
}