package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res []int

//-- ----------------------------
//--> @Description 前序遍历
//--> @Param
//--> @return
//-- ----------------------------
func preorderTraversal(root *TreeNode){
	if root == nil{
		return
	}
	//先加入跟节点，在加入左孩子，在加入右孩子
	res = append(res,root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return
}

//-- ----------------------------
//--> @Description 中序遍历
//--> @Param
//--> @return
//-- ----------------------------
func inorderTraversal(root *TreeNode){
	if root == nil {
		return
	}
	//先加入左孩子，在加入跟节点，在加入右孩子
	inorderTraversal(root.Left)
	res = append(res,root.Val)
	inorderTraversal(root.Right)
}

//-- ----------------------------
//--> @Description 后序遍历
//--> @Param
//--> @return
//-- ----------------------------
func postTraversal(root *TreeNode){
	if root == nil{
		return
	}
	postTraversal(root.Left)
	postTraversal(root.Right)
	res = append(res,root.Val)

}
/*****************************迭代法**************************************/
//-- ----------------------------
//--> @Description 前序遍历
//--> @Param
//--> @return
//-- ----------------------------
func scanPreorderTraversal(root *TreeNode)(res []int){
	if root == nil {
		return nil
	}
	var stack = list.New()
	stack.PushBack(root.Right)
	stack.PushBack(root.Left)
	res=append(res,root.Val)
	for stack.Len()>0 {
		e:=stack.Back()
		stack.Remove(e)
		//e是Element类型，其值为e.Value.由于Value为接口，所以要断言
		node := e.Value.(*TreeNode)
		if node==nil{
			continue
		}
		res=append(res,node.Val)
		stack.PushBack(node.Right)
		stack.PushBack(node.Left)
	}
	return res

}

//-- ----------------------------
//--> @Description
//--> @Param
//--> @return
//-- ----------------------------
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var stack = list.New()
	stack.PushBack(root.Left)
	stack.PushBack(root.Right)
	res=append(res,root.Val)
	for stack.Len()>0 {
		e:=stack.Back()
		stack.Remove(e)
		node := e.Value.(*TreeNode)//e是Element类型，其值为e.Value.由于Value为接口，所以要断言
		if node==nil{
			continue
		}
		res=append(res,node.Val)
		stack.PushBack(node.Left)
		stack.PushBack(node.Right)
	}
	for i:=0;i<len(res)/2;i++{
		res[i],res[len(res)-i-1] = res[len(res)-i-1],res[i]
	}
	return res
}

//-- ----------------------------
//--> @Description
//--> @Param
//--> @return
//-- ----------------------------
func scanInorderTraversal(root *TreeNode) []int {
	rootRes:=[]int{}
	if root==nil{
		return nil
	}
	stack:=list.New()
	node:=root
	//先将所有左节点找到，加入栈中
	for node!=nil{
		stack.PushBack(node)
		node=node.Left
	}
	//其次对栈中的每个节点先弹出加入到结果集中，再找到该节点的右节点的所有左节点加入栈中
	for stack.Len()>0{
		e:=stack.Back()
		node:=e.Value.(*TreeNode)
		stack.Remove(e)
		//找到该节点的右节点，再搜索他的所有左节点加入栈中
		rootRes=append(rootRes,node.Val)
		node=node.Right
		for node!=nil{
			stack.PushBack(node)
			node=node.Left
		}
	}
	return rootRes
}



func main(){
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//preorderTraversal(root)
	//inorderTraversal(root)
	//postTraversal(root)
	//fmt.Println(scanPreorderTraversal(root))
	fmt.Println(scanPreorderTraversal(root))
	fmt.Println(postorderTraversal(root))
	fmt.Println(scanInorderTraversal(root))
	//fmt.Println(res)
}