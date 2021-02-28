package btbuilder

import (
	"math"
	"strings"
)

func BuildTree(valuer interface{}) ([][]rune, bool) {
	var s string
	switch valuer.(type) {
	case string:
		s, _ = valuer.(string)
	case []string:
		temp, _ := valuer.([]string)
		s = sliceToString(temp)
	default:
		panic("type is not support")
	}
	if s == "" || len(s) == 0 {
		return nil, false
	}
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		if "#" == split[i] {
			split[i] = ""
		}
	}
	return modifyTree1(split), true
}

func BuildTreeLevelOrder(valuer interface{}) ([][]rune, bool) {
	var s string
	switch valuer.(type) {
	case string:
		s, _ = valuer.(string)
	case []string:
		temp, _ := valuer.([]string)
		s = sliceToString(temp)
	default:
		panic("type is not support")
	}
	if s == "" || len(s) == 0 {
		return nil, false
	}
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		if "#" == split[i] {
			split[i] = ""
		}
	}
	return modifyTreeLevelOrder1(split), true
}

func modifyTree1(objs []string) [][]rune {
	root, ok := buildTree2(objs)
	if !ok {
		panic("buildTree panic")
	}
	return modifyTree2(root)
}

func modifyTreeLevelOrder1(objs []string) [][]rune {
	root, ok := buildTreeLevelOrder(objs)
	if !ok {
		panic("buildTree panic")
	}
	return modifyTree2(root)
}

func modifyTree2(root *TreeNode) [][]rune {
	calcDisToParent(root)
	mymap, ok := getMap(root)
	if !ok {
		panic("getMap panic")
	}
	return mymap
}

func buildTree2(objs []string) (*TreeNode, bool) {
	if len(objs) == 0 || objs[0] == "" {
		return nil, false
	}
	var q []*TreeNode // change 1
	i := 0
	root := newTreeNode(objs[i], false)
	i++
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if i >= len(objs) {
			break
		}
		if objs[i] != "" {
			node.left = newTreeNode(objs[i], true)
			q = append(q, node.left)
		}
		i++
		if i >= len(objs) {
			break
		}
		if objs[i] != "" {
			node.right = newTreeNode(objs[i], false)
			q = append(q, node.right)
		}
		i++
	}
	return root, true
}

func buildTreeLevelOrder(objs []string) (*TreeNode, bool) {
	if len(objs) == 0 || objs[0] == "" {
		return nil, false
	}
	var q []*TreeNode
	i := 0
	root := newTreeNode(objs[i], false)
	i++
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if i >= len(objs) {
			break
		}
		if node == nil {
			i += 2
			q = append(q, &TreeNode{})
			q = append(q, &TreeNode{})
		} else {
			if objs[i] == "" {
				node.left = nil
			} else {
				node.left = newTreeNode(objs[i], true)
			}
			q = append(q, node.left)
			i++
			if i >= len(objs) {
				break
			}
			if objs[i] == "" {
				node.right = nil
			} else {
				node.right = newTreeNode(objs[i], false)
			}
			q = append(q, node.right)
			i++
		}
	}
	return root, true
}

func calcDisToParent(node *TreeNode) {
	if node == nil {
		return
	}
	calcDisToParent(node.left)
	calcDisToParent(node.right)
	max := 0
	var right, left int
	if node.left == nil {
		left = 0
	} else {
		left = len(node.left.rightList)
	}
	if node.right == nil {
		right = 0
	} else {
		right = len(node.right.leftList)
	}
	min := math.Min(float64(left), float64(right))
	for i := 0; float64(i) < min; i++ {
		max = int(math.Max(float64(max), float64(node.left.rightList[i]+node.right.leftList[i])))
	}
	dis := math.Max(float64((max+1)/2), 1)
	if node.left != nil {
		node.left.disToParent = int(dis)
	}
	if node.right != nil {
		node.right.disToParent = int(dis)
	}
	calcLeftList(node)
	calcRightList(node)
}

