def add(solution, candidate):
    solution.append(candidate)

def remove(solution):
    solution.pop()

def output(solution):
    print(solution)

def backtrack(root, first, next, reject, accept, add=add, remove=remove, output=output):
    root_stack = list()
    solution = list()

    add(solution, root)
    while root is not None: 
        candidate = first(root)
        while candidate is not None:
            if reject(solution, candidate):
                candidate = next(candidate)
                continue
            add(solution, candidate)
            if accept(solution):
                output(solution)
                remove(solution)
                candidate = next(candidate)
                continue
            root_stack.append(root)
            root = candidate
            break
        else:
            root = next(root)
            remove(solution)
            while root_stack:
                if root is None:
                    root = next(root_stack.pop())
                    remove(solution)
                    continue
                if reject(solution, root):
                    root = next(root)
                    continue
                break
            if root is not None:
                add(solution, root)  
