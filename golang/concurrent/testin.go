package main

import (
	"time"
	"log"
)

type Queen struct {
	Column int
	Row    int
}

func main() {
	N := 8
	start := time.Now()
	search from Queen{0,0} (Queen) {
		concurrent
		children node:
			column := node.Column + 1
			c := make(chan Queen, 0)
			// If the parent is in the final column
			// then there are no children.
			if column > N {
				close(c)
				return c
			}
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					c <- Queen{column, r}
				}
			}()
			return c
		accept solution:
			if len(solution) == N {
				log.Println(solution)
				return true
			}
			return false
		reject node, solution:
			row, column := node.Row, node.Column
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
	log.Println(time.Now().Sub(start))
}