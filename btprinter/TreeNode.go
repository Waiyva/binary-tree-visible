package btprinter

type TreeNode struct {
	val                 string
	left                *TreeNode
	right               *TreeNode
	disToParent         int
	leftList, rightList []int
}

type Object interface{}

func newTreeNode(obj Object, left bool) *TreeNode {
	var node TreeNode
	val, ok := obj.(string)
	if !ok {
		panic("type assert panic")
		return nil
	}
	disToParent := 0
	node.val = val
	node.disToParent = disToParent
	len := len(val)
	var joint int
	if len%2 == 1 {
		joint = (len + 1) / 2
	} else if left {
		joint = len / 2
	} else {
		joint = len/2 + 1
	}
	node.leftList = append(node.leftList, joint-1)
	node.rightList = append(node.rightList, len-joint)
	return &node
}
