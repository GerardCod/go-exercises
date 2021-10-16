package sorting

// Sorting is an interface that defines the behavior to sort a list of number
type Sorting interface {
	Sort(data []int) []int
}

func swap(data []int, left, right int) {
	aux := data[right]
	data[right] = data[left]
	data[left] = aux
}
