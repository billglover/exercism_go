package tree

import (
	"errors"
	"fmt"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	// return an empty tree rather than an error when we have been given
	// no records
	if len(records) == 0 {
		return nil, nil
	}

	// this is the root of our tree
	root := &Node{}

	// a slice of pointers to Nodes containing the root node
	todo := []*Node{root}

	n := 1
	for {
		if len(todo) == 0 {
			break
		}

		// a slice of pointers to Nodes containing a nil Node
		newTodo := []*Node(nil)

		// loop over slice containing our root node
		for _, c := range todo {

			// loop over the records we've been given
			for _, r := range records {

				// if current has this node as its parent
				if r.Parent == c.ID {

					// validate that this node ID is sensible
					if r.ID < c.ID {
						return nil, errors.New("a")
					} else if r.ID == c.ID {
						if r.ID != 0 {
							return nil, fmt.Errorf("b")
						}

						// add it to the current tree
					} else {
						n++
						switch len(c.Children) {

						case 0:
							// if record has no children add it as a child
							// TODO: unclear if this is zero children or not
							nn := &Node{ID: r.ID}
							c.Children = []*Node{nn}
							newTodo = append(newTodo, nn)
						case 1:
							// if the record has a single child add it to the node
							// TODO: unclear why order matters here
							nn := &Node{ID: r.ID}
							if c.Children[0].ID < r.ID {
								c.Children = []*Node{c.Children[0], nn}
								newTodo = append(newTodo, nn)
							} else {
								c.Children = []*Node{nn, c.Children[0]}
								newTodo = append(newTodo, nn)
							}
						default:
							// TODO: what is len(c.Children) in the default case
							nn := &Node{ID: r.ID}
							newTodo = append(newTodo, nn)
						breakpoint:
							for range []bool{false} {
								for i, cc := range c.Children {
									if cc.ID > r.ID {
										a := make([]*Node, len(c.Children)+1)
										copy(a, c.Children[:i])
										copy(a[i+1:], c.Children[i:])
										copy(a[i:i+1], []*Node{nn})
										c.Children = a
										break breakpoint
									}
								}
								c.Children = append(c.Children, nn)
							}
						}
					}
				}
			}
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, fmt.Errorf("c")
	}
	if err := chk(root, len(records)); err != nil {
		return nil, err
	}
	return root, nil
}

// chk validates that no node in the tree can have ID greater than or equal to the
// number of nodes in the tree.
// TODO: simplify by combining the conditions
// TODO: remove and see if we can add this check at build time
func chk(n *Node, m int) (err error) {
	fmt.Println(n)
	fmt.Println(n.ID, m)
	if n.ID > m {
		return fmt.Errorf("node can't have an ID greater than the number of nodes")
	} else if n.ID == m {
		return fmt.Errorf("node can't have an ID equal to the number of nodes")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				fmt.Println(n.Children[i], m, err)
				return
			}
		}
		return
	}
}
