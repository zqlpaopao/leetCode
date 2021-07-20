package main

import "fmt"

/**
题目： 有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。
注意，不是必须有这些素因子，而是必须不包含其他的素因子。
例如，前几个数按顺序应该是 1，3，5，7，9，15，21。

*/

/**
	以后的数一定是3，5，7的倍数 而且最后的商是3，5，7
	而得到的都是3，5，7的倍数乘积得来的
*/
func getKthMagicNumber(k int) int {
	num := make([]int, k)
	num[0] = 1
	i3,i5,i7 :=0,0,0
	for i:=1;i<k;i++ {
		res := min(num[i3]*3,min(num[i5]*5, num[i7]*7))
		if res == num[i3]*3 {
			i3++
		}
		if res == num[i5]*5 {
			i5++
		}
		if res == num[i7]*7 {
			i7++
		}
		num[i]= res
	}
	return num[k-1]
}

func min(a,b int)int {
	if a <b {
		return a
	}
	return b
}

func main(){
	fmt.Println(getKthMagicNumber(5))
}
