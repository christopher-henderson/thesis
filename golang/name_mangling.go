package main

import (
	"log"
	"time"
)

type USERTYPE struct {
	Column int
	Row    int
}

func backtrack() {
	// PRETEND USERLAND CLOSURE
	N := 4
	root := USERTYPE{0, 0}
	////

	// The rest is inlined code.

	type __ID_StackEntry struct {
		Parent   USERTYPE
		Children chan USERTYPE
	}

	// Engine initialization.
	__ID_stack := make([]__ID_StackEntry, 0)
	__ID_solution := make([]USERTYPE, 0)
	__ID_root := root

	var __ID_c chan USERTYPE
	////////////////////////////////////////////////////////
	// USERLAND
	for {
		// PARAMETER BINDINGS
		root := __ID_root
		/////////////////////
		column := root.Column + 1
		c := make(chan USERTYPE, 0)
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				c <- USERTYPE{column, r}
			}
		}()
		__ID_c = c
		goto INIT_CHILDREN
	}
	////////////////////////////////////////////////////////
INIT_CHILDREN:

	var __ID_candidate USERTYPE
	var __ID_ok bool
	var __ID_se __ID_StackEntry
	for {
		if __ID_candidate, __ID_ok = <-__ID_c; !__ID_ok {
			if len(__ID_stack) == 0 {
				return
			}
			__ID_solution = __ID_solution[:len(__ID_solution)-1]
			__ID_se = __ID_stack[len(__ID_stack)-1]
			__ID_stack = __ID_stack[:len(__ID_stack)-1]
			__ID_root = __ID_se.Parent
			__ID_c = __ID_se.Children
			continue
		}

		var __ID_reject bool
		////////////////////////////////////////////////////////
		// USERLAND - REJECT
		for {
			// PARAMETER BINDINGS
			candidate := __ID_candidate
			solution := __ID_solution
			/////////////////////
			row, column := candidate.Row, candidate.Column
			for _, q := range solution {
				r, c := q.Row, q.Column
				if row == r ||
					column == c ||
					row+column == r+c ||
					row-column == r-c {
					__ID_reject = true
					goto __ID_END_REJECT
				}
			}
			__ID_reject = false
			goto __ID_END_REJECT
		}
		////////////////////////////////////////////////////////
	__ID_END_REJECT:
		if __ID_reject {
			continue
		}
		__ID_solution = append(__ID_solution, __ID_candidate)

		var __ID_accept bool
		////////////////////////////////////////////////////////
		// USERLAND - ACCEPT
		for {
			// PARAMETER BINDINGS
			solution := __ID_solution
			/////////////////////

			__ID_accept = len(solution) == N
			goto __ID_END_ACCEPT
		}
		////////////////////////////////////////////////////////
	__ID_END_ACCEPT:
		if __ID_accept {
			log.Println(__ID_solution)
			__ID_solution = __ID_solution[:len(__ID_solution)-1] // Pop and throw final element away.
			continue
		}
		__ID_stack = append(__ID_stack, __ID_StackEntry{__ID_root, __ID_c})
		__ID_root = __ID_candidate

		////////////////////////////////////////////////////////
		// USERLAND - CHILDREN
		for {
			// PARAMETER BINDINGS
			root := __ID_root
			/////////////////////

			column := root.Column + 1
			c := make(chan USERTYPE, 0)
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					c <- USERTYPE{column, r}
				}
			}()
			__ID_c = c
			goto __ID_END_NEXT_CHILD
		}
		////////////////////////////////////////////////////////
	__ID_END_NEXT_CHILD:
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	t := time.Now()
	backtrack()
	log.Println(time.Now().Sub(t))
}
