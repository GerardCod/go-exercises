package sorting

type MergeSort struct{}

// Sort is the implementation of Sorting interface
func (ms *MergeSort) Sort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	middle := int(len(data) / 2)
	left := data[0:middle]
	right := data[middle:]

	leftSorted := ms.Sort(left)
	rightSorted := ms.Sort(right)

	return ms.stitch(leftSorted, rightSorted)
}

func (ms *MergeSort) stitch(left, right []int) []int {
	var result []int
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
