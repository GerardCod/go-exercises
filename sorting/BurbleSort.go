package sorting

type BurbleSort struct{}

func (bs *BurbleSort) Sort(data []int) []int {

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j+1] < data[j] {
				swap(data, j, j+1)
			}
		}
	}

	return data
}
