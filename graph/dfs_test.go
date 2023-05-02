package graph

import (
	"reflect"
	"testing"
)

func TestDFS_PreOrderTraverse(t *testing.T) {
	result := []int{}
	dfs := &DFS{
		root: &traversable_mock{
			value: 1,
			children: []Traversable{
				&traversable_mock{value: 2, children: []Traversable{
					&traversable_mock{value: 4, children: nil},
					&traversable_mock{value: 5, children: nil},
				}},
				&traversable_mock{value: 3, children: []Traversable{
					&traversable_mock{value: 6, children: nil},
					&traversable_mock{value: 7, children: nil},
				}},
			},
		},
		action: func(n Traversable) Break {
			node, _ := n.(*traversable_mock)
			result = append(result, node.value)

			return false
		},
	}

	dfs.PreOrderTraverse(dfs.root)

	if !reflect.DeepEqual(result, []int{1, 2, 4, 5, 3, 6, 7}) {
		t.Errorf("traversal order should be 1 2 4 5 3 6 7, instead got: %v", result)
	}
}

func TestDFS_PostOrderTraverse(t *testing.T) {
	result := []int{}
	dfs := &DFS{
		root: &traversable_mock{
			value: 1,
			children: []Traversable{
				&traversable_mock{value: 2, children: []Traversable{
					&traversable_mock{value: 4, children: nil},
					&traversable_mock{value: 5, children: nil},
				}},
				&traversable_mock{value: 3, children: []Traversable{
					&traversable_mock{value: 6, children: nil},
					&traversable_mock{value: 7, children: nil},
				}},
			},
		},
		action: func(n Traversable) Break {
			node, _ := n.(*traversable_mock)
			result = append(result, node.value)

			return false
		},
	}

	dfs.PostOrderTraverse(dfs.root)

	if !reflect.DeepEqual(result, []int{4, 5, 2, 6, 7, 3, 1}) {
		t.Errorf("traversal order should be 4 5 2 6 7 3 1, instead got: %v", result)
	}
}