func calcRightList(node *TreeNode) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		return
	}
	if node.left != nil && node.right == nil {
		disToParent := node.left.disToParent
		for i := 1; i <= disToParent; i++ {
			node.rightList = append(node.rightList, -i)
		}
		for _, d := range node.left.rightList {
			node.rightList = append(node.rightList, d-disToParent-1)
		}
	} else if node.left == nil {
		disToParent := node.right.disToParent
		for i := 1; i <= disToParent; i++ {
			node.rightList = append(node.rightList, i)
		}
		for _, d := range node.right.rightList {
			node.rightList = append(node.rightList, d+disToParent+1)
		}
	} else {
		disToParent := node.right.disToParent
		for i := 1; i <= disToParent; i++ {
			node.rightList = append(node.rightList, i)
		}
		for _, d := range node.right.rightList {
			node.rightList = append(node.rightList, d+disToParent+1)
		}
		if len(node.left.rightList) > len(node.right.rightList) {
			for i := len(node.right.rightList); i < len(node.left.rightList); i++ {
				node.rightList = append(node.rightList, node.left.rightList[i]-disToParent-1)
			}
		}
	}

}

func calcLeftList(node *TreeNode) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		return
	}
	if node.left != nil && node.right == nil {
		disToParent := node.left.disToParent
		for i := 1; i <= disToParent; i++ {
			node.leftList = append(node.leftList, i)
		}
		for _, d := range node.left.leftList {
			node.leftList = append(node.leftList, d+disToParent+1)
		}
	} else if node.left == nil {
		disToParent := node.right.disToParent
		for i := 1; i <= disToParent; i++ {
			node.leftList = append(node.leftList, -i)
		}
		for _, d := range node.right.leftList {
			node.leftList = append(node.leftList, d-disToParent-1)
		}
	} else {
		disToParent := node.left.disToParent
		for i := 1; i <= disToParent; i++ {
			node.leftList = append(node.leftList, i)
		}
		for _, d := range node.left.leftList {
			node.leftList = append(node.leftList, d+disToParent+1)
		}
		if len(node.right.leftList) > len(node.left.leftList) {
			for i := len(node.left.leftList); i < len(node.right.leftList); i++ {
				node.leftList = append(node.leftList, node.right.leftList[i]-disToParent-1)
			}
		}
	}
}

func getMap(node *TreeNode) ([][]rune, bool) {
	if node == nil {
		return make([][]rune, 0), false
	}
	leftWidth := 0
	rightWidth := 0
	for _, w := range node.leftList {
		leftWidth = int(math.Max(float64(leftWidth), float64(w)))
	}
	for _, w := range node.rightList {
		rightWidth = int(math.Max(float64(rightWidth), float64(w)))
	}
	width := leftWidth + rightWidth + 1
	height := len(node.leftList)
	var map2 [][]rune // change 2

	for j := 0; j < height; j++ {
		var map1 []rune
		for i := 0; i < width; i++ {
			map1 = append(map1, ' ')
		}
		map2 = append(map2, map1)
	}
	fillMap(map2, node, int(leftWidth), 0)
	return map2, true
}

func fillMap(mymap [][]rune, node *TreeNode, x int, y int) {
	if node == nil {
		return
	}
	s := node.val
	for i := 0; i < len(s); i++ {
		mymap[y][x-node.leftList[0]+i] = []rune(s)[i]
	}
	if node.left != nil {
		disToParent := node.left.disToParent
		for i := 1; i <= disToParent; i++ {
			mymap[y+i][x-i] = '/'
		}
		fillMap(mymap, node.left, x-disToParent-1, y+disToParent+1)
	}
	if node.right != nil {
		disToParent := node.right.disToParent
		for i := 1; i <= disToParent; i++ {
			mymap[y+i][x+i] = '\\'
		}
		fillMap(mymap, node.right, x+disToParent+1, y+disToParent+1)
	}
}

func sliceToString(values []string) string {
	if len(values) == 0 {
		return ""
	}
	var result = strings.Join(values, ",")
	return result
}
