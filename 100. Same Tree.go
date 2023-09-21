package main

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func IsSameTree2(p *TreeNode, q *TreeNode) bool {
	stack := []*TreeNode{p, q}
	for len(stack) > 0 {
		p, q = stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		stack = append(stack, p.Left, q.Left, p.Right, q.Right)
	}
	return true
}
