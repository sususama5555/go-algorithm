package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */

func threeOrders(root *TreeNode) [][]int {
	// write code here
	res := make([][]int, 3)
	for i := 0; i < 3; i++ {
		res[i] = make([]int, 0)
	}
	printTree(root, res)
	return res
}

func printTree(root *TreeNode, res [][]int) {
	if root == nil {
		return
	}
	res[0] = append(res[0], root.Val)
	printTree(root.Left, res)
	res[1] = append(res[1], root.Val)
	printTree(root.Right, res)
	res[2] = append(res[2], root.Val)
}
