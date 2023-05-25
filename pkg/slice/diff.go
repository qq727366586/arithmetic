package slice

// DiffSet 差集，已去重 只支持 comparable 类型 返回顺序会打乱
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap(src)
	for _, v := range dst {
		delete(srcMap, v)
	}
	var ret = make([]T, 0, len(srcMap))
	for k := range srcMap {
		ret = append(ret, k)
	}
	return ret
}

// DiffSetFunc 差集，已去重 返回顺序会打乱
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, val, equal) {
			ret = append(ret, val)
		}
	}
	return deduplicateFunc[T](ret, equal)
}
