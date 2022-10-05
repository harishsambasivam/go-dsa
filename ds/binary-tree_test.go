package ds

import (
	"testing"
)

func TestInsertToBinaryTree(t *testing.T) {
	t.Run("inserting a node to binary tree", func(t *testing.T) {
		BST := BinaryTree{
			root: &Node{
				1,
				nil,
				nil,
			},
		}

		BST.Insert(2, BST.root)

		if BST.root.val != 1 {
			t.Errorf("root should have value = 1")
		}

		if BST.root.right.val != 2 {
			t.Errorf("root should have value = 2")
		}
	})
}
