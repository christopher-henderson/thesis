package main

import (
	"log"
	"time"
)

type Queen struct {
	Column int

	Row int
}

func main() {

	N := 8

	winners := 0

	shutdown := make(chan int, 0)

	type __1_StackEntry struct {
		Parent   Queen
		Children chan Queen
	}

	// Engine initialization.
	__1_stack := make([]__1_StackEntry, 0)
	__1_solution := make([]Queen, 0)
	__1_root := Queen{0, 0}

	var __1_c chan Queen
	////////////////////////////////////////////////////////
	// USERLAND
	for {
		// PARAMETER BINDINGS
		parent := __1_root
		/////////////////////
		column := parent.Column + 1
		c := make(chan Queen, 0)
		if column > N {
			close(c)
			__1_c = c
			goto __1_END_INIT_CHILDREN
		}
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				select {
				case c <- Queen{column, r}:
				case <-shutdown:
					log.Println("THIS WOULDA BEEN A LEAK")
					break
				}
			}
		}()
		__1_c = c
		goto __1_END_INIT_CHILDREN

	}
	////////////////////////////////////////////////////////
__1_END_INIT_CHILDREN:

	var __1_candidate Queen
	var __1_ok bool
	var __1_se __1_StackEntry
	for {
		if __1_candidate, __1_ok = <-__1_c; !__1_ok {
			if len(__1_stack) == 0 {
				break
			}
			__1_solution = __1_solution[:len(__1_solution)-1]
			__1_se = __1_stack[len(__1_stack)-1]
			__1_stack = __1_stack[:len(__1_stack)-1]
			__1_root = __1_se.Parent
			__1_c = __1_se.Children
			continue
		}
		var __1_reject bool
		////////////////////////////////////////////////////////
		// USERLAND - REJECT
		for {
			// PARAMETER BINDINGS
			candidate := __1_candidate
			solution := __1_solution
			/////////////////////
			row, column := candidate.Row, candidate.Column
			for _, q := range solution {
				r, c := q.Row, q.Column
				if row == r ||
					column == c ||
					row+column == r+c ||
					row-column == r-c {
					__1_reject = true
					goto __1_END_REJECT
				}
			}
			__1_reject = false
			goto __1_END_REJECT

		}
		////////////////////////////////////////////////////////
	__1_END_REJECT:
		if __1_reject {
			continue
		}
		__1_solution = append(__1_solution, __1_candidate)
		var __1_accept bool
		////////////////////////////////////////////////////////
		// USERLAND - ACCEPT
		for {
			// PARAMETER BINDINGS
			solution := __1_solution
			/////////////////////
			if len(solution) == N {
				winners++
			}
			__1_accept = len(solution) == N
			goto __1_END_ACCEPT

		}
		////////////////////////////////////////////////////////
	__1_END_ACCEPT:
		if __1_accept {
			log.Println(__1_solution)
			__1_solution = __1_solution[:len(__1_solution)-1]
			continue
		}
		__1_stack = append(__1_stack, __1_StackEntry{__1_root, __1_c})
		__1_root = __1_candidate
		////////////////////////////////////////////////////////
		// USERLAND - CHILDREN
		for {
			// PARAMETER BINDINGS
			parent := __1_root
			/////////////////////
			column := parent.Column + 1
			c := make(chan Queen, 0)
			if column > N {
				close(c)
				__1_c = c
				goto __1_END_CHILDREN
			}
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					select {
					case c <- Queen{column, r}:
					case <-shutdown:
						log.Println("Woulda been a leak")
						break
					}
				}
			}()
			__1_c = c
			goto __1_END_CHILDREN

		}
		////////////////////////////////////////////////////////
	__1_END_CHILDREN:
	}
	close(shutdown)

	log.Println(winners)
	time.Sleep(time.Second * 4)
}
