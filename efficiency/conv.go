package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Size    int64
	Cores   int64
	Time    float64
	Mallocs int64
}

func parseCSV(r io.Reader) []Entry {
	b := bufio.NewReader(r)
	entries := make([]Entry, 0)
	for line, err := b.ReadString('\n'); err == nil; line, err = b.ReadString('\n') {
		e := strings.Split(strings.Trim(line, "\n"), ",")
		size, err := strconv.ParseInt(strings.Trim(e[0], " "), 10, 32)
		if err != nil {
			log.Panicln(err)
		}
		cores, err := strconv.ParseInt(strings.Trim(e[1], " "), 10, 32)
		if err != nil {
			log.Panicln(err)
		}
		time, err := time.ParseDuration(strings.Trim(e[2], " "))
		if err != nil {
			log.Panicln(err)
		}
		mallocs, err := strconv.ParseInt(strings.Trim(e[3], " "), 10, 32)
		if err != nil {
			log.Panicln(err)
		}
		entries = append(entries, Entry{size, cores, float64(time) / 1000000000, mallocs})
	}
	return entries
}

func main() {
	f, err := os.Open("timings.csv")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()
	csv := parseCSV(f)
	o, err := os.Create("new.csv")
	if err != nil {
		log.Panic(err)
	}
	defer o.Close()
	for _, c := range csv {
		o.WriteString(fmt.Sprintf("%v,%v,%f,%v\n", c.Size, c.Cores, c.Time, c.Mallocs))
	}
}
