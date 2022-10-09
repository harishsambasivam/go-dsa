package ds

import (
	"reflect"
	"testing"
)

/**
https://visualgo.net/en/bst
*/

var root = &Node{
	41,
	nil,
	nil,
}
var BST = BinaryTree{
	root,
}

func TestInsertToBinaryTree(t *testing.T) {
	t.Run("inserting a node to binary tree", func(t *testing.T) {

		BST.Insert(20, BST.root)
		BST.Insert(65, BST.root)
		BST.Insert(11, BST.root)
		BST.Insert(29, BST.root)
		BST.Insert(50, BST.root)
		BST.Insert(91, BST.root)
		BST.Insert(32, BST.root)
		BST.Insert(72, BST.root)
		BST.Insert(99, BST.root)

		if BST.root.val != 41 {
			t.Errorf("expected 5, received %d", BST.root.val)
		}

		_, err := BST.Insert(99, BST.root)
		if err.Error() != "value already exists in the binary tree" {
			t.Errorf("expected : error message, got: nil")
		}

		received := BST.LevelOrderTraversal(root)
		expected := []int{41, 20, 65, 11, 29, 50, 91, 32, 72, 99}
		if reflect.DeepEqual(received, expected) == false {
			t.Errorf("expected : %v, got: %v", expected, received)
		}
	})
}

func TestPreOrderTraversal(t *testing.T) {
	t.Run("tree traversal using pre-order traversal", func(t *testing.T) {
		results := []int{}
		BST.PreOrderTraversal(root, &results)
		expected := []int{41, 20, 11, 29, 32, 65, 50, 91, 72, 99}
		if reflect.DeepEqual(results, expected) == false {
			t.Errorf("expected : %v, got: %v", expected, results)
		}
	})

}

func TestInOrderTraversal(t *testing.T) {
	t.Run("tree traversal using in-order traversal", func(t *testing.T) {
		results := []int{}
		BST.InOrderTraversal(root, &results)
		expected := []int{11, 20, 29, 32, 41, 50, 65, 72, 91, 99}
		if reflect.DeepEqual(results, expected) == false {
			t.Errorf("expected : %v, got: %v", expected, results)
		}
	})

	// BST.Insert(5, BST.root)
}

func TestPostOrderTraversal(t *testing.T) {
	t.Run("tree traversal using post-order traversal", func(t *testing.T) {
		results := []int{}
		BST.PostOrderTraversal(root, &results)
		expected := []int{11, 32, 29, 20, 50, 72, 99, 91, 65, 41}
		if reflect.DeepEqual(results, expected) == false {
			t.Errorf("expected : %v, got: %v", expected, results)
		}
	})
}

func TestRemoveNode(t *testing.T) {
	t.Run("removing a node from binary tree", func(t *testing.T) {
		BST.RemoveNode(BST.root, 65)
		received := []int{}
		BST.PreOrderTraversal(BST.root, &received)
		expected := []int{41, 20, 11, 29, 32, 72, 50, 91, 99}
		if reflect.DeepEqual(received, expected) == false {
			t.Errorf("expected : %v, got: %v", expected, received)
		}
	})
}
