package bin

// SliceDistinct 切片去重
func SliceDistinct[T comparable](vs ...T) []T {
	result := []T{}
	for i := range vs {
		if SliceContain(result, vs[i]) {
			continue
		}
		result = append(result, vs[i])
	}
	return result
}

// SliceContain 判断切片中是否包含子内容
func SliceContain[T comparable](vs []T, v T) bool {
	for i := range vs {
		if vs[i] == v {
			return true
		}
	}
	return false
}

// SliceMaxNumber 获取切片中最大数
func SliceMaxNumber[T NumberGenericity](vs []T) T {
	var max T
	for _, v := range vs {
		if max < v {
			max = v
		}
	}
	return max
}

// SliceMinNumber 获取切片中最小数
func SliceMinNumber[T NumberGenericity](vs []T) T {
	var min T
	for i, v := range vs {
		if i == 0 {
			min = v
		}
		if min > v {
			min = v
		}
	}
	return min
}
