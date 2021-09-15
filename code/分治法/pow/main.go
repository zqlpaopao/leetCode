package main

import (
	"log"
)

func main(){
	log.Println(forPow(2,-2))
	//log.Println(forSum(2,-1))
	//log.Println(myPow(2,-2))
	log.Println(quickPow(2,-2))
}

//-- ----------------------------
//--> @Description
//--> @Param
//--> @return
//-- ----------------------------
func forPow(i ,j int)(res float64){
	res = 1
	f := j
	if j <0 {
		j = -j
	}
	for a := 0; a < j;a ++{
		if j > 0 {
			res *= float64(i)
		}
	}
	if f > 0{
		return res
	}else {
		return 1 /res
	}

}


func myPow(x float64, n int) float64 {
	if x == 0{
		return 0
	}else if n > 0{
		return pre(x,n)
	}else{
		return 1/pre(x,n)
	}
}

func pre(x float64, n int) float64{
	//终止条件 go向下取整 1 /2 = 0
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}
	y := pre(x,n/2)
	if n % 2 == 0{
		return y * y
	}else {
		return x * y * y
	}

}


//-- ----------------------------
//--> @Description
//--> @Param
//--> @return
//-- ----------------------------
func quickPow(x,y int)(ret float64){
	ret = 1
	f := y
	//如果是<0的
	if y < 0{
		y = -y
	}

	for y != 0{
		//与运算不为0的说明不可以被2整除
		if (y & 1) != 0{
			//多出来个x 要乘
			ret = ret*float64(x)
		}
		x = x*x
		//y是指数不断减小
		y >>= 1
	}

	if f > 0{
		return ret
	}else {
		return 1/ret
	}
}