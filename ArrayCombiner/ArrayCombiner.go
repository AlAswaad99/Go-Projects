package main

import (
	"fmt"
	"strconv"
)

func main() {
	Combiner()
}

func Combiner() {
	var combined []string

	list1, list2 := [4]int{1, 2, 3, 4}, [4]string{"a", "b", "c", "d"}

	for i := 0; i < len(list1); i++ {
		combined = append(combined, strconv.Itoa(list1[i]))
		combined = append(combined, list2[i])

	}

	fmt.Printf("Original Lists were %v and %v \nThe Combined list will be %v", list1, list2, combined)

}
