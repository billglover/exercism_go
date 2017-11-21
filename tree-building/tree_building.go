// Package tree implements the tree building logic for highly abstracted records.
// The records only contain an ID number and a parent ID number. The ID number is
// always between 0 (inclusive) and the length of the record list (exclusive).
// All records have a parent ID lower than their own ID, except for the root
// record, which has a parent ID that's equal to its own ID.
package tree

import (
	"fmt"
	"sort"
)

// Record represents an individual record. Each Record maintains a
// reference to its parent Record's ID.
type Record struct {
	ID, Parent int
}

// Node represents an individual node in a tree. Each node contains
// an identifier and a slice containing pointers to child nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Build takes a slice of records and builds a tree of nodes. It returns
// an error if the records contain loops, non-continuous nodes, or multiple
// root nodes.
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	// Sorting the records allows us to make some efficiency savings when
	// adding nodes to the tree.
	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent < records[j].Parent {
			return true
		}
		if records[i].Parent > records[j].Parent {
			return false
		}
		return records[i].ID < records[j].ID
	})

	// If the first node is not the root node the we should return an error
	// as there is no need to construct the tree.
	if records[0].ID != 0 {
		return nil, fmt.Errorf("no root node present")
	}

	tn := Node{}
	tn, err := tn.addChildren(records[1:], len(records))
	return &tn, err
}

// AddChildren is a recursive function that takes a sorted slice of records
// and adds them to a node in the tree.
func (n Node) addChildren(records []Record, max int) (Node, error) {
	for i, r := range records {

		// Nodes with an ID that is less than their parent or an ID
		// that is greater than the total number of nodes in the tree
		// are considered invalid. Nodes with an ID equal to their
		// parent are considered cyclical and also invalid.
		// Note: the root node has been handled above.
		if r.ID <= r.Parent || r.ID >= max {
			return n, fmt.Errorf("invalid node")
		}

		// Because we have a sorted list of records, we can ignore any
		// records with an ID greater than the current node ID.
		if r.Parent > n.ID {
			break
		}

		if r.Parent == n.ID {

			c, err := Node{ID: r.ID}.addChildren(records[i:], max)
			if err != nil {
				return n, err
			}

			n.Children = append(n.Children, &c)
		}
	}

	// we sort child nodes by ID in increasing numerical order
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].ID < n.Children[j].ID
	})
	return n, nil
}
