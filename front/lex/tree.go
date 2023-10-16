package lex

import (
	"errors"
)

/*
To match operators, we're using a prefix tree
*/

type treeNode struct {
	Char          byte
	Corresponding Type
	Next          []*treeNode
}

func newCharsTree() *treeNode {
	tree := new(treeNode)

	for key, value := range chars {
		if err := tree.insert(key, value); err != nil {
			panic(err)
		}
	}

	return tree
}

func (t *treeNode) insert(repr string, payload Type) error {
	for _, node := range t.Next {
		if node.Char == repr[0] {
			if len(repr) == 1 {
				if node.Corresponding != Unknown {
					return errors.New("character/tree.go:Node.insert(): overrides already existing value")
				}

				node.Corresponding = payload

				return nil
			}

			return node.insert(repr[1:], payload)
		}
	}

	node := &treeNode{
		Char: repr[0],
	}

	if len(repr) == 1 {
		node.Corresponding = payload
	} else {
		if err := node.insert(repr[1:], payload); err != nil {
			return err
		}
	}

	t.Next = append(t.Next, node)

	return nil
}

func (t *treeNode) Match(str string) Type {
	for _, node := range t.Next {
		if node.Char == str[0] {
			if len(str) == 1 {
				return node.Corresponding
			}

			return node.Match(str[1:])
		}
	}

	return Unknown
}

func (t *treeNode) HasPrefix(prefix string) bool {
	for _, node := range t.Next {
		if node.Char == prefix[0] {
			if len(prefix) == 1 {
				return true
			}

			return node.HasPrefix(prefix[1:])
		}
	}

	return false
}
