package slice

import "fmt"

func ExampleDelete() {
	res, _ := Delete[int]([]int{1, 2, 3, 4}, 2)
	fmt.Println(res)
	_, err := Delete[int]([]int{1, 2, 3, 4}, -1)
	fmt.Println(err)
	// Output:
	// [1 2 4]
	// ekit: 下标超出范围，长度 4, 下标 -1
}
