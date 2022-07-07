package main

import "fmt"

type node struct {
	val string
	left *node
	right *node
}
func inorder(root *node){
	if root==nil{
		return
	}

	inorder(root.left)
	fmt.Println(root.val)
	inorder(root.right)
}
func preorder(root *node){
	if root==nil{
		return
	}
	fmt.Println(root.val)
	preorder(root.left)
	preorder(root.right)
}
func postorder(root *node){
	if root==nil{
		return
	}

	postorder(root.left)
	postorder(root.right)
	fmt.Println(root.val)
}

func main() {
	root:=node{"+",nil,nil}
	root.left=&node{"a",nil,nil}
	root.right=&node{"-",nil,nil}
	root.right.left=&node{"b",nil,nil}
	root.right.right=&node{"c",nil,nil}
	// 	preorder(&root)
	postorder(&root)
}
