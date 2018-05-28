#include <stdio.h>
#include <stdlib.h>
using namespace std;

const int N = 4;

// type __GraphNode struct {
// 	Active	 bool
// 	ID       int
// 	Parent   int
// }
// // User CHILDREN declaration.
// USER_children := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) <-chan {UTYPE} {
// 	return make(chan {UTYPE}, 0)
// }
// // User ACCEPT declaration.
// USER_accept := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) bool {
// 	return false
// }
// // User REJECT declaration.
// USER_reject := func(node {UTYPE}, solution []{UTYPE}, gid *__GraphNode) bool {
// 	return false
// }
// root := changeme
// maxgoroutine := 1
// // Parent:Children PODO meant for stack management.
// type StackEntry struct {
// 	Parent   {UTYPE}
// 	Children <-chan {UTYPE}
// }
// lock := make(chan int, maxgoroutine)
// wg := make(chan int, maxgoroutine)
// ticket := make(chan int, maxgoroutine)
// // You have to declare first since the function can fire off a
// // goroutine of itself.
// var engine func(solution []{UTYPE}, root {UTYPE}, gid *__GraphNode)
// engine = func(solution []{UTYPE}, root {UTYPE}, gid *__GraphNode) {
// 	_children := USER_children(root, solution, gid)
// 	// Stack of Parent:Chidren pairs.
// 	stack := make([]StackEntry, 0)
// 	// Current candidate under consideration.
// 	var candidate {UTYPE}
// 	// Holds a StackEntry.
// 	var stackEntry StackEntry
// 	// Generic boolean variable
// 	var ok bool
// 	for {
// 		if candidate, ok = <-_children; !ok {
// 			// This node has no further children.
// 			if len(stack) == 0 {
// 				// Algorithm termination. No further nodes in the stack.
// 				break
// 			}
// 			// With no valid children left, we pop the latest node from the solution.
// 			solution = solution[:len(solution)-1]
// 			// Pop from the stack. Broken into two steps:
// 			// 	1. Get final element.
// 			//	2. Resize the stack.
// 			stackEntry = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			// Extract root and candidate fields from the StackEntry.
// 			root = stackEntry.Parent
// 			_children = stackEntry.Children
// 			continue
// 		}
// 		// Ask the user if we should reject this candidate.
// 		_reject := USER_reject(candidate, solution, gid)
// 		if _reject {
// 			// Rejected candidate.
// 			continue
// 		}
// 		// Append the candidate to the solution.
// 		solution = append(solution, candidate)
// 		// Ask the user if we should accept this solution.
// 		_accept := USER_accept(candidate, solution, gid)
// 		if _accept {
// 			// Accepted solution.
// 			// Pop from the solution thus far and continue on with the next child.
// 			solution = solution[:len(solution)-1]
// 			continue
// 		}
// 		select {
// 			case lock <- 1:
// 				wg <- 1
// 				s := make([]{UTYPE}, len(solution))
// 				copy(s, solution)
// 				go engine(s, candidate, &__GraphNode{Active: true, ID: <-ticket, Parent: gid.ID})
// 				// pretend we didn't see this
// 				solution = solution[:len(solution)-1]
// 				continue
// 			default:
// 		}
// 		// Push the current root to the stack.
// 		stack = append(stack, StackEntry{root, _children})
// 		// Make the candidate the new root.
// 		root = candidate
// 		// Get the new root's children channel.
// 		_children = USER_children(root, solution, gid)
// 	}
// 	<- lock
// 	wg <- -1
// 	gid.Active = false
// }
// shutdown := make(chan struct{}, 0)
// go func() {
// 	// Goroutine ticketing system.
// 	id := 0
// 	for {
// 		select {
// 		case ticket <- id:
// 			id++
// 		case <-shutdown:
// 			close(ticket)
// 			return
// 		}
// 	}
// }()
// lock <- 1
// wg <- 1
// go engine(make([]{UTYPE}, 0), root, &__GraphNode{Active: true, ID: <-ticket, Parent: 0})
// count := 0
// for c := range wg {
// 	count += c
// 	if count == 0 {
// 		break
// 	}
// }
// close(shutdown)
// close(wg)
	// close(lock)

typedef struct queen {
	int Row;
	int Column;
} Queen;

// typedef struct iter {
// 	Queen * Current;
// 	int Size;
// } Iter;

// Iter children(Queen * parent, int size) {
// 	Iter i;
// 	i.Size = size;
// 	i.Current = malloc(sizeof(Queen));
// 	i.Current->Row = 1;
// 	i.Current->Column = parent->Column + 1;
// 	return i;
// }

// Queen * next(Iter * iter) {
// 	if (iter->Current->Column > iter->Size || iter->Current->Row > iter->Size) {
// 		return NULL;
// 	}
// 	Queen * ret = iter->Current;
// 	iter->Current = malloc(sizeof(Queen));
// 	iter->Current->Row = ret->Row + 1;
// 	iter->Current->Column = ret->Column;
// 	return ret;
// }

typedef struct stacke {
	Queen * Parent;
	Iter Iter;
} StackEntry;


int reject(Queen * queen, Solution * solution) {
	int column = queen->Column;
	int row = queen->Row;
	for (int i = 0; i < solution->pointer; i++) {
		int r = getS(solution, i)->Row;
		int c = getS(solution, i)->Column;
		if ((row == r) || (column == c) || (row + column == r + c) || (row - column == r - c)) {
			return 1;
		}
	}
	return 0;
}

int accept(Solution * solution) {
	return solution->pointer == N;
}

void backtrack(Queen * FCG, int size) {
	Iter iter;
	Queen * root;
	Queen * candidate;
	StackEntry se;
	
	Solution * solution = NewSolution(size);
	Stack * stack = NewStack(size);
	
	iter = children(FCG, size);
	root = FCG;
	while (1) {
		candidate = next(&iter);
		printf("%i %i\n", root->Column, root->Row);
		printf("%i %i\n", candidate->Column, candidate->Row);
		if (candidate == NULL) {
			if (empty(stack)) {
				printf("%s\n", "We lose?");
				break;
			}
			popS(solution);
			se = pop(stack);
			root = se.Parent;
			iter = se.Iter;
			continue;
		}
		if (reject(candidate, solution)) {
			// free(candidate);
			continue;
		}
		// printf("%i %i\n", candidate->Column, candidate->Row);
		pushS(solution, candidate);
		if (accept(solution)) {
			printf("%s\n", "We win?");
			candidate = popS(solution);
			// free(candidate);
			continue;
		}
		se.Parent = root;
		se.Iter = iter;
		push(stack, se);
		root = candidate;
		iter = children(root, size);
	}
}


int main(int argc, char const *argv[]) {
	Queen FCG;
	FCG.Column = 0;
	FCG.Row = 0;
	backtrack(&FCG, N);
	return 0;
}