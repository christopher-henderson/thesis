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
	// for N := 0; N < 13; N++ {
	N := 12
	start := time.Now()
	winners := 0
	// 'if' used to scope this entire engine.
	if true {
		// User CHILDREN declaration.
		__5d9cc7ed_USER_children := func(parent Queen) chan Queen {
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
		}
		// User ACCEPT declaration.
		__5d9cc7ed_USER_accept := func(solution []Queen) bool {
			if len(solution) == N {
				// Print it, append it to a slice of solutions,
				// whatever you want here.
				// log.Println(solution)
				// Keeping track of number of solutions
				// for quick verification.
				winners++
				return true
			}
			return false
		}
		// User REJECT declaration.
		__5d9cc7ed_USER_reject := func(candidate Queen, solution []Queen) bool {
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
		type __5d9cc7ed_StackEntry struct {
			Parent   Queen
			Children chan Queen
		}
		/////////////// Engine initialization.
		// Stack of Parent:Chidren pairs.
		__5d9cc7ed_stack := make([]__5d9cc7ed_StackEntry, 0)
		// Solution thus far.
		__5d9cc7ed_solution := make([]Queen, 0)
		// Current root under consideration.
		__5d9cc7ed_root := Queen{0, 0}
		// Current candidate under consideration.
		var __5d9cc7ed_candidate Queen
		// Holds a StackEntry.
		var __5d9cc7ed_stackEntry __5d9cc7ed_StackEntry
		// Generic boolean variable
		var __5d9cc7ed_ok bool
		/////////////// Begin search.
		__5d9cc7ed_children := __5d9cc7ed_USER_children(__5d9cc7ed_root)
		for {
			if __5d9cc7ed_candidate, __5d9cc7ed_ok = <-__5d9cc7ed_children; !__5d9cc7ed_ok {
				// This node has no further children.
				if len(__5d9cc7ed_stack) == 0 {
					// Algorithm termination. No further nodes in the stack.
					break
				}
				// With no valid children left, we pop the latest node from the solution.
				__5d9cc7ed_solution = __5d9cc7ed_solution[:len(__5d9cc7ed_solution)-1]
				// Pop from the stack. Broken into two steps:
				// 	1. Get final element.
				//	2. Resize the stack.
				__5d9cc7ed_stackEntry = __5d9cc7ed_stack[len(__5d9cc7ed_stack)-1]
				__5d9cc7ed_stack = __5d9cc7ed_stack[:len(__5d9cc7ed_stack)-1]
				// Extract root and candidate fields from the StackEntry.
				__5d9cc7ed_root = __5d9cc7ed_stackEntry.Parent
				__5d9cc7ed_children = __5d9cc7ed_stackEntry.Children
				continue
			}
			// Ask the user if we should reject this candidate.
			__5d9cc7ed_reject := __5d9cc7ed_USER_reject(__5d9cc7ed_candidate, __5d9cc7ed_solution)
			if __5d9cc7ed_reject {
				// Rejected candidate.
				continue
			}
			// Append the candidate to the solution.
			__5d9cc7ed_solution = append(__5d9cc7ed_solution, __5d9cc7ed_candidate)
			// Ask the user if we should accept this solution.
			__5d9cc7ed_accept := __5d9cc7ed_USER_accept(__5d9cc7ed_solution)
			if __5d9cc7ed_accept {
				// Accepted solution.
				// Pop from the solution thus far and continue on with the next child.
				__5d9cc7ed_solution = __5d9cc7ed_solution[:len(__5d9cc7ed_solution)-1]
				continue
			}
			// Push the current root to the stack.
			__5d9cc7ed_stack = append(__5d9cc7ed_stack, __5d9cc7ed_StackEntry{__5d9cc7ed_root, __5d9cc7ed_children})
			// Make the candidate the new root.
			__5d9cc7ed_root = __5d9cc7ed_candidate
			// Get the new root's children channel.
			__5d9cc7ed_children = __5d9cc7ed_USER_children(__5d9cc7ed_root)
		}
	}
	log.Println(winners)
	log.Println(time.Now().Sub(start))
	// }
}
