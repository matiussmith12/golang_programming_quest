package main

import (
	"fmt"
)

func main() {
	studentsCities := []string{"Mumbai", "Delhi", "Ahmedabad", "Mumbai", "Bangalore", "Delhi", "Kolkata", "Pune"}
	studentIDs := []int{11, 33, 22, 1, 33, 12, 25, 11}

	fmt.Println(removeDuplicateStr(studentsCities))
	fmt.Println(removeDuplicateInt(studentIDs))


}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		fmt.Println("item", item)
		if _, value := allKeys[item]; !value {

			fmt.Println("Value", value)
			allKeys[item] = true
			list = append(list, item)
			fmt.Println("LIST", list)
		}

	}
	return list
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)

		}

	}
	return list
}
