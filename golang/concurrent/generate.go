package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const template = `
// 'if' used to scope this entire engine.
if true {
	// User declaration.
	__{ID}_USER_children := func({CHILDREN_PARAM1} {USERTYPE}) chan {USERTYPE} {
		{CHILDREN_BODY}
	}
	// User declaration.
	__{ID}_USER_accept := func({ACCEPT_PARAM1} []{USERTYPE}) bool {
		{ACCEPT_BODY}
	}
	// User declaration.
	__{ID}_USER_reject := func({REJECT_PARAM1} {USERTYPE}, {REJECT_PARAM2} []{USERTYPE}) bool {
		{REJECT_BODY}
	}
	// Parent:Children PODO meant for stack management. 
	type __{ID}_StackEntry struct {
		Parent   {USERTYPE}
		Children chan {USERTYPE}
	}
	/////////////// Engine initialization.
	// Stack of Parent:Chidren pairs.
	__{ID}_stack := make([]__{ID}_StackEntry, 0)
	// Solution thus far.
	__{ID}_solution := make([]{USERTYPE}, 0)
	// Current root under consideration.
	__{ID}_root := {ROOT_EXPRESSION}
	// Current candidate under consideration.
	var __{ID}_candidate {USERTYPE}
	// Holds a Stack Entry and popping.
	var __{ID}_se __{ID}_StackEntry
	// Generic bool holder.
	var __{ID}_ok bool

	// Begin search.
	__{ID}_c := __{ID}_USER_children(__{ID}_root)
	for {
		if __{ID}_candidate, __{ID}_ok = <-__{ID}_c; !__{ID}_ok {
			if len(__{ID}_stack) == 0 {
				break
			}
			__{ID}_solution = __{ID}_solution[:len(__{ID}_solution)-1]
			__{ID}_se = __{ID}_stack[len(__{ID}_stack)-1]
			__{ID}_stack = __{ID}_stack[:len(__{ID}_stack)-1]
			__{ID}_root = __{ID}_se.Parent
			__{ID}_c = __{ID}_se.Children
			continue
		}
		__{ID}_reject := __{ID}_USER_reject(__{ID}_candidate, __{ID}_solution)
		if __{ID}_reject {
			continue
		}
		__{ID}_solution = append(__{ID}_solution, __{ID}_candidate)
		__{ID}_accept := __{ID}_USER_accept(__{ID}_solution)
		if __{ID}_accept {
			log.Println(__{ID}_solution)
			__{ID}_solution = __{ID}_solution[:len(__{ID}_solution)-1]
			continue
		}
		__{ID}_stack = append(__{ID}_stack, __{ID}_StackEntry{__{ID}_root, __{ID}_c})
		__{ID}_root = __{ID}_candidate
		__{ID}_c = __{ID}_USER_children(__{ID}_root)
	}
}`

var searchRegex = regexp.MustCompile(`\s*search\s+from\s+(?P<ROOT>.*)\s+\((?P<UTYPE>.*)\)\s*{\s*accept\s+(?P<ACCEPT_PARAM_SOLUTION>.*)\s*:\s*\n\s*(?P<ACCEPT_BODY>(?:.*\n)*)\s*reject\s+(?P<REJECT_PARAM_CANDIDATE>.*)\s*,\s*(?P<REJECT_PARAM_SOLUTION>.*)\s*:\s*\n(?P<REJECT_BODY>(?:.*\n)*)\s*children\s+(?P<CHILDREN_PARAM_PARENT>.*)\s*:\s*\n(?P<CHILDREN_BODY>(?:.*\n)*)\s*}`)
var searchSignature = regexp.MustCompile(`\s*search\s+from.*`)
var nameMap = func(r *regexp.Regexp) map[string]int {
	m := make(map[string]int, 0)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = i
		}
	}
	return m
}(searchRegex)

type Search struct {
	ID       string
	Root     string
	UserType string
	Accept   struct {
		ParamSolution string
		Body          string
	}
	Reject struct {
		ParamCandidate string
		ParamSolution  string
		Body           string
	}
	Children struct {
		ParamParent string
		Body        string
	}
}

func (s *Search) String() string {
	r := strings.NewReplacer("{ID}", s.ID,
		"{USERTYPE}", s.UserType,
		"{ROOT_EXPRESSION}", s.Root,
		"{CHILDREN_PARAM1}", s.Children.ParamParent,
		"{REJECT_PARAM1}", s.Reject.ParamCandidate,
		"{REJECT_PARAM2}", s.Reject.ParamSolution,
		"{REJECT_BODY}", s.Reject.Body,
		"{ACCEPT_PARAM1}", s.Accept.ParamSolution,
		"{ACCEPT_BODY}", s.Accept.Body,
		"{CHILDREN_BODY}", s.Children.Body)
	return fmt.Sprintln(r.Replace(template))
}

func isolate(l string, b *bufio.Reader) string {
	buf := bytes.NewBuffer([]byte{})
	o := bufio.NewWriter(buf)
	o.WriteString(l)
	count := 1
	for l, err := b.ReadString('\n'); err == nil; l, err = b.ReadString('\n') {
		for _, c := range l {
			if c == '{' {
				count++
			} else if c == '}' {
				count--
			}
		}
		o.WriteString(l)
		if count == 0 {
			break
		}

	}
	o.Flush()
	return string(buf.Bytes())
}

func build(l string, b *bufio.Reader) string {
	isolated := isolate(l, b)
	match := searchRegex.FindStringSubmatch(isolated)
	s := NewSearch(match)
	return s.String()
}

func RandomID() string {
	src := make([]byte, 4)
	rand.Read(src)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

func NewSearch(match []string) *Search {
	s := new(Search)
	s.ID = RandomID()
	for name, i := range nameMap {
		m := match[i]
		switch name {
		case "ROOT":
			s.Root = m
		case "UTYPE":
			s.UserType = m
		case "ACCEPT_PARAM_SOLUTION":
			s.Accept.ParamSolution = m
		case "ACCEPT_BODY":
			s.Accept.Body = m
		case "REJECT_PARAM_CANDIDATE":
			s.Reject.ParamCandidate = m
		case "REJECT_PARAM_SOLUTION":
			s.Reject.ParamSolution = m
		case "REJECT_BODY":
			s.Reject.Body = m
		case "CHILDREN_PARAM_PARENT":
			s.Children.ParamParent = m
		case "CHILDREN_BODY":
			s.Children.Body = m
		}
	}
	return s
}

func parseFile(f io.Reader, w io.Writer) {
	b := bufio.NewReader(f)
	o := bufio.NewWriter(w)
	var l string
	var err error
	for l, err = b.ReadString('\n'); err == nil; l, err = b.ReadString('\n') {
		if searchSignature.MatchString(l) {
			l = build(l, b)
		}
		fmt.Fprint(w, l)
	}
	fmt.Fprint(w, l)
	o.Flush()
}

var in string
var out string

func init() {
	flag.StringVar(&in, "in", "testin.go", "Path to the input Go file.")
	flag.StringVar(&out, "out", "testout.go", "Path to the output Go file.")
	flag.Parse()
}

func main() {
	log.Println(in, out)
	inf, err := os.Open(in)
	if err != nil {
		log.Panic(err)
	}
	defer inf.Close()
	outf, err := os.Create(out)
	if err != nil {
		log.Println(err)
	}
	defer outf.Close()
	parseFile(inf, outf)
	cmd := exec.Command("goimports", "-w", out)
	cmd.Run()
}