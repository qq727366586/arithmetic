package slice

import "fmt"

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("pkg/slice: 下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want, got string) error {
	return fmt.Errorf("pkg/slice: 类型转换失败，want:%s, got:%s", want, got)
}

// Delete 删除符合条件的元素
func Delete[Src any](src []Src, index int) ([]Src, error) {
	l := len(src)
	if index < 0 || index >= l {
		return nil, NewErrIndexOutOfRange(l, index)
	}
	j := 0
	for k, v := range src {
		if k != index {
			src[j] = v
			j++
		}
	}
	src = src[:j]
	return src, nil
}

// FilterDelete 删除符合条件的元素
func FilterDelete[Src any](src []Src, m func(idx int, src Src) bool) []Src {
	// 记录被删除的元素位置，也称空缺的位置
	emptyPos := 0
	for idx := range src {
		// 判断是否满足删除的条件
		if m(idx, src[idx]) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
