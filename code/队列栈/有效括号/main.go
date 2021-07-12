package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
	时间复杂度O(n)
	空间复杂度o(n)
 */

func isValid(s []byte) bool {
	Mp := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}

	stack := []byte{}

	for _ ,v := range s{
		if v1,ok := Mp[v];!ok{
			stack = append(stack,v)
		}else {
			//stack < 1 防止全是右边的
			//和栈顶的比较，如果遇到的第一个右括号和栈顶的不匹配，肯定就不符合了
			if len(stack) <1 || v1 != stack[len(stack)-1]{
				return false

			}else{
				//匹配的话移除元素
				stack = stack[0:len(stack)-1]
			}
		}

	}
	//防止只有(( 左括号这种
	if len(stack) > 0{
		return false
	}

	return true
}

func main(){
	str := "((({{}}[][])))"
	fmt.Println(isValid([]byte(str)))
}
