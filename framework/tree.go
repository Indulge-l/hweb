package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{root: newNode()}
}

func (t *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := t.root
	if n.matchNode(uri) != nil {
		return errors.New("router exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segment := range segments {

		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node

		childrenNodes := n.filterChildrenNodes(segment)

		if len(childrenNodes) > 0 {
			for _, cnode := range childrenNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment

			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.children = append(n.children, cnode)
			objNode = cnode
		}

		n = objNode
	}
	return nil
}

func (t *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := t.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}

type node struct {
	isLast   bool
	segment  string
	handler  ControllerHandler
	children []*node
}

func newNode() *node {
	return &node{}
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层不满足的节点
func (n *node) filterChildrenNodes(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}
	// 如果是通配符，则下一层节点全都满足
	if isWildSegment(segment) {
		return n.children
	}
	var nodes []*node
	for _, cnode := range n.children {
		if isWildSegment(cnode.segment) || cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

// TODO 有优化空间
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildrenNodes(segment)
	if len(cnodes) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}
