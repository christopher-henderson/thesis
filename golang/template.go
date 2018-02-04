package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var searchRegex = regexp.MustCompile(`\s*search\s+from\s+(?P<ROOT>.*)\s+\((?P<UTYPE>.*)\)\s*{\s*accept\s+(?P<ACCEPT_PARAM_SOLUTION>.*)\s*:\s*\n\s*(?P<ACCEPT_BODY>(?:.*\n)*)\s*reject\s+(?P<REJECT_PARAM_CANDIDATE>.*)\s*,\s*(?P<REJECT_PARAM_SOLUTION>.*)\s*:\s*\n(?P<REJECT_BODY>(?:.*\n)*)\s*children\s+(?P<CHILDREN_PARAM_PARENT>.*)\s*:\s*\n(?P<CHILDREN_BODY>(?:.*\n)*)\s*}`)
var nameMap = make_map(searchRegex)

const example = `
package main

type Queen struct {
	Row    int
	Column int
}

func main() {
	N := 8
	search from Queen{0,0} (Queen) {
		accept solution:
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
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					c <- Queen{column, r}
				}
			}()
			return c
	}


	search from Queen{0,0} (Queen) {
		accept solution:
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
			go func() {
				defer close(c)
				for r := 1; r < N+1; r++ {
					c <- Queen{column, r}
				}
			}()
			return c
	}
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
}

func (s *Search) String() {

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
	goto INIT_CHILDREN
}
////////////////////////////////////////////////////////
INIT_CHILDREN:

var __{ID}_candidate {USERTYPE}
var __{ID}_ok bool
var __{ID}_se __{ID}_StackEntry
for {
	if __{ID}_candidate, __{ID}_ok = <-__{ID}_c; !__{ID}_ok {
		if len(__{ID}_stack) == 0 {
			return
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
		goto __{ID}_END_REJECT
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
		goto __{ID}_END_ACCEPT
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
		{CHILDREN_INIT_BODY}
		goto __{ID}_END_NEXT_CHILD
	}
	////////////////////////////////////////////////////////
__{ID}_END_NEXT_CHILD:
}`

// var signature = regexp.MustCompile(`\s*search\s+from\s+(?P<ROOT>.*)\s+\((?P<USER_TYPE>.*)\)\s+{\s*\n`)
// var reject = regexp.MustCompile(`.*reject\s+(?P<REJECT_PARAM_CANDIDATE>.*)\s*,\s*(?P<REJECT_PARAM_SOLUTION>.*):\n`) // ((.*\n?)*)

// var signature_map = make_map(signature)
// var reject_map = make_map(reject)

func make_map(r *regexp.Regexp) map[string]int {
	m := make(map[string]int, 0)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = i
		}
	}
	return m
}

// func search_and_replace(text io.Reader, out io.Writer) {
// 	in := bufio.NewReader(text)
// 	id := 0
// 	for l, err := in.ReadString('\n'); err == nil; l, err = in.ReadString('\n') {
// 		if signature.Match([]byte(l)) {
// 			l = parse_search(l, id, in, out)
// 		}
// 		if _, err := out.Write([]byte(l)); err != nil {
// 			log.Panic(err)
// 		}

// 	}
// }

// func parse_search(l string, id int, in *bufio.Reader, out io.Writer) string {
// 	match := signature.FindStringSubmatch(l)
// 	root := match[signature_map["ROOT"]]
// 	utype := match[signature_map["USER_TYPE"]]

// 	var err error
// 	open := 1
// 	for l, err = in.ReadString('\n'); open != 0 && err == nil; l, err = in.ReadString('\n') {
// 		for _, c := range l {
// 			if c == '{' {
// 				open += 1
// 			} else if c == '}' {
// 				open -= 1
// 			}
// 		}
// 	}
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	search := new(Search)
// 	search.ID = id
// 	search.Root = root
// 	search.UserType = utype

