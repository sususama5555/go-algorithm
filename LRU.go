package main

import "fmt"

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func main() {
	operators := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	fmt.Println(LRU(operators, 3))
}

func LRU(operators [][]int, k int) []int {
	// write code here
	res := make([]int, 0, len(operators))
	key := make([]int, k)
	value := make([]int, k)
	for _, item := range operators {
		switch item[0] {
		case 1:
			if len(key) == k {
				key = key[1:]
				value = value[1:]
			}
			key = append(key, item[1])
			value = append(value, item[2])
		case 2:
			index := -1
			for i := range key {
				if item[1] == key[i] {
					index = i
					break
				}
			}
			if index == -1 {
				res = append(res, -1)
			} else {
				res = append(res, value[index])
				if index < k - 1 {
					// 获取的 key 非最近最常使用
					key = append(key[:index], append(key[index+1:], key[index])...)
					value = append(value[:index], append(value[index+1:], value[index])...)
				}
			}

		}
	}
	return res
}
