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
	Column int
	Row    int
}

// North-Sounth border
func NSborder(maze [][]int, node MazeNode) bool {
	numRows := len(maze)
	return node.Row == 0 || node.Row == numRows - 1
}

// East-West border
func ESborder(maze [][]int, node MazeNode) bool {
	numColumns := len(maze[0])
	return node.Column == 0 || node.Column == numColumns - 1
}

func main() {
	maze := simple
	search from MazeNode{-1, 1} (MazeNode) {
		children parent:
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
		accept solution:
			if len(solution) == 1 {
				return false
			}
			final := solution[len(solution) - 1]
			return NSborder(maze, final) || ESborder(maze, final)
		reject candidate, solution:
			return false
		add candidate, solution:
			log.Println("WHOA NOW, ADDING")
		remove candidate, solution:
			log.Println("Guess it didn't work out")
	}
}
