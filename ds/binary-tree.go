package ds

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

func (binaryTree *BinaryTree) Insert(value int, root *Node) {
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
}

func (binaryTree *BinaryTree) RemoveNode(value int) {

}

func (binaryTree *BinaryTree) DepthFirstSearch(value int) {

}

func (binaryTree *BinaryTree) BreadthFirstSearch(value int) {

}
