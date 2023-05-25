package slice

import "arithmetic/pkg"

// Max 返回最大值
func Max[T pkg.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] > res {
			res = ts[i]
		}
	}
	return res
}

// Min min
func Min[T pkg.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res
}

// Sum sum
func Sum[T pkg.RealNumber](ts []T) T {
	var res T
	for _, n := range ts {
		res += n
	}
	return res
}
