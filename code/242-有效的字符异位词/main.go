package main

import "fmt"

/**
	输入: s = "anagram", t = "nagaram"
	输出: true
	示例 2:

	输入: s = "rat", t = "car"
	输出: false
*/

func isAnagram1(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	var strMap = make(map[rune]int64)

	for _ ,v := range s {
		strMap[v] ++
	}

	for _ ,v := range t{
		if _, ok := strMap[v];ok{
			strMap[v] --
			if strMap[v] < 0{
				return false
			}
		}else{
			return false
		}
	}
	return true
}
/**
	第二种
	字符对应ascii码，字符最多26个
	相同的数-相同的（a）最后肯定可以落在相同的槽
	如果是异位字符串，肯定最后相同槽的值为0，不为0说明不符合条件
*/
func isAnagram2(s string, t string) bool {
	var cnt [26]int
	if !(len(s) == len(t)){
		return false
	}

	for i := range s{
		cnt[s[i]- 'a'] ++
		cnt[t[i]- 'a'] --
	}
	for _ ,v := range cnt{
		if v >0 {
			return false
		}
	}
	return true
}

func main(){
	s:= "anagram"
	t := "nagaram"
	fmt.Println(isAnagram1(s,t))
	fmt.Println(isAnagram2(s,t))
}
