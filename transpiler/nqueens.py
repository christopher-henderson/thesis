N = 8
backtrack on (1, 1):
	reject(solution, candidate):
		column, row = candidate
		return any(
			row == r or
			column == c or
			row + column == r + c or 
			row - column == r - c for r, c in solution)
	accept(solution):
		return len(solution) == N
	children(parent):
		column, _ = parent
		for row in range(1, N + 1):
			yield tuple([column, row])


def NQueens(N):
	backtrack on (1, 1):
		reject(solution, candidate):
			column, row = candidate
			return any(
				row == r or
				column == c or
				row + column == r + c or 
				row - column == r - c for r, c in solution)
		accept(solution):
			return len(solution) == N
		children(parent):
			column, _ = parent
			for row in range(1, N + 1):
				yield tuple([column, row])

for N in range(1, 9):
	backtrack on (1, 1):
		reject(solution, candidate):
			column, row = candidate
			return any(
				row == r or
				column == c or
				row + column == r + c or 
				row - column == r - c for r, c in solution)
		accept(solution):
			return len(solution) == N
		children(parent):
			column, _ = parent
			for row in range(1, N + 1):
				yield tuple([column, row])