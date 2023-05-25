package main

import "fmt"

// 1. 最简单的线性查找 o(n)
func linerSearch(key int, data []int) (index int) {
	index = -1
	for k, v := range data {
		if v == key {
			index = k
		}
	}
	return
}

// 2. 最简单的线性查找 o(n) 支持泛型
func linerSearchV2[T comparable](key T, data []T) (index int) {
	index = -1
	for k, v := range data {
		if v == key {
			return k
		}
	}
	return
}

type student struct {
	name string
	age  int
}

func (s *student) ageEquals(stu *student) bool {
	return s.age == stu.age
}

func (s *student) this() *student {
	return s
}

type ageEquals interface {
	this() *student
	ageEquals(*student) bool
}

// 2. 最简单的线性查找 o(n2) 支持泛型
func linerSearchV3[T ageEquals](key T, data []T) (index int) {
	index = -1
	for k, v := range data {
		if v.ageEquals(key.this()) {
			return k
		}
	}
	return
}

func main() {
	// 1.
	res := linerSearch(1, []int{2, 3, 4, 5, 6, 1})
	fmt.Println(res)

	fmt.Println("---------------")

	// 2.
	res = linerSearchV2(1, []int{2, 3, 4, 5, 6, 1})
	fmt.Println(res)

	fmt.Println("---------------")

	// 3.
	res = linerSearchV3(&student{}, []*student{{age: 1}, {age: 10}})
	fmt.Println(res)
}
