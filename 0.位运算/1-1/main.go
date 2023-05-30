package main

import "fmt"

// zh:打印一个数的二进制形式
// en:print a number's binary form
func printBinaryForm(data int32) {
	for i := 31; i >= 0; i-- {
		fmt.Print(data >> i & 1)
	}
}

// zh:不用中间变量交换两个数
// en:swap two numbers without a temp variable
func swapTwoNumbersWithoutTempVariable(a, b int32) (int32, int32) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// zh:找到出现次数奇数次的数
// en:find the number that appears odd times
func findTheNumberThatAppearsOddTimes(data []int32) int32 {
	var tmp int32
	for _, v := range data {
		tmp ^= v
	}
	return tmp
}

// zh:找到出现次数K次的数
// en:find the number that appears K times
func findTheNumberThatAppearsKTimes(data []int32, k int32) int32 {
	tmpArr := make([]int32, 32, 32)
	for _, v := range data {
		for j := 31; j >= 0; j-- {
			tmpArr[31-j] += v >> j & 1
		}
	}
	var res int32
	for i := 0; i <= 31; i++ {
		if tmpArr[i] == 0 {
			continue
		}
		if tmpArr[i]%k == 0 {
			res |= 1 << (31 - i)
			fmt.Println(i, tmpArr[i], res)
		}
	}
	return res
}

func main() {
	printBinaryForm(10)
	fmt.Println()
	res := findTheNumberThatAppearsKTimes([]int32{1, 1, 1, 2, 2, 2, 3, 3, 3, 7, 7}, 2)
	fmt.Println(res)
}
