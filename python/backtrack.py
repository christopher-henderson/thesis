# def reject(solution, candidate):
#     column, row = candidate
#     return any(
#         row == r or
#         column == c or
#         row + column == r + c or 
#         row - column == r - c for r, c in solution)

# def accept(solution):
#     return globals(), locals(), len(solution) == N

# def children(parent):
#     column, _ = parent
#     for row in range(1, N + 1):
#         yield tuple([column + 1, row])

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
        return None

# def backtrack(root, first, next, reject, accept, add=add, remove=remove, output=output, single_solution=False):
def backtrack(root):
    N = 4

    root_stack = list()
    solution = list()

    def reject(solution, candidate):
        column, row = candidate
        return any(
            row == r or
            column == c or
            row + column == r + c or 
            row - column == r - c for c, r in solution)

    def accept(solution):
        if len(solution) == N:
            print(solution)

    def children(parent):
        column, _ = parent
        for row in range(1, N + 1):
            yield tuple([column + 1, row])

    __children = children(root)

    while True:
        # print(root)
        candidate = __next(__children)
        if candidate is None:
            if len(root_stack) is 0:
                break
            remove(solution)
            root, __children = root_stack.pop()
            continue
        __reject = reject(solution, candidate)
        if __reject:
            continue
        add(solution, candidate)
        __accept = accept(solution)
        if __accept:
            remove(solution)
            continue
        root_stack.append([root, __children])
        root = candidate
        __children = children(root)

def go():
    backtrack((0, 1))

go()