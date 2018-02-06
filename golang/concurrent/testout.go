package main

import (
	"log"
	"time"
)

type Queen struct {
	Column int
	Row    int
}

func main() {
	for N := 0; N < 9; N++ {
		start := time.Now()
		winners := 0
		shutdown := make(chan int, 0)

		// 'if' used to scope this entire engine.
		if true {
			// User declaration.
			__9abf22eb_USER_children := func(parent Queen) chan Queen {
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
			// User declaration.
			__9abf22eb_USER_accept := func(solution []Queen) bool {
				if len(solution) == N {
					winners++
				}
				return len(solution) == N

			}
			// User declaration.
			__9abf22eb_USER_reject := func(candidate Queen, solution []Queen) bool {
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
			// Parent:Children PODO meant for stack management.
			type __9abf22eb_StackEntry struct {
				Parent   Queen
				Children chan Queen
			}
			/////////////// Engine initialization.
			// Stack of Parent:Chidren pairs.
			__9abf22eb_stack := make([]__9abf22eb_StackEntry, 0)
			// Solution thus far.
			__9abf22eb_solution := make([]Queen, 0)
			// Current root under consideration.
			__9abf22eb_root := Queen{0, 0}
			// Current candidate under consideration.
			var __9abf22eb_candidate Queen
			// Holds a Stack Entry and popping.
			var __9abf22eb_se __9abf22eb_StackEntry
			// Generic bool holder.
			var __9abf22eb_ok bool

			// Begin search.
			__9abf22eb_c := __9abf22eb_USER_children(__9abf22eb_root)
			for {
				if __9abf22eb_candidate, __9abf22eb_ok = <-__9abf22eb_c; !__9abf22eb_ok {
					if len(__9abf22eb_stack) == 0 {
						break
					}
					__9abf22eb_solution = __9abf22eb_solution[:len(__9abf22eb_solution)-1]
					__9abf22eb_se = __9abf22eb_stack[len(__9abf22eb_stack)-1]
					__9abf22eb_stack = __9abf22eb_stack[:len(__9abf22eb_stack)-1]
					__9abf22eb_root = __9abf22eb_se.Parent
					__9abf22eb_c = __9abf22eb_se.Children
					continue
				}
				__9abf22eb_reject := __9abf22eb_USER_reject(__9abf22eb_candidate, __9abf22eb_solution)
				if __9abf22eb_reject {
					continue
				}
				__9abf22eb_solution = append(__9abf22eb_solution, __9abf22eb_candidate)
				__9abf22eb_accept := __9abf22eb_USER_accept(__9abf22eb_solution)
				if __9abf22eb_accept {
					log.Println(__9abf22eb_solution)
					__9abf22eb_solution = __9abf22eb_solution[:len(__9abf22eb_solution)-1]
					continue
				}
				__9abf22eb_stack = append(__9abf22eb_stack, __9abf22eb_StackEntry{__9abf22eb_root, __9abf22eb_c})
				__9abf22eb_root = __9abf22eb_candidate
				__9abf22eb_c = __9abf22eb_USER_children(__9abf22eb_root)
			}
		}
		close(shutdown)
		log.Println(winners)
		log.Println(time.Now().Sub(start))
	}

}
