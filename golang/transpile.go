package main

import (
	"log"
	"time"
)

var N = 14

type Queen struct {
	Column int
	Row    int
}

// func reject(candidate Queen, solution []Queen) bool {
// 	row, column := candidate.Row, candidate.Column
// 	for _, q := range solution {
// 		r, c := q.Row, q.Column
// 		if row == r ||
// 			column == c ||
// 			row+column == r+c ||
// 			row-column == r-c {
// 			return true
// 		}
// 	}
// 	return false
// }

// func children(parent Queen) chan Queen {
// 	column := parent.Column + 1
// 	c := make(chan Queen, 0)
// 	go func() {
// 		defer close(c)
// 		for r := 1; r < N+1; r++ {
// 			c <- Queen{column, r}
// 		}
// 	}()
// 	return c
// }

// func accept(solution []Queen) bool {
// 	return len(solution) == N
// }

type RootStackEntry struct {
	Parent   Queen
	Children chan Queen
}

type RootStack []RootStackEntry

func rpush(s RootStack, e RootStackEntry) RootStack {
	return append(s, e)
}

func rpop(s RootStack) (RootStackEntry, RootStack) {
	if len(s) == 0 {
		return RootStackEntry{}, s
	}
	return s[len(s)-1], s[:len(s)-1]
}

type SolutionStack []Queen

func spush(s SolutionStack, e Queen) SolutionStack {
	return append(s, e)
}

func spop(s SolutionStack) (Queen, SolutionStack) {
	if len(s) == 0 {
		return Queen{}, s
	}
	return s[len(s)-1], s[:len(s)-1]
}

func backtrack(root Queen) {
	stack := RootStack{}
	solution := SolutionStack{}

	var __c chan Queen
	for {
		////////////////////////////////////////////////////////
		column := root.Column + 1
		c := make(chan Queen, 0)
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				c <- Queen{column, r}
			}
		}()
		__c = c
		goto INIT_CHILDREN
		////////////////////////////////////////////////////////
	}
INIT_CHILDREN:

	var candidate Queen
	var ok bool
	var rse RootStackEntry
	for {
		if candidate, ok = <-__c; !ok {
			_, solution = spop(solution)
			if len(stack) == 0 {
				log.Println(winners)
				return
			}
			rse, stack = rpop(stack)
			root = rse.Parent
			__c = rse.Children
			continue
		}

		var __reject bool
		for {
			////////////////////////////////////////////////////////
			row, column := candidate.Row, candidate.Column
			for _, q := range solution {
				r, c := q.Row, q.Column
				if row == r ||
					column == c ||
					row+column == r+c ||
					row-column == r-c {
					__reject = true
					goto END_REJECT
				}
			}
			__reject = false
			goto END_REJECT
			////////////////////////////////////////////////////////
		}
	END_REJECT:
		if __reject {
			continue
		}
		solution = spush(solution, candidate)

		var __accept bool
		for {
			////////////////////////////////////////////////////////
			__accept = len(solution) == N
			goto END_ACCEPT
			////////////////////////////////////////////////////////
		}
	END_ACCEPT:
		if __accept {
			log.Println(solution)
			_, solution = spop(solution)
			continue
		}
		stack = rpush(stack, RootStackEntry{root, __c})
		root = candidate

		for {
			////////////////////////////////////////////////////////
			column := root.Column + 1
			c := make(chan Queen, 0)
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					c <- Queen{column, r}
				}
			}()
			__c = c
			goto END_NEXT_CHILD
			////////////////////////////////////////////////////////
		}
	END_NEXT_CHILD:
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	t := time.Now()
	backtrack(Queen{0, 0})
	log.Println(time.Now().Sub(t))
}
