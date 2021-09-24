package functional

func ClosureFactorial() func(num int) int {
	cache := make(map[int]*int)
	return func(num int) int {
		if cache[num] != nil {
			return *cache[num]
		} else if cache[num-1] != nil {
			result := *cache[num-1] * num
			cache[num] = &result
			return result
		} else {
			copy := num
			result := 1
			for copy > 0 {
				result *= copy
				copy--
			}
			cache[num] = &result
			return result
		}
	}
}
