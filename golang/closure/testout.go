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
			__e4931531_USER_children := func(Queen) chan Queen {

			}
			// User declaration.
			__e4931531_USER_accept := func([]Queen) bool {

			}
			// User declaration.
			__e4931531_USER_reject := func(Queen, []Queen) bool {

			}
			// Parent:Children PODO meant for stack management.
			type __e4931531_StackEntry struct {
				Parent   Queen
				Children chan Queen
			}
			/////////////// Engine initialization.
			// Stack of Parent:Chidren pairs.
			__e4931531_stack := make([]__e4931531_StackEntry, 0)
			// Solution thus far.
			__e4931531_solution := make([]Queen, 0)
			// Current root under consideration.
			__e4931531_root := Queen{0, 0}
			// Current candidate under consideration.
			var __e4931531_candidate Queen
			// Holds a Stack Entry and popping.
			var __e4931531_se __e4931531_StackEntry
			// Generic bool holder.
			var __e4931531_ok bool

			// Begin search.
			__e4931531_c := __e4931531_USER_children(__e4931531_root)
			for {
				if __e4931531_candidate, __e4931531_ok = <-__e4931531_c; !__e4931531_ok {
					if len(__e4931531_stack) == 0 {
						break
					}
					__e4931531_solution = __e4931531_solution[:len(__e4931531_solution)-1]
					__e4931531_se = __e4931531_stack[len(__e4931531_stack)-1]
					__e4931531_stack = __e4931531_stack[:len(__e4931531_stack)-1]
					__e4931531_root = __e4931531_se.Parent
					__e4931531_c = __e4931531_se.Children
					continue
				}
				__e4931531_reject := __e4931531_USER_reject(__e4931531_candidate, __e4931531_solution)
				if __e4931531_reject {
					continue
				}
				__e4931531_solution = append(__e4931531_solution, __e4931531_candidate)
				__e4931531_accept := __e4931531_USER_accept(__e4931531_solution)
				if __e4931531_accept {
					log.Println(__e4931531_solution)
					__e4931531_solution = __e4931531_solution[:len(__e4931531_solution)-1]
					continue
				}
				__e4931531_stack = append(__e4931531_stack, __e4931531_StackEntry{__e4931531_root, __e4931531_c})
				__e4931531_root = __e4931531_candidate
				__e4931531_c = __e4931531_USER_children(__e4931531_root)
			}
		}
		log.Println(winners)
		log.Println(time.Now().Sub(start))
	}
}
