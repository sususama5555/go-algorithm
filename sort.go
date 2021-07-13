package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func BubbleSort(arr []int) []int {
	// write code here
	for i, _ := range arr {
		for j, _ := range arr {
			if j <= i {
				continue
			}
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
