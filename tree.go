package tree

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 若不为空，所有节点的左子小于父节点，所有节点的右子大于父节点
// 子树也为二叉排序树
// 没有相等的节点

//插入节点，若为空，则插入根节点，小于此节点，则递归插入左子树 反之右
func Insert(n *Node, v int) *Node {
	if n == nil {
		fmt.Println("插入结点", v)
		return &Node{Val: v}
	}

	if v < n.Val {
		if n.Left == nil {
			a := &Node{Val: v}
			n.Left = a
			fmt.Printf("插入%d的左节点%d", n.Val, v)

			return a
		} else {
			Insert(n.Left, v)
		}

	} else if v > n.Val {

		if n.Right == nil {
			a := &Node{Val: v}
			n.Right = a
			fmt.Printf("插入%d的右节点%d", n.Val, v)
			return a
		} else {
			Insert(n.Right, v)
		}

	}

	return n

}

//删除节点，弱左子节点为空，右子树承接此节点。弱右子节点为空，左子树承接此节点。
// 弱左右都不为空，则选择左子树最大节点，或右子树最小节点来代替当前节点，并删除对应节点
func Delete(n *Node, v int) *Node {
	if n == nil {
		return nil
	}

	// 先判断输入节点是否为待删除节点，若不是找到待删除节点
	if v < n.Val {
		n.Left = Delete(n.Left, v)
		return n
	} else if v > n.Val {
		n.Right = Delete(n.Right, v)
		return n
	} else {
		if n.Left == nil {
			n = n.Right
		} else if n.Right == nil {
			n = n.Left
		} else {
			t := Min(n.Right)
			n.Val = t.Val
			n.Right = DeleteMin(n.Right)
			return n
		}
	}
	return nil
}

func DeleteMin(n *Node) *Node {

	if n == nil {
		return nil
	}

	if n.Left == nil {

		n = n.Right
	} else {
		n.Left = DeleteMin(n.Left)
	}

	return n
}

func Min(n *Node) *Node {
	a := n

	for a.Left != nil {
		a = a.Left
	}
	return a
}

func PreorderShow(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.Val)

	PreorderShow(n.Left)

	PreorderShow(n.Right)
	return
}

func SequentialShow(n *Node) {
	if n == nil {
		return
	}

	PreorderShow(n.Left)
	fmt.Println(n.Val)
	PreorderShow(n.Right)
	return
}

func PostorderShow(n *Node) {
	if n == nil {
		return
	}
	PreorderShow(n.Left)
	PreorderShow(n.Right)
	fmt.Println(n.Val)

}
