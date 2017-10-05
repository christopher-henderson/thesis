from __future__ import division
from nqueens import NQueens


def backtrack(root, single_solution=False):
    root_stack = list()
    solution = list()

    solution.append(root)
    while root is not None: 
        candidate = None if root.column >= root.N else NQueens(root.column + 1, 1)
        while candidate is not None:
            if any(queen.can_attack(candidate) for queen in solution):
                candidate = None if candidate.row >= candidate.N else NQueens(candidate.column, candidate.row + 1)
                continue
            solution.append(candidate)
            if len(solution) == candidate.N:
                print(solution)
                if single_solution:
                    return
                solution.pop()
                candidate = None if candidate.row >= candidate.N else NQueens(candidate.column, candidate.row + 1)
                continue
            root_stack.append(root)
            root = candidate
            break
        else:
            root = None if root.row >= root.N else NQueens(root.column, root.row + 1)
            solution.pop()
            while root_stack:
                if root is None:
                    parent = root_stack.pop()
                    root = None if parent.row >= parent.N else NQueens(parent.column, parent.row + 1)
                    solution.pop()
                    continue
                if any(queen.can_attack(root) for queen in solution):
                    root = None if root.row >= root.N else NQueens(root.column, root.row + 1)
                    continue
                break
            if root is not None:
                solution.append(root) 

from time import time
start = time()
backtrack(NQueens(1, 1))
inline = time() - start

from nqueens import main
start = time()
main()
framework = time() - start

print(inline / framework)