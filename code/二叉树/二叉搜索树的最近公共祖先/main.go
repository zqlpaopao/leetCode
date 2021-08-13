package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}




func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if (root.Val - p.Val) * (root.Val - q.Val) < 0{
		return root
	}
	if p.Val < root.Val{
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val{
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}


func main(){
	root1 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	l := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	r := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(lowestCommonAncestor(root1,l,r))
}