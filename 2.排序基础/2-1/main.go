package main

import (
	"fmt"
)

// 选择排序 o(n2)
func selectionSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		minIndex := i
		for k := i + 1; k < len(data); k++ {
			if data[minIndex] > data[k] {
				minIndex = k
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}

// 插入排序
func InsertionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for k := i + 1; k > 0 && data[k] < data[k-1]; k-- {
			data[k], data[k-1] = data[k-1], data[k]
		}
	}
	return data
}

// 插入排序V2 -从前面插入
func InsertionSortV2(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		key := data[i+1]
		k := i + 1
		for ; k > 0 && key < data[k-1]; k-- {
			data[k] = data[k-1]
		}
		data[k] = key
	}
	return data
}

// 插入排序V3  -从后面插入
func InsertionSortV3(data []int) []int {
	for i := len(data) - 1; i > 0; i-- {
		k := i - 1
		key := data[k]
		for ; k < len(data)-1 && key > data[k+1]; k++ {
			data[k] = data[k+1]
		}
		data[k] = key
	}
	return data
}

// 冒泡排序
func BubbleSort(data []int) []int {
	for i := len(data) - 1; i > 0; i-- {
		for k := 0; k+1 <= i; k++ {
			if data[k] > data[k+1] {
				data[k], data[k+1] = data[k+1], data[k]
			}
		}
	}
	return data
}

func main() {
	res := selectionSort([]int{10, 9, 1, 23, 3})
	fmt.Println(res)
	fmt.Println("=========================")
	res = InsertionSort([]int{10, 9, 1, 23, 3, -1, 99, 12, 123, 4, 4, 5, 5, 5})
	fmt.Println(res)
	fmt.Println("=========================")
	res = InsertionSortV2([]int{10, 9, 1, 23, 3, -1, 99, 12, 123, 4, 4, 5, 5, 5})
	fmt.Println(res)
	fmt.Println("=========================")
	res = InsertionSortV3([]int{10, 9, 1, 23, 3, -1, 99, 12, 123, 4, 4, 5, 5, 5})
	fmt.Println(res)
	fmt.Println("=========================")
	res = BubbleSort([]int{10, 9, 1, 23, 3, -1, 99, 12, 123, 4, 4, 5, 5, 5})
	fmt.Println(res)

}
