package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var searchRegex = regexp.MustCompile(`\s*search\s+from\s+(?P<ROOT>.*)\s+\((?P<UTYPE>.*)\)\s*{\s*accept\s+(?P<ACCEPT_PARAM_SOLUTION>.*)\s*:\s*\n\s*(?P<ACCEPT_BODY>(?:.*\n)*)\s*reject\s+(?P<REJECT_PARAM_CANDIDATE>.*)\s*,\s*(?P<REJECT_PARAM_SOLUTION>.*)\s*:\s*\n(?P<REJECT_BODY>(?:.*\n)*)\s*children\s+(?P<CHILDREN_PARAM_PARENT>.*)\s*:\s*\n(?P<CHILDREN_BODY>(?:.*\n)*)\s*}`)
var nameMap = make_map(searchRegex)

const example = `
package main

type Queen struct {
	Column int
	Row    int
}

func main() {
	N := 8
	winners := 0
	search from Queen{0,0} (Queen) {
		accept solution:
			if len(solution) == N {
				winners++
			}
			return len(solution) == N
		reject candidate, solution:
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
		children parent:
			column := parent.Column + 1
			c := make(chan Queen, 0)
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
	log.Println(winners)
}
`

type Search struct {
	ID       int
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
	ChildrenInit struct {
		ParamParent string
		Body        string
	}
}

func (s *Search) String() string {
	r := strings.NewReplacer("{ID}", strconv.FormatInt(int64(s.ID), 10),
		"{USERTYPE}", s.UserType,
		"{ROOT_EXPRESSION}", s.Root,
		"{CHILDREN_PARAM1}", s.Children.ParamParent,
		"{CHILDREN_INIT_BODY}", s.ChildrenInit.Body,
		"{REJECT_PARAM1}", s.Reject.ParamCandidate,
		"{REJECT_PARAM2}", s.Reject.ParamSolution,
		"{REJECT_BODY}", s.Reject.Body,
		"{ACCEPT_PARAM1}", s.Accept.ParamSolution,
		"{ACCEPT_BODY}", s.Accept.Body,
		"{CHILDREN_BODY}", s.Children.Body)
	return r.Replace(template)
}

const template = `
type __{ID}_StackEntry struct {
	Parent   {USERTYPE}
	Children chan {USERTYPE}
}

// Engine initialization.
__{ID}_stack := make([]__{ID}_StackEntry, 0)
__{ID}_solution := make([]{USERTYPE}, 0)
__{ID}_root := {ROOT_EXPRESSION}

var __{ID}_c chan {USERTYPE}
////////////////////////////////////////////////////////
// USERLAND
for {
	// PARAMETER BINDINGS
	{CHILDREN_PARAM1} := __{ID}_root
	/////////////////////
	{CHILDREN_INIT_BODY}
}
////////////////////////////////////////////////////////
__{ID}_END_INIT_CHILDREN:

var __{ID}_candidate {USERTYPE}
var __{ID}_ok bool
var __{ID}_se __{ID}_StackEntry
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
	var __{ID}_reject bool
	////////////////////////////////////////////////////////
	// USERLAND - REJECT
	for {
		// PARAMETER BINDINGS
		{REJECT_PARAM1} := __{ID}_candidate
		{REJECT_PARAM2} := __{ID}_solution
		/////////////////////
		{REJECT_BODY}
	}
	////////////////////////////////////////////////////////
__{ID}_END_REJECT:
	if __{ID}_reject {
		continue
	}
	__{ID}_solution = append(__{ID}_solution, __{ID}_candidate)
	var __{ID}_accept bool
	////////////////////////////////////////////////////////
	// USERLAND - ACCEPT
	for {
		// PARAMETER BINDINGS
		{ACCEPT_PARAM1} := __{ID}_solution
		/////////////////////
		{ACCEPT_BODY}
	}
	////////////////////////////////////////////////////////
__{ID}_END_ACCEPT:
	if __{ID}_accept {
		log.Println(__{ID}_solution)
		__{ID}_solution = __{ID}_solution[:len(__{ID}_solution)-1]
		continue
	}
	__{ID}_stack = append(__{ID}_stack, __{ID}_StackEntry{__{ID}_root, __{ID}_c})
	__{ID}_root = __{ID}_candidate
	////////////////////////////////////////////////////////
	// USERLAND - CHILDREN
	for {
		// PARAMETER BINDINGS
		{CHILDREN_PARAM1} := __{ID}_root
		/////////////////////
		{CHILDREN_BODY}
	}
	////////////////////////////////////////////////////////
__{ID}_END_CHILDREN:
}`

