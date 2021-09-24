package array

import (
	"errors"
	"fmt"
)

//ArrayList is the array implementation of List interface.
type ArrayList struct {
	List []int
	HeadIndex int
	TailIndex int
	Count int	
}

//NewArrayList returns a new instance of ArrayList.
func NewArrayList() *ArrayList {
	return &ArrayList{List: make([]int, 16),Count: 0, HeadIndex: 0, TailIndex: -1}
}

//IsEmpty returns if the list is empty
func (al *ArrayList) IsEmpty() bool {
	return al.Count == 0
}

//Size returns the size of the list.
func (al *ArrayList) Size() int {
	return al.Count
}

//First returns the first element of the list.
func (al *ArrayList) First() (int, error) {
	if al.IsEmpty() {
		return 0, errors.New("the list is empty")
	}
	return al.List[al.HeadIndex], nil
}

//Last returns the last element of the list
func (al *ArrayList) Last() (int, error) {
	if al.IsEmpty() {
		return 0, errors.New("the list is empty")
	}
	return al.List[al.TailIndex], nil
}

//Append inserts a new element at the end of the list.
func (al *ArrayList) Append(number int) {
	al.TailIndex++
	al.List[al.TailIndex] = number
  	al.Count++
}

//Find search an element and returns its value if the number is in the list otherwise returns 0.
func (al *ArrayList) Find(number int) (int, error) {
	if al.IsEmpty() {
		return 0, errors.New("the list is empty")
	}

	min := al.HeadIndex
	max := al.TailIndex
	middle := int((max + min) / 2)
	for min <= max {
		if number == al.List[middle] {
			return number, nil
		} else if number < al.List[middle] {
			max = middle - 1
		} else {
			min = middle + 1
		}
		middle = int((max + min) / 2)
	}
	return 0, nil
}

//IndexOf returns the index of the given number. Returns -1 if the number doesn't exist in the list.
func (al *ArrayList) IndexOf(number int) int {
	min := al.HeadIndex
	max := al.TailIndex
	middle := int((max + min) / 2)
	for min <= max {
		if number == al.List[middle] {
			return middle
		} else if number < al.List[middle] {
			max = middle - 1
		} else {
			min = middle + 1
		}
		middle = int((min + max) / 2)
	}
	return -1
}

//Shift adds a new element at the beginning of the list.
func (al *ArrayList) Shift(number int) {
  newList := make([]int, 16)
  newList[0] = number
  for idx, num := range al.List {
    newList[idx + 1] = num
  }
  al.List = newList
  al.Count++;
}

//Pop removes the last element of the list and returns its value.
func (al *ArrayList) Pop() (int, error) {
  if al.IsEmpty() {
    return 0, errors.New("the list is empty")
  }
  value := al.List[al.TailIndex]
  al.TailIndex--
  al.Count--
  return value, nil
}

//Unshift removes the first element of the list and retrieves its value.
func (al *ArrayList) Unshift() (int, error) {
  if al.IsEmpty() {
    return 0, errors.New("the list is empty")
  }
  value := al.List[al.HeadIndex]
  al.HeadIndex++
  al.Count--
  return value, nil
}

//Show prints all the list values by white space separated.
func (al *ArrayList) Show() error {
  if al.IsEmpty() {
    return errors.New("the list is empty")
  }

  for _, num := range al.List {
    fmt.Printf("%d ", num)
  }
  return nil
}