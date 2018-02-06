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
			__aecbfeab_USER_children := func(parent Queen) chan Queen {
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
			__aecbfeab_USER_accept := func(solution []Queen) bool {
				if len(solution) == N {
					winners++
				}
				return len(solution) == N
			}
			// User declaration.
			__aecbfeab_USER_reject := func(candidate Queen, solution []Queen) bool {
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
			type __aecbfeab_StackEntry struct {
				Parent   Queen
				Children chan Queen
			}
			/////////////// Engine initialization.
			// Stack of Parent:Chidren pairs.
			__aecbfeab_stack := make([]__aecbfeab_StackEntry, 0)
			// Solution thus far.
			__aecbfeab_solution := make([]Queen, 0)
			// Current root under consideration.
			__aecbfeab_root := Queen{0, 0}
			// Current candidate under consideration.
			var __aecbfeab_candidate Queen
			// Holds a Stack Entry and popping.
			var __aecbfeab_se __aecbfeab_StackEntry
			// Generic bool holder.
			var __aecbfeab_ok bool

			// Begin search.
			__aecbfeab_c := __aecbfeab_USER_children(__aecbfeab_root)
			for {
				if __aecbfeab_candidate, __aecbfeab_ok = <-__aecbfeab_c; !__aecbfeab_ok {
					if len(__aecbfeab_stack) == 0 {
						break
					}
					__aecbfeab_solution = __aecbfeab_solution[:len(__aecbfeab_solution)-1]
					__aecbfeab_se = __aecbfeab_stack[len(__aecbfeab_stack)-1]
					__aecbfeab_stack = __aecbfeab_stack[:len(__aecbfeab_stack)-1]
					__aecbfeab_root = __aecbfeab_se.Parent
					__aecbfeab_c = __aecbfeab_se.Children
					continue
				}
				__aecbfeab_reject := __aecbfeab_USER_reject(__aecbfeab_candidate, __aecbfeab_solution)
				if __aecbfeab_reject {
					continue
				}
				__aecbfeab_solution = append(__aecbfeab_solution, __aecbfeab_candidate)
				__aecbfeab_accept := __aecbfeab_USER_accept(__aecbfeab_solution)
				if __aecbfeab_accept {
					log.Println(__aecbfeab_solution)
					__aecbfeab_solution = __aecbfeab_solution[:len(__aecbfeab_solution)-1]
					continue
				}
				__aecbfeab_stack = append(__aecbfeab_stack, __aecbfeab_StackEntry{__aecbfeab_root, __aecbfeab_c})
				__aecbfeab_root = __aecbfeab_candidate
				__aecbfeab_c = __aecbfeab_USER_children(__aecbfeab_root)
			}
		}
		log.Println(winners)
		log.Println(time.Now().Sub(start))
	}
}
