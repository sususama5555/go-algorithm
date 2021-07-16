package main

import (
	"fmt"
	"reflect"
)

/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */

func main() {
	//numbers := [3]int{3, 2, 4}
	//var numbers [3]int
	numbers := []int{3, 2, 4}
	//numbers := make([]int, 3)
	//var numbers = new([3]int)
	//numbers = append(numbers, 3)
	//numbers := [...]string{"a", "b", "c", "d"}

	test := numbers
	//test[0] = 99
	fmt.Println(reflect.TypeOf(numbers))
	fmt.Println(numbers)
	fmt.Println(test)
	for i, j := range numbers {
		fmt.Println(i, j)
	}
	fmt.Println(twoSum(numbers, 6))
}

func twoSum(numbers []int, target int) []int {
	hashMap := make(map[int]int, len(numbers))
	for index, value := range numbers {
		tmp := target - value
		if i, ok := hashMap[tmp]; ok {
			return []int{i + 1, index + 1}
		}
		hashMap[value] = index
	}
	return nil
}
