package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	arr := tree2Arr(root)
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left] + arr[right] == k {
			return true
		} else if arr[left] + arr[right] > k {
			right --
		} else {
			left ++
		}
	}
	return false
}

func tree2Arr(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	var stack []*TreeNode
	stack = make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		c := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if c.Right != nil {
			stack = append(stack, c.Right)
		}
		if c.Left == nil && c.Right == nil {
			res = append(res, c.Val)
		} else {
			stack = append(stack, &TreeNode{Val: c.Val})
		}
		if c.Left != nil {
			stack = append(stack, c.Left)
		}

	}
	return res
}

func main() {
	findTarget(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, 3)
}
