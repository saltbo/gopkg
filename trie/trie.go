//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package trie

import (
	"fmt"
	"log"
	"strings"

	"github.com/saltbo/gopkg/safemap"
)

const (
	defaultDelimiter = "/"
	defaultWildcard  = ".+"
)

// Trie is a tree
type Tree struct {
	delimiter string
	root      *Node
}

// New returns a new Trie object.
func New() *Tree {
	return &Tree{
		delimiter: defaultDelimiter,
		root:      NewNode(),
	}
}

// SetDelimiter sets the delimiter of the trie object.
func (trie *Tree) SetDelimiter(delimiter string) {
	trie.delimiter = delimiter
}

func (trie *Tree) splitPattern(pattern string) []string {
	p := strings.TrimRight(pattern, trie.delimiter)
	parts := strings.Split(p, trie.delimiter)
	if parts[0] == "" {
		parts[0] = string(pattern[0])
	}
	return parts
}

func (trie *Tree) Dump() {
	log.Printf("--------- Tree Dump ------------\n")
	log.Printf("-%s: %v \n", trie.root.Key, trie.root.entry)
	lastNode := trie.root
	lastNode.dump("")
	log.Printf("--------- Tree Dump End ----------\n\n")
}

func (trie *Tree) Empty() bool {
	if !trie.root.Children.Empty() {
		return false
	} else if trie.root.entry != nil {
		return false
	}

	return true
}

func (trie *Tree) Exist(pattern string) (interface{}, bool) {
	node := trie.root
	parts := trie.splitPattern(pattern)[1:]
	for _, part := range parts {
		if child, ok := node.directChild(part); ok {
			node = child
			continue
		}

		return nil, false
	}

	if node.entry == nil {
		return nil, false
	}

	return node.entry, true
}

func (trie *Tree) Put(pattern string, entry interface{}) {
	node := trie.root
	parts := trie.splitPattern(pattern)[1:]
	for _, part := range parts {
		node = node.insertChild(part)
	}

	node.entry = entry
}

func (trie *Tree) Match(pattern string) (interface{}, bool) {
	node := trie.root
	parts := trie.splitPattern(pattern)[1:]
	for _, part := range parts {
		if child, ok := node.directChild(part); ok {
			node = child
		} else if wildcardNode, ok := node.wildcardChild(part); ok {
			return wildcardNode.entry, true
		} else {
			goto traverseParent
		}
	}

	if node.entry != nil {
		return node.entry, true
	}

traverseParent:
	if wildcardNode, ok := node.wildcardParent(); ok {
		return wildcardNode.entry, true
	}

	return nil, false
}

func (trie *Tree) Del(pattern string) error {
	node := trie.root
	parts := trie.splitPattern(pattern)[1:]
	for _, part := range parts {
		if child, ok := node.directChild(part); ok {
			node = child
		} else {
			return fmt.Errorf("pattern %v is in not existed", pattern)
		}
	}

	// 删除当前节点entry
	node.entry = nil
	// 判断当前节点是否有子节点，如果有则直接返回，如果没有子节点则执行向上遍历清理逻辑
	if !node.Children.Empty() {
		return nil
	}

	// 向上遍历查找上游节点是否为空节点，如果是则一并清理
	for node.Parent != nil {
		if node.Children.Empty() && node.entry == nil {
			node.delete()
		}

		node = node.Parent
	}

	return nil
}

// Node is the tree node of the Trie.
type Node struct {
	Key      string
	Parent   *Node
	Children *safemap.Map //nodeKey => Node
	entry    interface{}
}

func NewNode() *Node {
	node := &Node{
		Key:      "/",
		Children: safemap.New(),
	}

	return node
}

// return a child
func (n *Node) insertChild(part string) *Node {
	if v, ok := n.Children.Load(part); ok {
		return v.(*Node)
	}

	newNode := &Node{Key: part, Parent: n, Children: safemap.New()} //创建新节点
	n.Children.Store(part, newNode)
	return newNode
}

func (n *Node) directChild(part string) (*Node, bool) {
	if v, ok := n.Children.Load(part); ok {
		return v.(*Node), true
	}

	return nil, false
}

func (n *Node) wildcardChild(part string) (*Node, bool) {
	if wildcardNode, ok := n.Children.Load(part + defaultWildcard); ok {
		return wildcardNode.(*Node), true
	} else if wildcardNode, ok := n.Children.Load(defaultWildcard); ok {
		return wildcardNode.(*Node), true
	}

	return nil, false
}

// 向上遍历父节点是否存在通配路由
func (n *Node) wildcardParent() (*Node, bool) {
	parentNode := n.Parent
	for parentNode != nil {
		if wildcardNode, ok := parentNode.Children.Load(defaultWildcard); ok {
			return wildcardNode.(*Node), true
		}

		parentNode = parentNode.Parent
	}

	return nil, false
}

// delete a child
func (n *Node) delete() {
	n.Parent.Children.Delete(n.Key)
}

func (n *Node) dump(format string) {
	if format == "" {
		format = "|--%s: %v \n"
	}

	n.Children.Range(func(k string, v interface{}) bool {
		node := v.(*Node)
		log.Printf(format, node.Key, node.entry)
		if node.Children.Empty() {
			return true
		}

		node.dump("    " + format)
		return true
	})
}
