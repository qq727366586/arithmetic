package main

import (
	"fmt"
)

// 归并排序
func MergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) []int {
	tmp := make([]int, 0, r-l+1)

	p1, p2 := l, m+1

	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
			tmp = append(tmp, arr[p1])
			p1++
		} else {
			tmp = append(tmp, arr[p2])
			p2++
		}
	}
	for p1 <= m {
		tmp = append(tmp, arr[p1])
		p1++
	}
	for p2 <= r {
		tmp = append(tmp, arr[p2])
		p2++
	}
	for i, v := range tmp {
		arr[l+i] = v
	}
	return arr
}

func main() {
	arr := []int{2, 4, 3, 2, 5, 12, 3, 32, 12}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
