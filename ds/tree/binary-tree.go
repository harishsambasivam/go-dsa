package ds

import (
	"errors"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func createNode(val int) *Node {
	node := Node{val, nil, nil}
	return &node
}

type BinaryTree struct {
	root *Node
}

func (binaryTree *BinaryTree) Insert(value int, root *Node) (*Node, error) {

	node, err := binaryTree.BreadthFirstSearch(value)

	if err == nil && node.val != 0 {
		return node, errors.New("value already exists in the binary tree")
	}

	newNode := createNode(value)
	if value > root.val {
		if root.right != nil {
			binaryTree.Insert(value, root.right)
		} else {
			root.right = newNode
		}
	} else {
		if root.left != nil {
			binaryTree.Insert(value, root.left)
		} else {
			root.left = newNode
		}
	}

	return newNode, nil
}

func (binaryTree *BinaryTree) RemoveNode(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val > root.val {
		temp := binaryTree.RemoveNode(root.right, val)
		root.right = temp
	} else if val < root.val {
		root.left = binaryTree.RemoveNode(root.left, val)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			node, _ := binaryTree.InorderSuccessor(root.right)
			node.right = root.right
			node.left = root.left
			return node
		}
	}
	return root
}

func (binaryTree *BinaryTree) LevelOrderTraversal(root *Node) []int {
	stack := []*Node{root}
	result := []int{}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		result = append(result, node.val)
		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
	}
	return result
}

func (binaryTree *BinaryTree) PreOrderTraversal(root *Node, results *[]int) *[]int {
	if root == nil {
		return results
	}

	*results = append(*results, root.val)

	if root.left != nil {
		binaryTree.PreOrderTraversal(root.left, results)
	}

	if root.right != nil {
		binaryTree.PreOrderTraversal(root.right, results)
	}

	return results
}

func (binaryTree *BinaryTree) InOrderTraversal(root *Node, results *[]int) *[]int {
	if root == nil {
		return results
	}

	if root.left != nil {
		binaryTree.InOrderTraversal(root.left, results)
	}

	*results = append(*results, root.val)

	if root.right != nil {
		binaryTree.InOrderTraversal(root.right, results)
	}

	return results
}

func (binaryTree *BinaryTree) PostOrderTraversal(root *Node, results *[]int) *[]int {
	if root == nil {
		return results
	}

	if root.left != nil {
		binaryTree.PostOrderTraversal(root.left, results)
	}

	if root.right != nil {
		binaryTree.PostOrderTraversal(root.right, results)
	}

	*results = append(*results, root.val)

	return results
}

func (binaryTree *BinaryTree) BreadthFirstSearch(target int) (*Node, error) {
	stack := []*Node{binaryTree.root}
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr.val == target {
			return curr, nil
		}
		if curr.left != nil {
			stack = append(stack, curr.left)
		}
		if curr.right != nil {
			stack = append(stack, curr.right)
		}
	}
	return &Node{0, nil, nil}, errors.New("value not found")
}

func (binaryTree *BinaryTree) DepthFirstSearch(root *Node, target int) (*Node, error) {
	if root != nil {
		return &Node{0, nil, nil}, errors.New("args[0] is not a valid node")
	}

	if root.val == target {
		return root, nil
	}

	if root.left != nil {
		binaryTree.DepthFirstSearch(root.left, target)
	}

	if root.right != nil {
		binaryTree.DepthFirstSearch(root.right, target)
	}

	return &Node{0, nil, nil}, errors.New("value not found")
}

func (binaryTree *BinaryTree) InorderSuccessor(node *Node) (*Node, error) {
	if node == nil {
		return node, errors.New("node is not a valid one")
	}
	curr := node
	prev := curr
	for curr.left != nil {
		prev = curr
		curr = curr.left
	}
	// removing the link to avoid cycle
	prev.left = nil
	return curr, nil
}

func (binaryTree *BinaryTree) InorderPredeccessor(node *Node) (*Node, error) {
	if node == nil {
		return node, errors.New("node is not a valid one")
	}
	curr := node
	prev := curr
	for curr.right != nil {
		curr = curr.right
	}
	// removing the link to avoid cycle
	prev.right = nil
	return node, nil
}
