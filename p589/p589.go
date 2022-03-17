package p589

//https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if nil == root {
		return []int{}
	}
	ans := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		c := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, c.Val)
		for i := len(c.Children) - 1; i >= 0; i -- {
			stack = append(stack, c.Children[i])
		}
	}
	return ans
}