func make_map(r *regexp.Regexp) map[string]int {
	m := make(map[string]int, 0)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = i
		}
	}
	return m
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

var id = 0

func NewSearch(match []string) *Search {
	id++
	s := new(Search)
	s.ID = id
	acceptVar := fmt.Sprintf("__%v_accept", s.ID)
	acceptLabel := fmt.Sprintf("__%v_END_ACCEPT", s.ID)
	rejectVar := fmt.Sprintf("__%v_reject", s.ID)
	rejectLabel := fmt.Sprintf("__%v_END_REJECT", s.ID)
	childrenVar := fmt.Sprintf("__%v_c", s.ID)
	childrenLabel := fmt.Sprintf("__%v_END_CHILDREN", s.ID)
	childrenInitLabel := fmt.Sprintf("__%v_END_INIT_CHILDREN", s.ID)
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
			s.Accept.Body = inlineFunction(m, acceptVar, acceptLabel)
		case "REJECT_PARAM_CANDIDATE":
			s.Reject.ParamCandidate = m
		case "REJECT_PARAM_SOLUTION":
			s.Reject.ParamSolution = m
		case "REJECT_BODY":
			s.Reject.Body = inlineFunction(m, rejectVar, rejectLabel)
		case "CHILDREN_PARAM_PARENT":
			s.Children.ParamParent = m
		case "CHILDREN_BODY":
			s.Children.Body = inlineFunction(m, childrenVar, childrenLabel)
			s.ChildrenInit.Body = inlineFunction(m, childrenVar, childrenInitLabel)
		}
	}
	return s
}

var returnRegex = regexp.MustCompile(`\s*return`)

func inlineFunction(body string, ret string, label string) string {
	// r := bufio.NewReader(strings.NewReader(body))
	r := bufio.NewScanner(strings.NewReader(body))
	buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(buf)
	var err error
	var l string
	for r.Scan() {
		l = r.Text()
		// for l, err = r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		if !returnRegex.Match([]byte(l)) {
			fmt.Fprintln(w, l)
			// w.WriteString(l)
			continue
		}
		l = returnRegex.ReplaceAllString(l, fmt.Sprintf("%v =", ret))
		fmt.Fprintln(w, l)
		curly := strings.Count(l, "{") - strings.Count(l, "}")
		bracket := strings.Count(l, "[") - strings.Count(l, "]")
		parens := strings.Count(l, "(") - strings.Count(l, ")")
		for curly != 0 && bracket != 0 && parens != 0 {
			if !r.Scan() {
				log.Panic(r.Err())
			}
			l := r.Text()
			if err != nil {
				log.Panic(err)
			}
			fmt.Fprintln(w, l)
			for _, c := range l {
				switch c {
				case '{':
					curly++
				case '}':
					curly--
				case '[':
					bracket++
				case ']':
					bracket--
				case '(':
					parens++
				case ')':
					parens--
				}
			}
		}
		fmt.Fprintln(w, fmt.Sprintf("goto %v", label))
	}
	w.Flush()
	s := string(buf.Bytes())
	if !strings.HasSuffix(s, fmt.Sprintf("goto %v", label)) {
		fmt.Fprintln(w, fmt.Sprintf("goto %v", label))
	}
	return string(buf.Bytes())
}

func parseFile(f io.Reader, w io.Writer) {
	trim := regexp.MustCompile(`\s*(.*)`)
	b := bufio.NewReader(f)
	o := bufio.NewWriter(w)
	for l, err := b.ReadString('\n'); err == nil; l, err = b.ReadString('\n') {
		trimmed := trim.ReplaceAllString(l, "$1")
		if strings.HasPrefix(trimmed, "search") {
			l = build(l, b)
		}
		o.WriteString(l)
	}
	o.Flush()
	// log.Println(string(buf.Bytes()))
}

func main() {
	f, err := os.Create("pleaseGod.go")
	if err != nil {
		log.Println(err)
	}
	parseFile(strings.NewReader(example), f)
	cmd := exec.Command("goimports", "-w", "pleaseGod.go")
	cmd.Run()
}
