package algodaily_test

import (
	"fmt"
	"testing"

	"github.com/GerardCod/exercises/algodaily"
)

type TestCase struct {
	Name           string
	Input          interface{}
	ExpectedOutput interface{}
}

func TestReverseString(t *testing.T) {
	tests := []TestCase{
		{
			Name:           "It should return an string reversed",
			Input:          "Hello World",
			ExpectedOutput: "dlroW olleH",
		},
		{
			Name:           "It should return an empty string",
			Input:          "",
			ExpectedOutput: "",
		},
	}

	for _, value := range tests {
		fmt.Println(value.Name)
		result := algodaily.ReverseString(fmt.Sprintf("%v", value.Input))
		expected := fmt.Sprintf("%v", value.ExpectedOutput)
		if result != expected {
			fmt.Println("======================= ERROR =======================")
			fmt.Printf("Expected %v, got %v\n", expected, result)
		} else {
			fmt.Println("======================= PASSED =======================")
		}
	}
}
