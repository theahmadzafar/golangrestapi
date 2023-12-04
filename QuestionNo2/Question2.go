package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value      string
	path       string
	dictionary map[string]int
}

var swaps_count int

func main() {

	dictionary := make(map[string]int)
	s := "aab"

	if len(s) > 500 && len(s) < 1 {

		fmt.Println("String Length is not valid please Enter a Valid String")
	} else {
		fmt.Println("Good job")

	}
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if dictionary[string(chars[i])] == 0 {
			dictionary[string(chars[i])] = 1
		} else {
			dictionary[string(chars[i])]++
		}

	}
	fmt.Println(dictionary)
	var parentNode Node

	parentNode.dictionary = dictionary
	getChild(parentNode)
	if swaps_count == 0 {
		fmt.Println("\"\"")
	} else {
		fmt.Printf("Total possible combinations are %s", strconv.Itoa(swaps_count))
	}

}

func getChild(parentNode Node) {
	for k, v := range parentNode.dictionary {
		if v > 0 {
			if parentNode.value != k {

				var childNode Node
				childNode.value = k
				childNode.path = parentNode.path + childNode.value
				childNode.dictionary = make(map[string]int)

				for kc, vc := range parentNode.dictionary {
					childNode.dictionary[kc] = vc
				}

				childNode.dictionary[childNode.value]--

				if childNode.dictionary[childNode.value] == 0 {
					delete(childNode.dictionary, childNode.value)
				}
				if len(childNode.dictionary) == 0 {
					fmt.Println(childNode.path)
					swaps_count++
				} else {

					getChild(childNode)
				}
			}

		}

	}
}
