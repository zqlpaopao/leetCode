package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
 输出数据流中第k大的元素
 例如给定[2,5,6,1,8,12,15,13]
 输出3 输出12
*/


/**
	第一种sort.Int排序，然后找到长度k-1的下标就是
	时间复杂度是快速排序O(nlogn)
*/
func sortAndPrint(arr []int,k int)int{
	if k >= len(arr){
		return 0
	}
	sort.Ints(arr)
	return arr[len(arr)-k]

}

/**
	构造k小顶堆，最后对顶的元素就是要输出的
*/


func main(){
fmt.Println(sortAndPrint([]int{3,6,2,9,4,1,6,13,12,16},4))



	fmt.Println("小顶堆")
	k := 4
	h := &IntHeap{3,6,2,9,4,1,6,13,12,16}
	heap.Init(h)
	i := h.Len()
	for i > k{
		heap.Pop(h)
		i = h.Len()
	}
	fmt.Println(heap.Pop(h))
}

/**************************** 小顶堆************************************/
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
//最小堆构造
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
//最大堆构造
//func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
