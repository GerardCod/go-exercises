package sorting

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	middle := int(len(arr) / 2)
	left := arr[0:middle]
	right := arr[middle:]

	leftSorted := MergeSort(left)
	rightSorted := MergeSort(right)

	return stitch(leftSorted, rightSorted)
}

func stitch(left, right []int) []int {
	var result []int
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
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
