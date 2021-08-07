package main

import (
	"fmt"
)

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
func isValidBST(root *TreeNode) bool {
	return isValidBSTImpl(root, -1 << 63, 1 << 63 -1)
}


func isValidBSTImpl(n *TreeNode, min int, max int) bool {
	if n.Right == nil && n.Left == nil{
		return true
	}

	if n.Val < min || n.Val >= max{
		return false
	}

	return isValidBSTImpl(n.Left,min,n.Val) && isValidBSTImpl(n.Right,n.Val,max)
}

var Pre = -1 << 63

func isValidBST1(root *TreeNode)bool{
	if root == nil{
		return true
	}
	if !isValidBST1(root.Left){
		return false
	}
	if root.Val <= Pre{
		return false
	}
	Pre = root.Val

	return isValidBST1(root.Right)

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
	//fmt.Println(isValidBST(root1))
	fmt.Println(isValidBST1(root1))

}