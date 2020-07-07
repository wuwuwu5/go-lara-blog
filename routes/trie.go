package routes

import "strings"

type node struct {
	pattern  string // 只有最后一级才有有值 默认为""
	children []*node
	part     string // url中的一部分
	isWild   bool   // 是否精确匹配  part 含有:或*为true
}

// 查找第一个符合条件的数据
func (this *node) matchChild(part string) *node {
	for _, child := range this.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// 查找所有符合条件的数据
func (this *node) matchChildRen(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range this.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// 插入节点
func (this *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		this.pattern = pattern
		return
	}

	part := parts[height]

	child := this.matchChild(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		this.children = append(this.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// 查找
func (this *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(this.part, "*") {
		if this.pattern == "" {
			return nil
		}
		return this
	}

	part := parts[height]

	nodes := this.matchChildRen(part)

	for _, child := range nodes {
		result := child.search(parts, height+1)

		if result != nil {
			return result
		}
	}

	return nil
}
