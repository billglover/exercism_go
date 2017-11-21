package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

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
			return n, nil
		}

		if r.Parent == n.ID {

			c, err := Node{ID: r.ID}.addChildren(records[i:], max)
			if err != nil {
				return n, err
			}

			n.Children = append(n.Children, &c)
		}
	}

	// we expect node children to be sorted in increasing numerical order
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].ID < n.Children[j].ID
	})
	return n, nil
}

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
