package sorting

type SelectionSort struct{}

// Sort
func (ss *SelectionSort) Sort(data []int) []int {
	dataCopy := data

	for i := 0; i < len(dataCopy)-1; i++ {
		minimumIdx := i

		for j := i + 1; j < len(dataCopy)-1; j++ {
			if dataCopy[j] < dataCopy[minimumIdx] {
				minimumIdx = j
			}
		}

		if minimumIdx != i {
			swap(dataCopy, i, minimumIdx)
		}
	}

	return dataCopy
}
