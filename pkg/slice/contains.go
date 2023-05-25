package slice

func ContainsFunc[T any](src []T, dst T, equal equalFunc[T]) bool {
	for _, v := range src {
		if equal(v, dst) {
			return true
		}
	}
	return false
}

// Contains 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, dst, func(src, dst T) bool {
		return src == dst
	})
}

// ContainsAny 判断 src里面是否存在 dst中任何一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, ok := srcMap[v]; ok {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap(src)
	for _, v := range dst {
		if _, ok := srcMap[v]; !ok {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
func ContainsAllFunc[T comparable](src, dst []T, equal equalFunc[T]) bool {
	for _, v := range dst {
		if !ContainsFunc(src, v, equal) {
			return false
		}
	}
	return true
}
