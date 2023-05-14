package main

import "fmt"

const null = -1

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func CreateBinaryTree(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	root := &TreeNode{input[0], nil, nil}
	treeQue := []*TreeNode{root}

	pingpong := true
	for i := 1; i < len(input); i++ {
		if input[i] == null {
			if !pingpong {
				treeQue = treeQue[1:]
			}
			pingpong = !pingpong
		} else {
			node := &TreeNode{input[i], nil, nil}
			if pingpong {
				treeQue[0].left = node
			} else {
				treeQue[0].right = node
				treeQue = treeQue[1:]
			}
			pingpong = !pingpong
			treeQue = append(treeQue, node)
		}
	}

	return root
}

func ShowSequence(root *TreeNode, operator func(*TreeNode)) {
	operator(root)
	fmt.Println()
}

func levelOrder(root *TreeNode) {
	if root == nil {
		fmt.Printf("Nil Tree")
		return
	}
	treeQue := []*TreeNode{root}
	for len(treeQue) > 0 {
		if treeQue[0].left != nil {
			treeQue = append(treeQue, treeQue[0].left)
		}
		if treeQue[0].right != nil {
			treeQue = append(treeQue, treeQue[0].right)
		}
		fmt.Printf("%d ", treeQue[0].val)
		treeQue = treeQue[1:]
	}
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Printf("%d ", root.val)
	inorder(root.right)
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Printf("%d ", root.val)
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	preorder(root.left)
	preorder(root.right)
}

func main() {
	input := []int{5, 3, 7, 2, 4, 6, 8, 1, null, null, null, null, null, null, 9}
	root := CreateBinaryTree(input)
	ShowSequence(root, levelOrder)
	ShowSequence(root, inorder)
	ShowSequence(root, postorder)
	ShowSequence(root, preorder)
}
