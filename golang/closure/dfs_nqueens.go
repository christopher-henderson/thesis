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
	N := 8
	start := time.Now()
	// 'if' used to scope this entire engine.
	if true {
		// User declaration.
		__23d7d4eb_USER_children := func(parent Queen) chan Queen {
			column := parent.Column + 1
			c := make(chan Queen, 0)
			if column > N {
				close(c)
				return c
			}
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					q := Queen{column, r}
					log.Println(q)
					c <- q
				}
			}()
			return c
		}
		// Parent:Children PODO meant for stack management.
		type __23d7d4eb_StackEntry struct {
			Parent   Queen
			Children chan Queen
		}
		/////////////// Engine initialization.
		// Stack of Parent:Chidren pairs.
		__23d7d4eb_stack := make([]__23d7d4eb_StackEntry, 0)
		// Solution thus far.
		// __23d7d4eb_solution := make([]Queen, 0)
		// Current root under consideration.
		__23d7d4eb_root := Queen{0, 0}
		// Current candidate under consideration.
		var __23d7d4eb_candidate Queen
		// Holds a Stack Entry and popping.
		var __23d7d4eb_stackEntry __23d7d4eb_StackEntry
		// Generic bool holder.
		var __23d7d4eb_ok bool

		// Begin search.
		__23d7d4eb_children := __23d7d4eb_USER_children(__23d7d4eb_root)
		for {
			if __23d7d4eb_candidate, __23d7d4eb_ok = <-__23d7d4eb_children; !__23d7d4eb_ok {
				if len(__23d7d4eb_stack) == 0 {
					break
				}
				// __23d7d4eb_solution = __23d7d4eb_solution[:len(__23d7d4eb_solution)-1]
				__23d7d4eb_stackEntry = __23d7d4eb_stack[len(__23d7d4eb_stack)-1]
				__23d7d4eb_stack = __23d7d4eb_stack[:len(__23d7d4eb_stack)-1]
				__23d7d4eb_root = __23d7d4eb_stackEntry.Parent
				__23d7d4eb_children = __23d7d4eb_stackEntry.Children
				continue
			}
			// __23d7d4eb_solution = append(__23d7d4eb_solution, __23d7d4eb_candidate)
			__23d7d4eb_stack = append(__23d7d4eb_stack, __23d7d4eb_StackEntry{__23d7d4eb_root, __23d7d4eb_children})
			__23d7d4eb_root = __23d7d4eb_candidate
			__23d7d4eb_children = __23d7d4eb_USER_children(__23d7d4eb_root)
		}
	}
	log.Println(time.Now().Sub(start))
}
