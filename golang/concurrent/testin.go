package main

import "sync"

type Queen struct {
	Column int
	Row    int
}

func main() {
	N := 12
	start := time.Now()
	winners := 0
	l := sync.Mutex{}
	search from Queen{0,0} (Queen) {
		concurrent
		children parent:
			column := parent.Column + 1
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
				// Print it, append it to a slice of solutions,
				// whatever you want here.
				// log.Println(solution)
				// Keeping track of number of solutions
				// for quick verification.
				l.Lock()
				winners++
				l.Unlock()
				return true
			}
			return false
		reject candidate, solution:
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
	log.Println(winners)
	log.Println(time.Now().Sub(start))
}