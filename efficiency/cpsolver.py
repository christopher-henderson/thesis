"""
  OR-tools solution to the N-queens problem.
"""
from __future__ import print_function
import sys
from ortools.constraint_solver import pywrapcp

def main(board_size):
  # Creates the solver.
  solver = pywrapcp.Solver("n-queens")
  # Creates the variables.
  # The array index is the column, and the value is the row.
  queens = [solver.IntVar(0, board_size - 1, "x%i" % i) for i in range(board_size)]
  # Creates the constraints.

  # All rows must be different.
  solver.Add(solver.AllDifferent(queens))

  # All columns must be different because the indices of queens are all different.

  # No two queens can be on the same diagonal.
  solver.Add(solver.AllDifferent([queens[i] + i for i in range(board_size)]))
  solver.Add(solver.AllDifferent([queens[i] - i for i in range(board_size)]))

  db = solver.Phase(queens,
                    solver.CHOOSE_FIRST_UNBOUND,
                    solver.ASSIGN_MIN_VALUE)
  solver.NewSearch(db)

  # Iterates through the solutions, displaying each.
  num_solutions = 0

  while solver.NextSolution():
    pass
    num_solutions += 1

  solver.EndSearch()

  return solver.WallTime()
  # csv.write("{SIZE},{TIME}\n".format(SIZE=board_size, TIME=solver.WallTime()))
  # print()
  # print("Solutions found:", num_solutions)
  # print("Time:", solver.WallTime(), "ms")

# By default, solve the 8x8 problem.
board_size = 8

if __name__ == "__main__":
  with open("google_timings.csv", "w+") as csv:
    for i in range(1, 16):
      totalTime = 0
      for j in range(40):
        totalTime += main(i)
      csv.write("{SIZE},{TIME}\n".format(SIZE=i, TIME=totalTime//40))
      csv.flush()