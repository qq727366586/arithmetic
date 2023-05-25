package main

import "fmt"

// 二分法 （有序数组）
func dichotomy(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid] < key {
			l = mid + 1
		} else if arr[mid] > key {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 二分法 （有序数组） 找到大于等于某个数的 最左侧下标
func dichotomyPracticeV1(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	index := -1
	for l <= r {
		mid := l + (r-l)>>1
		if arr[mid] >= key {
			index = mid
			r = mid - 1
		} else if arr[mid] < key {
			l = mid + 1
		}
	}
	return index
}

// 二分法 （有序数组） 找到小于等于某个数的 最右侧下标
func dichotomyPracticeV2(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	index := -1
	for l <= r {
		mid := l + (r-l)>>1
		if arr[mid] >= key {
			r = mid - 1
		} else if arr[mid] <= key {
			l = mid + 1
			index = mid
		}
	}
	return index
}

// ***** 二分法 （有序数组） 局部最小值 相邻两个数不相等 len > 1
func dichotomyPracticeV3(arr []int) int {
	if arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	l := 1
	r := len(arr) - 2
	for l < r {
		mid := l + (r-l)>>1
		if arr[mid] > arr[mid-1] {
			r = mid - 1
		} else if arr[mid] > arr[mid+1] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}

func main() {
	arr := []int{1, 2, 2, 2, 3, 4, 5, 6, 10, 11, 15}
	index := dichotomy(arr, 11)
	fmt.Println(index)

	fmt.Println("----------------------------")

	index = dichotomyPracticeV1(arr, 21)
	fmt.Println(index)

	fmt.Println("----------------------------")

	index = dichotomyPracticeV2(arr, 3)
	fmt.Println(index)

	fmt.Println("----------------------------")

	arr = []int{10, 2, 3, 4, 5, 6, 51, 6, 10, 11, 15}
	index = dichotomyPracticeV3(arr)
	fmt.Println(index)

}
