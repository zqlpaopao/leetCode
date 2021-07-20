package main

import "fmt"

/**
给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

*/

func maxSlidingWindow(nums []int, k int) []int {
	//双端队列
	var queue []int
	//要输出的元素队列
	var rtn []int


	for f := 0; f < len(nums); f++ {
		//如果当前元素比单调队列的最后一个大，那么一直比较，算出最大值
		for len(queue) > 0 && nums[f] > nums[queue[len(queue)-1]] {
			fmt.Println(queue)
			queue = queue[:len(queue)-1]
		}

		// 如果当前值比单调队列的最后一个值小，直接入队
		//入队的是下标
		if len(queue) == 0 || nums[f] <= nums[queue[len(queue)-1]] {
			queue = append(queue, f)
		}

		//当k到达指定窗口的时候，开始往输出队列加入第一个元素
		if f >= k - 1 {
			rtn = append(rtn, nums[queue[0]])
			//移除单调队列的值，因为下一次活动窗口往后移动了
			//当单调队列的下标不在滑动窗口内的时候，移除元素
			if f - k + 1 == queue[0] {
				queue = queue[1:]
			}
		}
	}

	return rtn

}

func is (nums []int,k int)[]int{
	var queue , rnt []int

	for f := 0;f <len(nums);f++{
		for len(queue) > 0 && nums[f] > queue[len(queue)-1]{
			queue = queue[:len(queue)-1]
		}

		if len(queue) == 0 || nums[f] <= queue[len(queue)-1]{
			queue = append(queue,f)
		}

		if f >= k - 1 {
			rnt = append(rnt,nums[queue[0]])
			if f- k + 1 == queue[0]{
				queue = queue[1:]
			}
		}
	}
	return rnt
}




func main(){
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 2
	fmt.Println(maxSlidingWindow(nums,k))
	fmt.Println(is(nums,k))
}