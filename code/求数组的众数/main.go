package main

import (
	"fmt"
	"sort"
)

/**

示例1：

输入：[3,2,3]
输出：3
示2：

输入：[2,2,1,1,1,2,2]
输出：2


进阶：

尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

*/

//-- ----------------------------
//--> @Description map 时间复杂度O(N),空间复杂度O(N)
//--> @Param
//--> @return
//-- ----------------------------
func majorityElementMap(nums []int)int{
	var count = make(map[int]int,len(nums))
	for _ ,v := range nums{
		count[v] ++
		if count[v] > len(nums)>>1{
			return v
		}
	}
	return 0
}

//-- ----------------------------
//--> @Description 排序选随机数
//--> @Param
//--> @return
//-- ----------------------------
func majorityElementSortAmdRand(nums []int)int{
	sort.Ints(nums)
	return nums[len(nums)>>1]
}


func main(){
	nums := []int{1,2,1,2,2}
	fmt.Println(majorityElementMap(nums))
	fmt.Println(majorityElementSortAmdRand(nums))
}