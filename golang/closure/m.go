package main

import "log"

var simple = [][]int{{0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
}

var cycle = [][]int{{0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 0, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
}

var multipleEntrancesExits = [][]int{{0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1},
	{0, 1, 0, 1, 0, 0},
	{0, 1, 0, 1, 0, 0},
}

type MazeNode struct {
	Column int
	Row    int
}

// North-Sounth border
func NSborder(maze [][]int, node MazeNode) bool {
	numRows := len(maze)
	return node.Row == 0 || node.Row == numRows-1
}

// East-West border
func ESborder(maze [][]int, node MazeNode) bool {
	numColumns := len(maze[0])
	return node.Column == 0 || node.Column == numColumns-1
}

func main() {
	maze := simple

	// 'if' used to scope this entire engine.
	if true {
		// User declaration.
		__9584c2e7_USER_children := func(parent MazeNode) chan MazeNode {
			log.Println(parent)
			c := make(chan MazeNode, 0)
			children := []MazeNode{MazeNode{parent.Column - 1, parent.Row},
				MazeNode{parent.Column, parent.Row - 1},
				MazeNode{parent.Column + 1, parent.Row},
				MazeNode{parent.Column, parent.Row + 1}}
			go func() {
				for _, child := range children {
					if child.Column < 0 || child.Column >= len(maze[0]) || child.Row < 0 || child.Row >= len(maze) || maze[child.Row][child.Column] == 0 {
						continue
					}
					c <- child
				}
				close(c)
			}()
			return c
		}
		// User declaration.
		__9584c2e7_USER_accept := func(solution []MazeNode) bool {
			if len(solution) == 1 {
				return false
			}
			final := solution[len(solution)-1]
			return NSborder(maze, final) || ESborder(maze, final)
		}
		// User declaration.
		__9584c2e7_USER_reject := func(candidate MazeNode, solution []MazeNode) bool {
			return false
		}
		// Parent:Children PODO meant for stack management.
		type __9584c2e7_StackEntry struct {
			Parent   MazeNode
			Children chan MazeNode
		}
		/////////////// Engine initialization.
		// Stack of Parent:Chidren pairs.
		__9584c2e7_stack := make([]__9584c2e7_StackEntry, 0)
		// Solution thus far.
		__9584c2e7_solution := make([]MazeNode, 0)
		// Current root under consideration.
		__9584c2e7_root := MazeNode{-1, 1}
		// Current candidate under consideration.
		var __9584c2e7_candidate MazeNode
		// Holds a Stack Entry and popping.
		var __9584c2e7_se __9584c2e7_StackEntry
		// Generic bool holder.
		var __9584c2e7_ok bool

		// Begin search.
		__9584c2e7_c := __9584c2e7_USER_children(__9584c2e7_root)
		for {
			if __9584c2e7_candidate, __9584c2e7_ok = <-__9584c2e7_c; !__9584c2e7_ok {
				if len(__9584c2e7_stack) == 0 {
					break
				}
				__9584c2e7_solution = __9584c2e7_solution[:len(__9584c2e7_solution)-1]
				__9584c2e7_se = __9584c2e7_stack[len(__9584c2e7_stack)-1]
				__9584c2e7_stack = __9584c2e7_stack[:len(__9584c2e7_stack)-1]
				__9584c2e7_root = __9584c2e7_se.Parent
				__9584c2e7_c = __9584c2e7_se.Children
				continue
			}
			__9584c2e7_reject := __9584c2e7_USER_reject(__9584c2e7_candidate, __9584c2e7_solution)
			if __9584c2e7_reject {
				continue
			}
			__9584c2e7_solution = append(__9584c2e7_solution, __9584c2e7_candidate)
			__9584c2e7_accept := __9584c2e7_USER_accept(__9584c2e7_solution)
			if __9584c2e7_accept {
				log.Println(__9584c2e7_solution)
				__9584c2e7_solution = __9584c2e7_solution[:len(__9584c2e7_solution)-1]
				continue
			}
			__9584c2e7_stack = append(__9584c2e7_stack, __9584c2e7_StackEntry{__9584c2e7_root, __9584c2e7_c})
			__9584c2e7_root = __9584c2e7_candidate
			__9584c2e7_c = __9584c2e7_USER_children(__9584c2e7_root)
		}
	}
}
