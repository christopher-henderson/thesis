package main

type Queen struct {
	Column int
	Row    int
}

func main() {
	for N := 0; N < 9; N ++ {
		start := time.Now()
		winners := 0
		shutdown := make(chan int, 0)
		search from Queen{0,0} (Queen) {
			accept solution:
				if len(solution) == N {
					winners++
				}
				return len(solution) == N
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
			children parent:
				column := parent.Column + 1
				c := make(chan Queen, 0)
				if column > N {
					close(c)
					return c
				}
				go func() {
					defer close(c)
					for r := 1; r < N+1; r++ {
						select {
							case c <- Queen{column, r}:
							case <-shutdown:
								break
						}
					}
				}()
				return c
		}
		close(shutdown)
		log.Println(winners)
		log.Println(time.Now().Sub(start))
	}
	
	
}