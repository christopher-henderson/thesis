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

	Row int
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

	type __1_StackEntry struct {
		Parent   MazeNode
		Children chan MazeNode
	}

	// Engine initialization.
	__1_stack := make([]__1_StackEntry, 0)
	__1_solution := make([]MazeNode, 0)
	__1_root := MazeNode{-1, 1}

	var __1_c chan MazeNode
	////////////////////////////////////////////////////////
	// USERLAND
	for {
		// PARAMETER BINDINGS
		parent := __1_root
		/////////////////////
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
		__1_c = c
		goto __1_END_INIT_CHILDREN

	}
	////////////////////////////////////////////////////////
__1_END_INIT_CHILDREN:

	var __1_candidate MazeNode
	var __1_ok bool
	var __1_se __1_StackEntry
	for {
		if __1_candidate, __1_ok = <-__1_c; !__1_ok {
			if len(__1_stack) == 0 {
				break
			}
			__1_solution = __1_solution[:len(__1_solution)-1]
			__1_se = __1_stack[len(__1_stack)-1]
			__1_stack = __1_stack[:len(__1_stack)-1]
			__1_root = __1_se.Parent
			__1_c = __1_se.Children
			continue
		}
		var __1_reject bool
		////////////////////////////////////////////////////////
		// USERLAND - REJECT
		for {
			// PARAMETER BINDINGS
			candidate := __1_candidate
			solution := __1_solution
			/////////////////////
			log.Println(candidate, solution)
			__1_reject = false
			goto __1_END_REJECT

		}
		////////////////////////////////////////////////////////
	__1_END_REJECT:
		if __1_reject {
			continue
		}
		__1_solution = append(__1_solution, __1_candidate)
		var __1_accept bool
		////////////////////////////////////////////////////////
		// USERLAND - ACCEPT
		for {
			// PARAMETER BINDINGS
			solution := __1_solution
			/////////////////////
			log.Println(solution)
			if len(solution) == 1 {
				__1_accept = false
				goto __1_END_ACCEPT
			}
			final := solution[len(solution)-1]
			__1_accept = NSborder(maze, final) || ESborder(maze, final)
			goto __1_END_ACCEPT

		}
		////////////////////////////////////////////////////////
	__1_END_ACCEPT:
		if __1_accept {
			log.Println(__1_solution)
			__1_solution = __1_solution[:len(__1_solution)-1]
			continue
		}
		__1_stack = append(__1_stack, __1_StackEntry{__1_root, __1_c})
		__1_root = __1_candidate
		////////////////////////////////////////////////////////
		// USERLAND - CHILDREN
		for {
			// PARAMETER BINDINGS
			parent := __1_root
			/////////////////////
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
			__1_c = c
			goto __1_END_CHILDREN

		}
		////////////////////////////////////////////////////////
	__1_END_CHILDREN:
	}
}
