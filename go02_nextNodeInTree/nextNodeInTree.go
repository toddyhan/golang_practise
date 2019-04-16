package go02_nextNodeInTree

/*
题目描述：
给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。

public class TreeLinkNode {

    int val;
    TreeLinkNode left = null;
    TreeLinkNode right = null;
    TreeLinkNode next = null;

    TreeLinkNode(int val) {
        this.val = val;
    }
}

思路：
1.如果当前节点的右子节点不为空，则下一节点是右子树的最左节点
2.如果当前节点的右子节点为空，则下一节点是向上找第一个左链接指向的树包含该节点的祖先节点。

*/

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
	data int
}

func GetNext(nodeInput *TreeNode) (result interface{}) {
	if nodeInput.Right != nil {
		node := nodeInput.Right
		for {
			if node.Left != nil {
				node = node.Left
			} else {
				break
			}
		}
		return node
	} else {
		node := nodeInput
		parent := nodeInput.Parent
		for {
			if parent == nil {
				break
			}
			if parent.Left == node {
				return parent
			}
			parent = parent.Parent
		}
	}
	return nil
}