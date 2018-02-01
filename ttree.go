package ttreego

import (
	"github.com/juju/errors"
)

// TTreeNodeMax is the max size of child count, current is 0xff
const TTreeNodeMax = 0xff + 1

var (
	// Can't find string while RemoveString
	ErrTTreeNodeNotFound = errors.New("TTreeNode not found")
)

// TTreeNode represents a trie tree node
type TTreeNode struct {
	term     bool
	childCnt int
	level    int
	value    byte
	parent   *TTreeNode
	nodes    []*TTreeNode
}

// addChild add a child node to the node, return the child node
func (n *TTreeNode) addChild(idx byte) *TTreeNode {
	if nil == n.nodes {
		n.nodes = make([]*TTreeNode, TTreeNodeMax)
	}
	node := n.nodes[idx]
	if nil == node {
		node = NewTTreeNode()
		n.nodes[idx] = node
		// If allocate new, update child cnt
		n.childCnt++
		node.parent = n
		node.level = n.level + 1
		node.value = idx
	}

	return node
}

// removeChild remove a child node from the node, return the child node
func (n *TTreeNode) removeChild(idx byte) *TTreeNode {
	if nil == n.nodes {
		return nil
	}
	node := n.nodes[idx]
	if nil == node {
		return nil
	}
	n.nodes[idx] = nil
	n.childCnt--
	if 0 == n.childCnt {
		n.nodes = nil
	}
	return node
}

// clear clears all child node
func (n *TTreeNode) clear() {
	n.nodes = nil
	n.childCnt = 0
}

// getChild get the child node by charactor
func (n *TTreeNode) getChild(idx byte) *TTreeNode {
	if nil == n.nodes {
		return nil
	}
	return n.nodes[idx]
}

// NewTTreeNode creates a new trie tree node
func NewTTreeNode() *TTreeNode {
	return &TTreeNode{}
}

// TTree is a trie tree node with deep zero and with nil parent node
type TTree TTreeNode

// NewTTree creates a new trie tree
func NewTTree() *TTree {
	return &TTree{}
}

// Dump not implement
func (t *TTree) Dump(level int) string {
	return ""
}

// AddString add a string to the tree
func (t *TTree) AddString(str string) bool {
	return t.Add([]byte(str))
}

// Add add a byte sequence of string to the tree
func (t *TTree) Add(strData []byte) bool {
	node := (*TTreeNode)(t)
	for i := range strData {
		value := strData[i]
		cnode := node.addChild(value)
		// Set terminate flag
		if i == len(strData)-1 {
			cnode.term = true
			break
		}
		// Continue adding next charactor
		node = cnode
	}

	return true
}

// RemoveString remove a string from the tree
func (t *TTree) RemoveString(str string) error {
	return t.Remove([]byte(str))
}

// Remove remove a byte sequence of the string from the tree
func (t *TTree) Remove(strData []byte) error {
	if nil == strData ||
		0 == len(strData) {
		return nil
	}
	// First reach the last element of the strData in ptree
	termNode := t.findTermNode(strData, true)
	if nil == termNode {
		// Not in the tree
		return ErrTTreeNodeNotFound
	}
	// Check term node has child
	node := termNode
	for i := len(strData) - 1; i >= 0; i-- {
		if nil == node {
			return errors.Errorf("Can't find tree node with charactor %c", strData[i])
		}

		pnode := node.parent
		if nil == pnode {
			return errors.Errorf("Can't find parent node with charactor %c", strData[i])
		}

		if nil == node.nodes {
			// If current node has no child, remove the node
			if pnode.removeChild(strData[i]) != node {
				return errors.Errorf("Remove current node from parent node failed, charactor = %v", strData[i])
			}
		} else {
			// Current node has child node, just remove term flag
			node.term = false
		}

		node = pnode
	}

	return nil
}

// MatchString check if a string matches the tree
func (t *TTree) MatchString(str string) bool {
	return t.Match([]byte(str))
}

// Match check if a byte sequence of a string matches the tree
func (t *TTree) Match(strData []byte) bool {
	termNode := t.findTermNode(strData, false)
	if nil == termNode {
		return false
	}
	return true
}

// toNode convert a tree to tree node
func (t *TTree) toNode() *TTreeNode {
	return (*TTreeNode)(t)
}

// findTermNode find the terminate node of the string.If full is true, only return when
// last charactor of the string is terminate node. If full is false, any node of the
// string with terminate flag will be returned
func (t *TTree) findTermNode(strData []byte, full bool) *TTreeNode {
	node := t.toNode()

	for i := range strData {
		value := strData[i]

		cnode := node.getChild(value)
		if nil == cnode {
			return nil
		}

		if !full {
			if cnode.term {
				return cnode
			}
		}

		if i == len(strData)-1 {
			// Last charactor
			if !cnode.term {
				return nil
			}
			return cnode
		}

		node = cnode
	}

	return nil
}

// Reset reset the tree, remove all child node
func (t *TTree) Reset() {
	node := t.toNode()
	node.clear()
}
