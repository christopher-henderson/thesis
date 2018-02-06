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
		__f0a76785_USER_children := func(parent MazeNode) chan MazeNode {
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
		__f0a76785_USER_accept := func(solution []MazeNode) bool {
			if len(solution) == 1 {
				return false
			}
			final := solution[len(solution)-1]
			return NSborder(maze, final) || ESborder(maze, final)
		}
		// User declaration.
		__f0a76785_USER_reject := func(candidate MazeNode, solution []MazeNode) bool {
			return false
		}
		// Parent:Children PODO meant for stack management.
		type __f0a76785_StackEntry struct {
			Parent   MazeNode
			Children chan MazeNode
		}
		/////////////// Engine initialization.
		// Stack of Parent:Chidren pairs.
		__f0a76785_stack := make([]__f0a76785_StackEntry, 0)
		// Solution thus far.
		__f0a76785_solution := make([]MazeNode, 0)
		// Current root under consideration.
		__f0a76785_root := MazeNode{-1, 1}
		// Current candidate under consideration.
		var __f0a76785_candidate MazeNode
		// Holds a Stack Entry and popping.
		var __f0a76785_se __f0a76785_StackEntry
		// Generic bool holder.
		var __f0a76785_ok bool

		// Begin search.
		__f0a76785_c := __f0a76785_USER_children(__f0a76785_root)
		for {
			if __f0a76785_candidate, __f0a76785_ok = <-__f0a76785_c; !__f0a76785_ok {
				if len(__f0a76785_stack) == 0 {
					break
				}
				__f0a76785_solution = __f0a76785_solution[:len(__f0a76785_solution)-1]
				__f0a76785_se = __f0a76785_stack[len(__f0a76785_stack)-1]
				__f0a76785_stack = __f0a76785_stack[:len(__f0a76785_stack)-1]
				__f0a76785_root = __f0a76785_se.Parent
				__f0a76785_c = __f0a76785_se.Children
				continue
			}
			__f0a76785_reject := __f0a76785_USER_reject(__f0a76785_candidate, __f0a76785_solution)
			if __f0a76785_reject {
				continue
			}
			__f0a76785_solution = append(__f0a76785_solution, __f0a76785_candidate)
			__f0a76785_accept := __f0a76785_USER_accept(__f0a76785_solution)
			if __f0a76785_accept {
				log.Println(__f0a76785_solution)
				__f0a76785_solution = __f0a76785_solution[:len(__f0a76785_solution)-1]
				continue
			}
			__f0a76785_stack = append(__f0a76785_stack, __f0a76785_StackEntry{__f0a76785_root, __f0a76785_c})
			__f0a76785_root = __f0a76785_candidate
			__f0a76785_c = __f0a76785_USER_children(__f0a76785_root)
		}
	}
}
