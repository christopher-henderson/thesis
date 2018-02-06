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

		// 'if' used to scope this entire engine.
		if true {
			// User declaration.
			__d3708415_USER_children := func(parent Queen) chan Queen {
				column := parent.Column + 1
				c := make(chan Queen, 0)
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
			}
			// User declaration.
			__d3708415_USER_accept := func(solution []Queen) bool {
				if len(solution) == N {
					winners++
				}
				return len(solution) == N
			}
			// User declaration.
			__d3708415_USER_reject := func(candidate Queen, solution []Queen) bool {
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
			type __d3708415_StackEntry struct {
				Parent   Queen
				Children chan Queen
			}
			/////////////// Engine initialization.
			// Stack of Parent:Chidren pairs.
			__d3708415_stack := make([]__d3708415_StackEntry, 0)
			// Solution thus far.
			__d3708415_solution := make([]Queen, 0)
			// Current root under consideration.
			__d3708415_root := Queen{0, 0}
			// Current candidate under consideration.
			var __d3708415_candidate Queen
			// Holds a Stack Entry and popping.
			var __d3708415_se __d3708415_StackEntry
			// Generic bool holder.
			var __d3708415_ok bool

			// Begin search.
			__d3708415_c := __d3708415_USER_children(__d3708415_root)
			for {
				if __d3708415_candidate, __d3708415_ok = <-__d3708415_c; !__d3708415_ok {
					if len(__d3708415_stack) == 0 {
						break
					}
					__d3708415_solution = __d3708415_solution[:len(__d3708415_solution)-1]
					__d3708415_se = __d3708415_stack[len(__d3708415_stack)-1]
					__d3708415_stack = __d3708415_stack[:len(__d3708415_stack)-1]
					__d3708415_root = __d3708415_se.Parent
					__d3708415_c = __d3708415_se.Children
					continue
				}
				__d3708415_reject := __d3708415_USER_reject(__d3708415_candidate, __d3708415_solution)
				if __d3708415_reject {
					continue
				}
				__d3708415_solution = append(__d3708415_solution, __d3708415_candidate)
				__d3708415_accept := __d3708415_USER_accept(__d3708415_solution)
				if __d3708415_accept {
					log.Println(__d3708415_solution)
					__d3708415_solution = __d3708415_solution[:len(__d3708415_solution)-1]
					continue
				}
				__d3708415_stack = append(__d3708415_stack, __d3708415_StackEntry{__d3708415_root, __d3708415_c})
				__d3708415_root = __d3708415_candidate
				__d3708415_c = __d3708415_USER_children(__d3708415_root)
			}
		}
		log.Println(winners)
		log.Println(time.Now().Sub(start))
	}
}
