package main

import "fmt"

func main() {
	search 0; int {
	children:
		c := make(chan int, 0)
		go func() {
			for i := node + 1; i < 3; i++ {
				c <- i
			}
			close(c)
		}()
		return c
	accept:
		if node == 2 {
			search 0; int {
			children:
				fmt.Println("LOL SUB SEARCH")
				c := make(chan int, 0)
				close(c)
				return c
			}
			return true
		}
		return false
	}
}