// 	var left string
// 	for l, err := in.ReadString('\n'); err == nil; l, err = in.ReadString('\n') {
// 		switch {
// 		case reject.Match([]byte(l)):
// 			search.Reject, left = parse_reject(l, in)
// 		}
// 	}
// 	return left
// }

// func parse_reject(l string, in *bufio.Reader) {

// }

// func build(match []string, m map[string]int) *Search {
// 	s := new(Search)
// 	for name, _ := range m {
// 		switch name {
// 		case "SEARCH":
// 		case "ROOT":
// 		case "UTYPE":
// 		case "ACCEPT_PARAM_SOLUTION":
// 		case "ACCEPT_BLOCK":
// 		case "REJECT_PARAM_CANDIDATE":
// 		case "REJECT_PARAM_SOLUTION":
// 		case "REJECT_BLOCK":
// 		case "CHILDREN_PARAM_PAREN":
// 		case "CHILDREN_BLOCK":
// 		}
// 	}
// 	return s
// }

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
	log.Println(s)
	return isolated
}

var id = 0

func NewSearch(match []string) *Search {
	id++
	s := new(Search)
	s.ID = id
	for name, i := range nameMap {
		m := match[i]
		switch name {
		case "ROOT":
			s.Root = m
		case "UTYPE":
			s.UserType = m
		case "ACCEPT_PARAM_SOLUTION":
			s.Accept.ParamSolution = m
		case "ACCEPT_Body":
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

// var returnRegex = regexp.MustCompile(`(?P<LEAD>\s*)return(?P<EXP>.*)`)
var returnRegex = regexp.MustCompile(`\s*return`)

// var returnMap = make_map(returnRegex)

func inlineFunction(body string, ret string, label string) string {
	r := bufio.NewReader(strings.NewReader(body))
	buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(buf)
	var err error
	var l string
	for l, err = r.ReadString('\n'); err == nil; l, err = r.ReadString('\n') {
		log.Println(l)
		if !returnRegex.Match([]byte(l)) {
			w.WriteString(l)
			continue
		}
		// retMatches := returnRegex.FindStringSubmatch(l)
		// l = fmt.Sprintf("%v%v = %v\n", retMatches[returnMap["LEAD"]], ret, retMatches[returnMap["EXP"]])
		l = returnRegex.ReplaceAllString(l, fmt.Sprintf("%v =", ret))
		w.WriteString(l)
		curly := strings.Count(l, "{") - strings.Count(l, "}")
		bracket := strings.Count(l, "[") - strings.Count(l, "]")
		parens := strings.Count(l, "(") - strings.Count(l, ")")
		for curly != 0 && bracket != 0 && parens != 0 {
			l, err = r.ReadString('\n')
			if err != nil {
				log.Panic(err)
			}
			w.WriteString(l)
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
		w.WriteString(fmt.Sprintf("goto %v\n", label))
	}
	w.WriteString(fmt.Sprintf("goto %v\n", label))
	w.Flush()
	return string(buf.Bytes())
}

func parseFile(f io.Reader) {
	trim := regexp.MustCompile(`\s*(.*)`)
	b := bufio.NewReader(f)
	buf := bytes.NewBuffer([]byte{})
	o := bufio.NewWriter(buf)
	for l, err := b.ReadString('\n'); err == nil; l, err = b.ReadString('\n') {
		trimmed := trim.ReplaceAllString(l, "$1")
		if strings.HasPrefix(trimmed, "search") {
			l = build(l, b)
		}
		o.WriteString(l)
	}
	o.Flush()
	log.Println(string(buf.Bytes()))
}

func main() {
	// parseFile(strings.NewReader(example))
	f := `row, column := candidate.Row, candidate.Column
			for _, q := range solution {
			    r, c := q.Row, q.Column
			    if row == r ||
			        column == c ||
			        row+column == r+c ||
			        row-column == r-c {
			        return true
			    }
			}
			return false`
	log.Println(inlineFunction(f, "__01_reject", "END_REJECT"))
}
