package main

import (
	"runtime"
	"sync"
)

type Queen struct {
	Column int
	Row    int
}

func NQueens(N int) {
	// start := time.Now()
	// winners := 0
	// l := sync.Mutex{}
	// 'if' used to scope this entire engine.
	if true {
		// User CHILDREN declaration.
		USER_children := func(parent Queen) chan Queen {
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
		USER_accept := func(solution []Queen) bool {
			if len(solution) == N {
				// Print it, append it to a slice of solutions,
				// whatever you want here.
				// log.Println(solution)
				// Keeping track of number of solutions
				// for quick verification.
				// l.Lock()
				// winners++
				// l.Unlock()
				return true
			}
			return false
		}
		// User REJECT declaration.
		USER_reject := func(candidate Queen, solution []Queen) bool {
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
		type StackEntry struct {
			Parent   Queen
			Children chan Queen
		}
		lock := struct {
			max   int
			count int
			sync.WaitGroup
			sync.Mutex
		}{}
		lock.max = runtime.NumCPU()
		// capture := func() bool {
		// 	lock.Mutex.Lock()
		// 	defer lock.Mutex.Unlock()
		// 	if lock.count >= lock.max {
		// 		return false
		// 	}
		// 	lock.count++
		// 	lock.WaitGroup.Add(1)
		// 	return true
		// }
		// done := func() {
		// 	lock.Mutex.Lock()
		// 	lock.WaitGroup.Done()
		// 	lock.count -= 1
		// 	lock.Mutex.Unlock()
		// }
		// You have to declare first since the function can fire off a
		// goroutine of itself.
		var engine func(solution []Queen, root Queen, children chan Queen)
		engine = func(solution []Queen, root Queen, children chan Queen) {
			// Stack of Parent:Chidren pairs.
			stack := make([]StackEntry, 0)
			// Holds a StackEntry.
			var stackEntry StackEntry
			candidate, ok := <-children
			for {
				if !ok {
					// This node has no further children.
					if len(stack) == 0 {
						// Algorithm termination. No further nodes in the stack.
						break
					}
					// With no valid children left, we pop the latest node from the solution.
					solution = solution[:len(solution)-1]
					// Pop from the stack. Broken into two steps:
					// 	1. Get final element.
					//	2. Resize the stack.
					stackEntry = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					// Extract root and candidate fields from the StackEntry.
					root = stackEntry.Parent
					children = stackEntry.Children
					candidate, ok = <-children
					continue
				}
				// Ask the user if we should reject this candidate.
				reject := USER_reject(candidate, solution)
				if reject {
					candidate, ok = <-children
					// Rejected candidate.
					continue
				}
				// Append the candidate to the solution.
				solution = append(solution, candidate)
				// Ask the user if we should accept this solution.
				accept := USER_accept(solution)
				if accept {
					// Accepted solution.
					// Pop from the solution thus far and continue on with the next child.
					solution = solution[:len(solution)-1]
					candidate, ok = <-children
					continue
				}

				lock.Lock()
				if lock.count < lock.max {
					nextChild, ok := <-children
					if ok {
						lock.count++
						lock.Add(1)
						lock.Unlock()
						dst := make([]Queen, len(solution))
						copy(dst, solution)
						go engine(dst, candidate, USER_children(candidate))
						candidate = nextChild
						// pretend we didn't see this
						solution = solution[:len(solution)-1]
						continue
					}
				}
				lock.Unlock()
				// Push the current root to the stack.
				stack = append(stack, StackEntry{root, children})
				// Make the candidate the new root.
				root = candidate
				// Get the new root's children channel.
				children = USER_children(root)
				candidate, ok = <-children

				// if capture() {
				// 	dst := make([]Queen, len(solution))
				// 	copy(dst, solution)
				// 	go engine(dst, candidate, USER_children(candidate))
				// 	// pretend we didn't see this
				// 	solution = solution[:len(solution)-1]
				// 	continue
				// }

				// // Push the current root to the stack.
				// stack = append(stack, StackEntry{root, children})
				// // Make the candidate the new root.
				// root = candidate
				// // Get the new root's children channel.
				// children = USER_children(root)
			}
			lock.Lock()
			lock.count -= 1
			lock.Done()
			lock.Mutex.Unlock()
		}
		root := Queen{0, 0}
		lock.count++
		lock.Add(1)
		go engine(make([]Queen, 0), root, USER_children(root))
		lock.WaitGroup.Wait()
	}
	// log.Println(winners)
	// log.Println(time.Now().Sub(start))

}

func main() {
	NQueens(6)
}
