package main

import "fmt"

func main() {
	a := make([]int, 0)
	c := cap(a)
	for i := 0; i < 100; i++ {
		a = append(a, i)
		tmp := cap(a)
		if tmp != c {
			fmt.Println(tmp)
			c = tmp
		}
	}
}
