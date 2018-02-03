package main

import "log"

var N = 8

type Queen struct {
	Column int
	Row    int
}

func reject(candidate Queen, solution []Queen) bool {
	row, column := candidate.Row, candidate.Column
	for _, q := range solution {
		r, c := q.Row, q.Column
		if row == r ||
			column == c ||
			row+column == r+c ||
			row-column == r-c {
			return true
		}
	}
	return false
}

func children(parent Queen) chan Queen {
	column := parent.Column + 1
	c := make(chan Queen, 0)
	go func() {
		defer close(c)
		for r := 1; r < N+1; r++ {
			c <- Queen{column, r}
		}
	}()
	return c
}

func accept(solution []Queen) bool {
	return len(solution) == N
}

type RootStackEntry struct {
	Parent   Queen
	Children chan Queen
}

type RootStack []RootStackEntry

func rpush(s RootStack, e RootStackEntry) RootStack {
	return append(s, e)
}

func rpop(s RootStack) (RootStackEntry, RootStack) {
	if len(s) == 0 {
		return RootStackEntry{}, s
	}
	return s[len(s)-1], s[:len(s)-1]
}

type SolutionStack []Queen

func spush(s SolutionStack, e Queen) SolutionStack {
	return append(s, e)
}

func spop(s SolutionStack) (Queen, SolutionStack) {
	if len(s) == 0 {
		return Queen{}, s
	}
	return s[len(s)-1], s[:len(s)-1]
}

func backtrack(root Queen) {
	stack := RootStack{}
	solution := SolutionStack{}

	c := children(root)

	var candidate Queen
	var ok bool
	var rse RootStackEntry
	for {
		log.Println(solution)
		if candidate, ok = <-c; !ok {
			_, solution = spop(solution)
			if len(stack) == 0 {
				log.Println("NO ANSWER.")
				return
			}
			rse, stack = rpop(stack)
			root = rse.Parent
			c = rse.Children
			continue
		}
		if reject(candidate, solution) {
			continue
		}
		solution = spush(solution, candidate)
		if accept(solution) {
			log.Println("WINNER")
			log.Println(solution)
			return
		}
		stack = rpush(stack, RootStackEntry{root, c})
		root = candidate
		c = children(root)
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	backtrack(Queen{0, 0})
}

// 1, 0, 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 0, 1, 0, 0, 0,
// 0, 0, 0, 0, 0, 0, 0, 0,
// 0, 1, 0, 0, 0, 0, 0, 0,
// 0, 0, 0, 1, 0, 0, 0, 0,
// 0, 0, 0, 0, 0, 1, 0, 0,
// 0, 0, 1, 0, 0, 0, 0, 0,
// [
// 1, 0, 0, 0,
// 0, 0, 0, 0,
// 0, 1, 0, 0,
// 0, 0, 0, 0,
// ]
