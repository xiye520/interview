package rbtree

import (
	"errors"
	"math/rand"
	"testing"
)

// 测试插入和删除的时候红黑树的性质有没有发生改变：
var tree *RBTree

// go test -v -test.run TestInsert
func TestInsert(t *testing.T) {
	tree = &RBTree{}
	for i := 1; i <= 1000; i++ {
		x := rand.Int()
		tree.insert(x)
	}
	preperties(t)
}

// go test -v -test.run TestDelete
func TestDelete(t *testing.T) {
	tree = &RBTree{}
	for i := 1; i <= 1000; i++ {
		tree.insert(i)
	}
	for i := 1; i <= 1000; i++ {
		tree.delete(rand.Intn(1000))
		tree.delete(i)
		preperties(t)
	}
}

// five properties of rbtree
//1 each node be black or red
//2 root node is black
//3 each leaf node(nil) is black
//4 if one node is red ,its children are both black
//5 for each node,the same number of black nodes are included in
//the simple path from the node to all its descendant leaf nodes
func preperties(t *testing.T) {
	//the numbers 2,4,5 need to test
	//2
	if !isBlack(tree.root) {
		t.Error("tree's root is not black")
	}
	//4
	err := colorOfChildren(tree.root)
	if err != nil {
		t.Error("red nodes have red children")
	}
	//5
	_, err = theNumOfBlack(tree.root)
	if err != nil {
		t.Error("tree's num of black nodes are different")
	}
}
func colorOfChildren(n *node) (err error) {
	if n == nil {
		return
	}
	err = colorOfChildren(n.leftNode)
	if err != nil {
		return errors.New("the forth property is destroyed")
	}
	err = colorOfChildren(n.rightNode)
	if err != nil {
		return errors.New("the forth property is destroyed")
	}
	if n.color == RBTRed {
		if isBlack(n.leftNode) && isBlack(n.rightNode) {
			return
		} else {
			return errors.New("the forth property is destroyed")
		}
	}
	return
}

func theNumOfBlack(n *node) (num int, err error) {
	if n == nil {
		return 0, nil
	}
	leftNum, err := theNumOfBlack(n.leftNode)
	if err != nil {
		return 0, errors.New("the fifth property is destroyed")
	}
	rightNum, err := theNumOfBlack(n.rightNode)
	if err != nil {
		return 0, errors.New("the fifth property is destroyed")
	}
	if leftNum != rightNum {
		return 0, errors.New("the fifth property is destroyed")
	}
	if n.color == RBTBlack {
		return leftNum + 1, nil
	} else {
		return leftNum, nil
	}
}
