package main

import (
	"unicode"
)

//import "fmt"

func RemoveEven(input []int) (result []int) {
	//result := []int
	for i := range input{
		if input[i]%2 != 0 {
			result = append(result, input[i])
		}
	}

	return
}

func PowerGenerator(i int) (func() int) {
	n := 1
	return func() (int) {
		n *= i
		return n
	}
}

func DifferentWordsCount(str string) (result int) {
	result = 0
	word := ""
	set := make(map[string]bool)
	str = str + " "

	for _, sym := range str {
		if unicode.IsLetter(sym) {
			word += string(unicode.ToLower(sym))
		} else if word != "" {
			if !set[word] {
				result++
			}
			set[word] = true
			word = ""
		}
	}
	return
}

/*
func main() {
	/*
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]
	//*/
	/*
	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8
	//*/
	/*
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2
	//*/
//}
