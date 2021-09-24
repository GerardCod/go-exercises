package algodaily

import (
	"fmt"
	"strconv"
)

func FizzBuzz(number string) string {
	fizz := ""
	limit, err := strconv.Atoi(number)

	if err != nil {
		fmt.Println("The arg is not a number")
	}

	for i := 1; i <= limit; i++ {
		if i%3 == 0 {
			fizz += "fizz"
		} else if i%5 == 0 {
			fizz += "buzz"
		} else if i%3 == 0 && i%5 == 0 {
			fizz += "fizzbuzz"
		}
	}

	return fizz
}
