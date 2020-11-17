package gee

import "strings"

type node struct {
	wholePath string
	part      string
	children  []*node
	isWild    bool
}

/*
 * 根据part获取一个子节点
 */
func (n *node) getChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

/*
 * 根据part获取全部子节点
 */
func (n *node) getChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/*
 * 插入新的路径
 */
func (n *node) insert(path string, parts []string, level int) {
	if len(parts) == level {
		n.wholePath = path
		return
	}
	part := parts[level]
	child := n.getChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(path, parts, level+1)
}

/*
 * 查找路径，根据level匹配
 */
func (n *node) search(parts []string, level int) *node {
	if len(parts) == level || strings.HasPrefix(n.part, "*") {
		if n.wholePath == "" {
			return nil
		}
		return n
	}

	part := parts[level]
	//获取路径匹配的子节点
	children := n.getChildren(part)

	for _, child := range children {
		//bfs
		res := child.search(parts, level+1)
		if res != nil {
			return res
		}
	}
	return nil
}
