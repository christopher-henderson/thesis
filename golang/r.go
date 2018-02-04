package main

import (
	"log"
	"regexp"
)

const lol = ` lol1    lol2   `

func main() {
	r := regexp.MustCompile(`(?P<DERP>lol[0-9]\s*)`)
	m := r.FindAllStringSubmatch(lol, -1)
	log.Println(m)
	i := r.FindAllStringIndex(lol, -1)
	log.Println(i)
}
