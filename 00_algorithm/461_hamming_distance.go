package main

import "fmt"

func hammingDistance(x int, y int) int {
	xor, res := x ^ y, 0
	for ; xor > 0; xor &= xor - 1 {
		res++
	}
	return res
}

func hammingDistance1(x int, y int) int {
	i :=x^y    //x和y异或,得到一个新的数(异或相同为0,不同为1,此时咱们需要统计1的个数)
	count :=0  //定义数量的初始值为0
	for i!=0 { //只要i不为0,那就继续循环
		if (i&1)==1 { //如果i和1相与,值为1的话就count++   这一步我要特别注意，go要继续判断是否等于1
			count++
		}
		i = i>>1//i右移一位
	}
	return count
}




func main()  {
	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance1(1, 2))
}