package slice

// IntersectSet 取交集，只支持 comparable 类型
// 已去重
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap(src)
	var ret = make([]T, 0, len(src))
	// 交集小于等于两个集合中的任意一个
	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			ret = append(ret, val)
		}
	}
	return deduplicate[T](ret)
}

// IntersectSetFunc 支持任意类型
// 你应该优先使用 Intersect
// 已去重
func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, valSrc := range src {
		for _, valDst := range dst {
			if equal(valDst, valSrc) {
				ret = append(ret, valSrc)
				break
			}
		}
	}
	return deduplicateFunc[T](ret, equal)
}
