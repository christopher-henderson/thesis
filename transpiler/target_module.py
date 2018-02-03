def reject(g, l, solution, candidate):
    exec("") in locals(); globals().update(g); locals().update(l)
    column, row = candidate
    return globals(), locals(), any(
        row == r or
        column == c or
        row + column == r + c or 
        row - column == r - c for r, c in solution)
        

def accept(g, l, solution):
    exec("") in locals(); globals().update(g); locals().update(l)
    return globals(), locals(), len(solution) == N
        

def children(g, l, parent):
    exec("") in locals(); globals().update(g); locals().update(l)
    column, _ = parent
    for row in range(1, N + 1):
        yield globals(), locals(), tuple([column + 1, row])

def add(solution, candidate):
    solution.append(candidate)

def remove(solution):
    solution.pop()

def output(solution):
    print(solution)

def __next(gen):
    try:
        return next(gen)
    except StopIteration:
        return None, None, None

# def backtrack(root, first, next, reject, accept, add=add, remove=remove, output=output, single_solution=False):
def backtrack(root, g, l):
    root_stack = list()
    solution = list()

    add(solution, root)
    __children = children(g, l, root)
    g, l, candidate = __next(__children)
    root_stack.append([root, __children])
    while root is not None:
        while candidate is not None:
            print(candidate)
            g, l, __reject = reject(g, l, solution, candidate)
            if __reject:
                g, l, candidate = __next(__children)
                continue
            add(solution, candidate)
            print(solution)
            g, l, __accept = accept(g, l, solution)
            if __accept:
                output(solution)
                return
            root_stack.append([root, __children])
            root = candidate
            __children = children(g, l, root)
            g, l, candidate = __next(__children)
        try:
            root, __children = root_stack.pop()
            remove(solution)
        except IndexError:
            root, __children = None, None
        else:
            g, l, candidate = __next(__children)

def go():
    N = 8
    backtrack((1, 1), globals(), locals())

go()