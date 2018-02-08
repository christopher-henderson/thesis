package main

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
	Valid  bool
	Row    int
	Column int
	maze   [][]int
}

// North-Sounth border
func (mn *MazeNode) NSborder() bool {
	return mn.Row == 0 || mn.Row == len(mn.maze)-1
}

// East-West border
func (mn *MazeNode) ESborder() bool {
	return mn.Column == 0 || mn.Column == len(mn.maze[0])-1
}

func MakeNodeMap(maze [][]int) [][]*MazeNode {
	m := make([][]*MazeNode, len(maze))
	for i, r := range maze {
		m[i] = make([]*MazeNode, len(maze[i]))
		for j, c := range r {
			m[i][j] = &MazeNode{c == 1, i, j, maze}
		}
	}
	return m
}

func main() {
	maze := simple
	nodeMap := MakeNodeMap(maze)
	seen := make(map[*MazeNode]bool)
	// 'if' used to scope this entire engine.
	if true {
		// User CHILDREN declaration.
		__858bf52c_USER_children := func(parent *MazeNode) chan *MazeNode {
			c := make(chan *MazeNode, 0)
			go func() {
				defer close(c)
				north := []int{parent.Row - 1, parent.Column}
				south := []int{parent.Row + 1, parent.Column}
				east := []int{parent.Row, parent.Column - 1}
				west := []int{parent.Row, parent.Column + 1}
				if inbounds(north, maze) {
					child := nodeMap[north[0]][north[1]]
					if _, ok := seen[child]; !ok {
						c <- child
					}
				}
				if inbounds(south, maze) {
					child := nodeMap[south[0]][south[1]]
					if _, ok := seen[child]; !ok {
						c <- child
					}
				}
				if inbounds(east, maze) {
					child := nodeMap[east[0]][east[1]]
					if _, ok := seen[child]; !ok {
						c <- child
					}
				}
				if inbounds(west, maze) {
					child := nodeMap[west[0]][west[1]]
					if _, ok := seen[child]; !ok {
						c <- child
					}
				}
			}()
			return c
		}
		// User ACCEPT declaration.
		__858bf52c_USER_accept := func(solution []*MazeNode) bool {
			if len(solution) == 1 {
				// this is the entry point
				return false
			}
			final := solution[len(solution)-1]
			return final.Valid && (final.NSborder() || final.ESborder())
		}
		// User REJECT declaration.
		__858bf52c_USER_reject := func(candidate *MazeNode, solution []*MazeNode) bool {
			return false
		}
		// Parent:Children PODO meant for stack management.
		type __858bf52c_StackEntry struct {
			Parent   *MazeNode
			Children chan *MazeNode
		}
		/////////////// Engine initialization.
		// Stack of Parent:Chidren pairs.
		__858bf52c_stack := make([]__858bf52c_StackEntry, 0)
		// Solution thus far.
		__858bf52c_solution := make([]*MazeNode, 0)
		// Current root under consideration.
		__858bf52c_root := &MazeNode{false, -1, 1, maze}
		// Current candidate under consideration.
		var __858bf52c_candidate *MazeNode
		// Holds a StackEntry.
		var __858bf52c_stackEntry __858bf52c_StackEntry
		// Generic boolean variable
		var __858bf52c_ok bool
		/////////////// Begin search.
		__858bf52c_children := __858bf52c_USER_children(__858bf52c_root)
		for {
			if __858bf52c_candidate, __858bf52c_ok = <-__858bf52c_children; !__858bf52c_ok {
				// This node has no further children.
				if len(__858bf52c_stack) == 0 {
					// Algorithm termination. No further nodes in the stack.
					break
				}
				// With no valid children left, we pop the latest node from the solution.
				__858bf52c_solution = __858bf52c_solution[:len(__858bf52c_solution)-1]
				// Pop from the stack. Broken into two steps:
				// 	1. Get final element.
				//	2. Resize the stack.
				__858bf52c_stackEntry = __858bf52c_stack[len(__858bf52c_stack)-1]
				__858bf52c_stack = __858bf52c_stack[:len(__858bf52c_stack)-1]
				// Extract root and candidate fields from the StackEntry.
				__858bf52c_root = __858bf52c_stackEntry.Parent
				__858bf52c_children = __858bf52c_stackEntry.Children
				continue
			}
			// Ask the user if we should reject this candidate.
			__858bf52c_reject := __858bf52c_USER_reject(__858bf52c_candidate, __858bf52c_solution)
			if __858bf52c_reject {
				// Rejected candidate.
				continue
			}
			// Append the candidate to the solution.
			__858bf52c_solution = append(__858bf52c_solution, __858bf52c_candidate)
			// Ask the user if we should accept this solution.
			__858bf52c_accept := __858bf52c_USER_accept(__858bf52c_solution)
			if __858bf52c_accept {
				// Accepted solution.
				// Pop from the solution thus far and continue on with the next child.
				__858bf52c_solution = __858bf52c_solution[:len(__858bf52c_solution)-1]
				continue
			}
			// Push the current root to the stack.
			__858bf52c_stack = append(__858bf52c_stack, __858bf52c_StackEntry{__858bf52c_root, __858bf52c_children})
			// Make the candidate the new root.
			__858bf52c_root = __858bf52c_candidate
			// Get the new root's children channel.
			__858bf52c_children = __858bf52c_USER_children(__858bf52c_root)
		}
	}
}
