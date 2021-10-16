package sorting

type InsertionSort struct{}

func (is *InsertionSort) Sort(data []int) []int {
	dataCopy := data

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[i] {
				swap(data, i, j)
			}
		}
	}

	return dataCopy
}
