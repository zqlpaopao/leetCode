# 1、数组

- 数组指定下表的时间复杂度是O(1),不置顶下表是O(n)
- 数组的插入最后一个是O(1)，如果不是，涉及数据搬迁，时间复杂度是O(n)
- 数组删除，如果是最后一个，时间复杂度是O(1),不是的话时间复杂度是O(n)

查找O(1)

插入O(n)

删除O(n)



# 2、链表

插入O(1)

删除O(1)

查找O(n)



# 2、链表

## 一、链表反转

## 二、链表元素两两交换



## 三、链表是否有环及环的位置
给定链表，判断链表是否有环及环的位置
如下

![image-20210706104315663](readme.assets/image-20210706104315663.png)

<font color=red size=5x>**第一种-时间判断**</font>

给懂时间1s或者0.5s，来判定是否循环完毕。

此方法会有误差，如果链表足够长，并且没有环，而超过时间就会认定是有环



<font color=red size=5x>**第二种-set查重**</font>

设置map集合，将每次循环的链表元素地址存储在map中，当出现在map中存在的情况时候，就是存在环，而且能找出==环的入口==

```go
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


func isRing(head *nodeList)*nodeList{
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


	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n4

	fmt.Println(*isRing(n1))

}
```

<font color=red size=5x>**第三种-快慢指针**</font>

![image-20210706140852681](readme.assets/image-20210706140852681.png)

==第一步==

![image-20210706140910943](readme.assets/image-20210706140910943.png)

指定快慢指针，slow步幅为1，fast步幅为2

==第1次移动==

![image-20210706141019772](readme.assets/image-20210706141019772.png)

==第二次移动==

![image-20210706141040453](readme.assets/image-20210706141040453.png)



==第三次移动==

![image-20210706141101223](readme.assets/image-20210706141101223.png)

此时fast已经进入第二圈

==第四次移动==

![image-20210706141118452](readme.assets/image-20210706141118452.png)

此时fast和slow相遇，说明存在环

==代码==

```go
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


	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n4

	//fmt.Println(*isRingSet(n1))
	fmt.Println(*isRingCycle(n1))

}
```



```go
slow &{2 0xc00008e200}
fast &{3 0xc00008e210}
slow &{3 0xc00008e210}
fast &{5 0xc00008e230}
{4 0xc00008e220}
```



<font color=red size=5x>**环入口位置的确定**</font>

==第一种==

我们可以用hashset的方式，既能确定是否有环，能确定环的入口

==第二种==

快慢指针

快指针fast和慢指针slow从头开始出发，快指针每次走2步，慢指针每次走1步，当第一次相遇的时候，快指针从头出发，每次走一步，再次相遇的时候就是环的入口

![image-20210709144019404](readme.assets/image-20210709144019404.png)

```go
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
```

我们来论证一下

![image-20210709144703903](readme.assets/image-20210709144703903.png)

假设头节点到环入口的举例为x，环入口到第一次相遇的举例为y，第一次相遇到环入口的位置为z，
则slow走的举例为x+y，而fast走的距离为x+y +n(y+z)，也就是fast在环中走了n圈（n>=0），因为
fast每次走2步，slow走一步，所以fast是slow的2倍

```go
2(x+y) = x+y +n(y+z)
x = n(y+z)-y
x=(n-1)(y+z)+z
```

y+z是环的长度

有两种情况：

n=1 时，x=z，此时我们将 fast 放到链表头，然后 fast 和 slow 每次走一步，相遇节点就是环的入口；
n>1 时，我们将 fast 放到链表头，当 fast 和 slow 相遇时，说明 slow 在环里转了 n-1 圈后又走了z步，等价于 n=1 的情况。



## 四、两个链表是否相交及交点

![image-20210710110803815](readme.assets/image-20210710110803815.png)

<font color=red size=5x>链表结构</font>

```go
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

	fmt.Println(*isCycleSet(n1,e1))
```



### HashSet

此种方法简单，但是时间复杂度和空间复杂度都是O(n)

思路

- 遍历链表A，将链表地址存储在hashSet中
- 遍历链表B，如果链表B存在hashSet中，则此时就是交点



==代码==

```go
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
```



```go
{5 0xc00008e230}
```



### 快慢指针

空间复杂度O(1)

时间复杂度O(n)



==如果两个链表相交，那么相交后的长度是相同的==

构造相同长度的链表

1. 指针 pA 指向 A 链表，指针 pB 指向 B 链表，依次往后遍历
2. 如果 pA 到了末尾，则 pA = headB 继续遍历
3. 如果 pB 到了末尾，则 pB = headA 继续遍历
4. 比较长的链表指针指向较短链表head时，长度差就消除了
5. 如此，只需要将最短链表遍历两次即可找到位置



![相交链表.png](readme.assets/e86e947c8b87ac723b9c858cd3834f9a93bcc6c5e884e41117ab803d205ef662-相交链表.png)



==代码==

```go
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
```



```go
{5 0xc00008e230}
```



## 五、删除or输出链表倒数的第k个元素

题目描述

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

 

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]



<font color=red size=5x>获取链表长度，遍历拼接</font>



```go
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

	return slC[0]
}

```



```go
1
&{2 0xc00008e200}
2
&{3 0xc00008e220}
3
&{5 0xc00008e230}
5
&{6 <nil>}
```





<font color=red size=5x>**快慢指针**</font>

- fast首先走n + 1步 ，为什么是n+1呢，因为只有这样同时移动的时候slow才能指向删除节点的上一个节点（方便做删除操作）
- fast和slow同时移动，之道fast指向末尾
- 删除slow指向的下一个节点，如图



![img](readme.assets/cc43daa8cbb755373ce4c5cd10c44066dc770a34a6d2913a52f8047cbf5e6e56-file_1559548337458.gif)





```go
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

```









































































