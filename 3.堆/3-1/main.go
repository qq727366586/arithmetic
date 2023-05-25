package main

import "fmt"

// left child = 2*i+1
func leftChild(i int) int {
	return 2*i + 1
}

// right child = 2*i+2
func rightChild(i int) int {
	return 2*i + 2
}

// parent = (i-1)/2
func parent(i int) int {
	return (i - 1) / 2
}

// small heap insert
// 时间复杂度：O(logn)
func heapInsert(arr []int, index int) {
	for arr[index] < arr[parent(index)] {
		arr[index], arr[parent(index)] = arr[parent(index)], arr[index]
		index = parent(index)
	}
}

// heapify
// 时间复杂度 O(logn)
func heapify(arr []int, index int) {
	left, right := leftChild(index), rightChild(index)
	smallest := index

	if left < len(arr) && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != index {
		arr[index], arr[smallest] = arr[smallest], arr[index]
		heapify(arr, smallest)
	}
}

// zh: 按照从小到大排序堆
// en: sort heap from small to large
func heapSort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr[:i], 0)
	}
}

// zh: 获取堆的最小值， 并保持最小堆
// en: get the min value of heap, and keep the heap
func getMinAndKeepHeap(arr []int) int {
	min := arr[0]
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	heapify(arr, 0)
	return min
}

// zh: 在一个几乎有序的数组 每个元素移动距离不超过k， 请排序
// en: sort an almost sorted array, each element move distance not more than k
func sortKSortedArray(arr []int, k int) {
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
		if i == k {
			arr = arr[1:]
			sortKSortedArray(arr, k)
		}
	}
}

func main() {
	arr := []int{1, 4, 3, 2, 5, 6, 12, 4}
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	fmt.Print(arr)
	getMinAndKeepHeap(arr)
	fmt.Print(arr)
	heapSort(arr)
	fmt.Print(arr)
	fmt.Println("====")
	arr = []int{3, 4, 1, 2, 5}
	sortKSortedArray(arr, 2)
	fmt.Print(arr)

}
