package main

import (
	"fmt"
	"sort"
)

/**
	两数之和
	示例 1：

	输入：nums = [2,7,11,15], target = 9
	输出：[0,1]
	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
	示例 2：

	输入：nums = [3,2,4], target = 6
	输出：[1,2]
	示例 3：

	输入：nums = [3,3], target = 6
	输出：[0,1]

	遇到技术之和要首先想到map来做
*/
func twoSum(nums []int,target int)(index []int){
	var (
		key  = make(map[int]int,len(nums))
	)
	if len(nums) < 2 {
		return
	}

	for i ,v := range nums{
		key[v] = i
	}

	for i , v := range nums{
		if vl, ok := key[target - v];ok && vl != i{
			index = append(index,i,vl)
			//return
		}
	}
	return
}

/**
	三数之和
	给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
	使得a + b + c = 0，请你找出所有和为 0 且不重复的三元组。

	注意：答案中不可以包含重复的三元组。

	示例 1：

	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]

	解析：
	通过读题，我们知道其实就是求
	a + b + c = 0

	1⃣️、暴力破解，通过三层loop，时间复杂度O(N)

	2⃣️、set 加 2 loop
	c = -(a+b)
	set 时间复杂度是O(1),2loop是O(N^2)
	3⃣️、排序加双指针，时间复杂度是O(nlog(N)) O(n^2)
*/

func threeSump(nums []int,target int)(index [][]int){
	if len(nums) < 3{
		return
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2;i++{
		n1 := nums[i]
		//如果当前是大于0的，后边一定是大于0的
		if n1 > 0{
			break
		}
		//去重
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		l,r := i +1,len(nums)-1
		for l < r{
			n2,n3 := nums[l],nums[r]
			if n1 +n2 +n3 == 0{
				index = append(index,[]int{n1,n2,n3})
				//去重
				for l < r &&nums[l] == n2{
					l ++
				}
				//去重
				for l < r && nums[r] == n3{
					r --
				}
			}else if n1 + n2 +n3 < 0{
				l++
			}else{
				r--
			}
		}
	}

	return
}

/**
	四数只和
	示例 1：

	输入：nums = [1,0,-1,0,-2,2], target = 0
	输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	示例 2：

	输入：nums = [], target = 0
	输出：[]

	和三数之和是一样的 排序(O(nlogN)) + O(n^3)
*/
func fourSum(num []int)(index [][]int){
	if len(num) < 4{
		return
	}
	sort.Ints(num)
	for i := 0;i < len(num)-3;i++{
		if num[i] > 0{
			break
		}
		//去重
		if i > 0 && num[i] == num[i-1]{
			continue
		}

		for j := i+1;j < len(num)-2;j++{
			if j > i+1 && num[j] == num[j-1]{
				continue
			}
			l ,r := j+1,len(num)-1
			for l < r {
				n1,n2,n3,n4 := num[i],num[j],num[l],num[r]
				if n1+n2+n3+n4 == 0{
					index = append(index,[]int{n1,n2,n3,n4})
					for l < r && n3 == num[l]{
						l++
					}

					for l < r && n4 == num[r]{
						r--
					}
					l++
					r--
				}else if  n1+n2+n3+n4 < 0{
					l++
				}else{
					r--
				}
			}

		}
	}
	return
}



func main(){

	nums := []int{2,7,9,15}
	target := 9
	fmt.Println(twoSum(nums,target))

	fmt.Println("三数之和")
	num :=  []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSump(num,0))
	fmt.Println(fourSum(num))
}