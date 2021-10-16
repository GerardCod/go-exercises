package sorting

type QuickSort struct{}

func (qs *QuickSort) Sort(data []int) []int {
	dataCopy := data
	qs.sortHelper(data, 0, len(data)-1)
	return dataCopy
}

func (qs *QuickSort) sortHelper(data []int, min, max int) {
	if min < max {
		pti := qs.partition(data, min, max)
		qs.sortHelper(data, 0, pti-1)
		qs.sortHelper(data, pti+1, max)
	}
}

func (qs *QuickSort) partition(data []int, min, max int) int {
	pivot := data[max]
	i := min - 1

	for j := min; j <= max; j++ {
		if data[j] < pivot {
			i++
			swap(data, i, j)
		}
	}

	swap(data, max, i+1)
	return (i + 1)
}
