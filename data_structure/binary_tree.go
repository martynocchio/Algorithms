package main

import "errors"

type TreeNode struct {
	val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.val == value {
		return errors.New("this node val already exists")
	}

	if t.val > value {
		if t.LeftNode == nil {
			t.LeftNode = &TreeNode{val: value}
			return nil
		}
		return t.LeftNode.Insert(value)
	}

	if t.val < value {
		if t.RightNode == nil {
			t.RightNode = &TreeNode{val: value}
			return nil
		}
		return t.RightNode.Insert(value)
	}
	return nil
}

func (t *TreeNode) Remove(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if value < t.val {
		t.LeftNode = t.LeftNode.Remove(value)
		return t
	}
	if value > t.val {
		t.RightNode = t.RightNode.Remove(value)
		return t
	}

	if t.LeftNode == nil && t.RightNode == nil {
		t = nil
		return nil
	}

	if t.LeftNode == nil {
		t = t.RightNode
		return t
	}

	if t.RightNode == nil {
		t = t.LeftNode
		return t
	}

	smallestValOnRight := t.RightNode

	for {
		if smallestValOnRight != nil && smallestValOnRight.LeftNode != nil {
			smallestValOnRight = smallestValOnRight.LeftNode
		} else {
			break
		}
	}
	t.val = smallestValOnRight.val
	t.RightNode = t.RightNode.Remove(t.val)
	return t
}
