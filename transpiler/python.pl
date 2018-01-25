N = 8
backtrack on (1, 1):
    # Is the given candidate valid in the presence
    # of the solution computed thus far?
    reject(solution, candidate):
        column, row = candidate
        return any(
            row == r or
            column == c or
            row + column == r + c or 
            row - column == r - c for r, c in solution)
    # Is the solution computed thus far a complete solution?
    accept(solution):
        return len(solution) == N
    # Given a parent node, what are all of its children, if any?
    children(parent):
        column, _ = parent
        for row in range(1, N + 1):
            yield tuple([column, row])