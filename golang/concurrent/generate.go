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
if true {
// This is the one piece of internals that the userland
	// can potentially see.
	type __GraphNode struct {
		Active	 bool
		ID       int
		Parent   int
	}
	// User CHILDREN declaration.
	USER_children := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) <-chan {UTYPE} {
		{CHILDREN_BODY}
	}
	// User ACCEPT declaration.
	USER_accept := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) bool {
		{ACCEPT_BODY}
	}
	// User REJECT declaration.
	USER_reject := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) bool {
		{REJECT_BODY}
	}
	root := {ROOT_EXPRESSION}
	maxgoroutine := {MAX_GOROUTINE}
	// Parent:Children PODO meant for stack management.
	type StackEntry struct {
		Parent   {UTYPE}
		Children <-chan {UTYPE}
	}
	lock := make(chan int, maxgoroutine)
	wg := make(chan int, maxgoroutine)
	ticket := make(chan int, maxgoroutine)
	// You have to declare first since the function can fire off a
	// goroutine of itself.
	var engine func(solution []{UTYPE}, root {UTYPE}, gid *__GraphNode)
	engine = func(solution []{UTYPE}, root {UTYPE}, gid *__GraphNode) {
		_children := USER_children(root, solution, gid)
		// Stack of Parent:Chidren pairs.
		stack := make([]StackEntry, 0)
		// Current candidate under consideration.
		var candidate {UTYPE}
		// Holds a StackEntry.
		var stackEntry StackEntry
		// Generic boolean variable
		var ok bool
		for {
			if candidate, ok = <-_children; !ok {
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
				_children = stackEntry.Children
				continue
			}
			// Ask the user if we should reject this candidate.
			_reject := USER_reject(candidate, solution, gid)
			if _reject {
				// Rejected candidate.
				continue
			}
			// Append the candidate to the solution.
			solution = append(solution, candidate)
			// Ask the user if we should accept this solution.
			_accept := USER_accept(candidate, solution, gid)
			if _accept {
				// Accepted solution.
				// Pop from the solution thus far and continue on with the next child.
				solution = solution[:len(solution)-1]
				continue
			}
			select {
				case lock <- 1:
					wg <- 1
					s := make([]{UTYPE}, len(solution))
					copy(s, solution)
					go engine(s, candidate, &__GraphNode{Active: true, ID: <-ticket, Parent: gid.ID})
					// pretend we didn't see this
					solution = solution[:len(solution)-1]
					continue
				default:
			}
			// Push the current root to the stack.
			stack = append(stack, StackEntry{root, _children})
			// Make the candidate the new root.
			root = candidate
			// Get the new root's children channel.
			_children = USER_children(root, solution, gid)
		}
		<- lock
		wg <- -1
		gid.Active = false
	}
	shutdown := make(chan struct{}, 0)
	go func() {
		// Goroutine ticketing system.
		id := 0
		for {
			select {
			case ticket <- id:
				id++
			case <-shutdown:
				close(ticket)
				return
			}
		}
	}()
	lock <- 1
	wg <- 1
	go engine(make([]{UTYPE}, 0), root, &__GraphNode{Active: true, ID: <-ticket, Parent: 0})
	count := 0
	for c := range wg {
		count += c
		if count == 0 {
			break
		}
	}
	close(shutdown)
	close(wg)
	close(lock)
}
`

const DFS = `// 'if' used to scope this entire engine.
if true {
	// User declaration.
	__{ID}_USER_children := func({CHILDREN_PARAM1} {USERTYPE}) chan {USERTYPE} {
		{CHILDREN_BODY}
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
	// __{ID}_solution := make([]{USERTYPE}, 0)
	// Current root under consideration.
	__{ID}_root := {ROOT_EXPRESSION}
	// Current candidate under consideration.
	var __{ID}_candidate {USERTYPE}
	// Holds a Stack Entry and popping.
	var __{ID}_stackEntry __{ID}_StackEntry
	// Generic bool holder.
	var __{ID}_ok bool

	// Begin search.
	__{ID}_children := __{ID}_USER_children(__{ID}_root)
	for {
		if __{ID}_candidate, __{ID}_ok = <-__{ID}_children; !__{ID}_ok {
			if len(__{ID}_stack) == 0 {
				break
			}
			// __{ID}_solution = __{ID}_solution[:len(__{ID}_solution)-1]
			__{ID}_stackEntry = __{ID}_stack[len(__{ID}_stack)-1]
			__{ID}_stack = __{ID}_stack[:len(__{ID}_stack)-1]
			__{ID}_root = __{ID}_stackEntry.Parent
			__{ID}_children = __{ID}_stackEntry.Children
			continue
		}
		// __{ID}_solution = append(__{ID}_solution, __{ID}_candidate)
		__{ID}_stack = append(__{ID}_stack, __{ID}_StackEntry{__{ID}_root, __{ID}_children})
		__{ID}_root = __{ID}_candidate
		__{ID}_children = __{ID}_USER_children(__{ID}_root)
	}
}`

// s	let . match \n (default false)
// U	ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)
var regexFlags = `(?sU)`
var signature = `\s*search\s+from\s+(?P<ROOT>.*)\s+\((?P<UTYPE>.*)\)\s*{`
var concurrent = `(?P<CONCURRENT>\s*concurrent\s*)?`
var children = `\s*children\s+(?P<CHILDREN_PARAM_PARENT>.*)\s*:\s*(?P<CHILDREN_BODY>(?:.*)*)`
var accept = `\s*(?:accept\s+(?P<ACCEPT_PARAM_SOLUTION>.*)\s*:\s*\s*(?P<ACCEPT_BODY>(?:.*)*))?`
var reject = `\s*(?:reject\s+(?P<REJECT_PARAM_CANDIDATE>.*)\s*,\s*(?P<REJECT_PARAM_SOLUTION>.*)\s*:\s*(?P<REJECT_BODY>(?:.*)*))?`
var add = `\s*(?:add\s+(?P<ADD_PARAM_CANDIDATE>.*),\s+(?P<ADD_PARAM_SOLUTION>.*)\s*:\s*(?P<ADD_BODY>(?:.*)*))?`
var remove = `\s*(?:remove\s+(?P<REMOVE_PARAM_CANDIDATE>.*),\s+(?P<REMOVE_PARAM_SOLUTION>.*)\s*:\s*(?P<REMOVE_BODY>(?:.*)*))?`
var end = `\s*}$`

var completeText = fmt.Sprintf("%v%v%v%v%v%v%v%v%v", regexFlags, signature, concurrent, children, accept, reject, add, remove, end)

var searchRegex = regexp.MustCompile(completeText)

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
	ID           string
	Root         string
	UserType     string
	MaxGoroutine string
	Accept       struct {
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
	Add struct {
		ParamCandidate string
		ParamSolution  string
		Body           string
	}
	Remove struct {
		ParamCandidate string
		ParamSolution  string
		Body           string
	}
}

func (s *Search) String() string {
	r := strings.NewReplacer("{ID}", s.ID,
		"{UTYPE}", s.UserType,
		"{ROOT_EXPRESSION}", s.Root,
		"{CHILDREN_PARAM1}", s.Children.ParamParent,
		"{REJECT_PARAM1}", s.Reject.ParamCandidate,
		"{REJECT_PARAM2}", s.Reject.ParamSolution,
		"{REJECT_BODY}", s.Reject.Body,
		"{ACCEPT_PARAM1}", s.Accept.ParamSolution,
		"{ACCEPT_BODY}", s.Accept.Body,
		"{CHILDREN_BODY}", s.Children.Body,
		"{MAX_GOROUTINE}", s.MaxGoroutine)
	if s.Accept.Body == "" && s.Reject.Body == "" {
		return fmt.Sprintln(r.Replace(DFS))
	}
	return fmt.Sprintln(r.Replace(template))
}

var leftTrim = regexp.MustCompile(`(?s)^\s+(.*)`)
var rightTrim = regexp.MustCompile(`(?s)(.*})\s+$`)

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
	return rightTrim.ReplaceAllString(string(buf.Bytes()), "$1")
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
			s.Accept.Body = leftTrim.ReplaceAllString(m, "$1")
		case "REJECT_PARAM_CANDIDATE":
			s.Reject.ParamCandidate = m
		case "REJECT_PARAM_SOLUTION":
			s.Reject.ParamSolution = m
		case "REJECT_BODY":
			s.Reject.Body = leftTrim.ReplaceAllString(m, "$1")
		case "CHILDREN_PARAM_PARENT":
			s.Children.ParamParent = m
		case "CHILDREN_BODY":
			s.Children.Body = leftTrim.ReplaceAllString(m, "$1")

		case "ADD_PARAM_CANDIDATE":
			s.Add.ParamCandidate = m
		case "ADD_PARAM_SOLUTION":
			s.Add.ParamSolution = m
		case "ADD_BODY":
			s.Add.Body = leftTrim.ReplaceAllString(m, "$1")

		case "REMOVE_PARAM_CANDIDATE":
			s.Remove.ParamCandidate = m
		case "REMOVE_PARAM_SOLUTION":
			s.Remove.ParamSolution = m
		case "REMOVE_BODY":
			s.Remove.Body = leftTrim.ReplaceAllString(m, "$1")

		case "CONCURRENT":
			if m == "" {
				s.MaxGoroutine = "1"
			} else {
				s.MaxGoroutine = "runtime.NumCPU()"
			}
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
