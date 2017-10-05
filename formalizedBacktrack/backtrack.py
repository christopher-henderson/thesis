def _add(solution, candidate):
    solution.append(candidate)

def _remove(solution):
    solution.pop()

def _output(solution):
    print(solution) 


ACCEPT = '__accept__'
REJECT = '__reject__'
CHILD = '__child__'
SIBLING = '__sibling__'
ADD = '__add__'
REMOVE = '__remove__'
OUTPUT = '__output__'
SINGLE_SOLUTION = '__single_solution__'

def backtrack(root):
    # Procedure resolutions.
    accept = getattr(type(root), ACCEPT) 
    reject = getattr(type(root), REJECT)
    child = getattr(type(root), CHILD)
    sibling = getattr(type(root), SIBLING)
    add = getattr(type(root), ADD) if hasattr(root, ADD) else _add
    remove = getattr(type(root), REMOVE) if hasattr(root, REMOVE) else _remove
    output = getattr(type(root), OUTPUT) if hasattr(root, OUTPUT) else _output

    single_solution = getattr(type(root), SINGLE_SOLUTION) if hasattr(root, SINGLE_SOLUTION) else False

    # Stack initialization.
    root_stack = list()
    solution = list()

    # Algorithm.
    add(solution, root)
    while root is not None: 
        candidate = child(root)
        while candidate is not None:
            if reject(solution, candidate):
                candidate = sibling(candidate)
                continue
            add(solution, candidate)
            if accept(solution):
                output(solution)
                if single_solution:
                    return
                remove(solution)
                candidate = sibling(candidate)
                continue
            root_stack.append(root)
            root = candidate
            break
        else:
            root = sibling(root)
            remove(solution)
            while root_stack:
                if root is None:
                    root = sibling(root_stack.pop())
                    remove(solution)
                    continue
                if reject(solution, root):
                    root = sibling(root)
                    continue
                break
            if root is not None:
                add(solution, root)  