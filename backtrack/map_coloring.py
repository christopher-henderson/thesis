from backtrack import backtrack
import copy

RED = 0
BLUE = 1
GREEN = 2

MAP = [(0, 1), (1, 2), (2, 3), (3, 0)]


def reject(P, candidate):
    vertex, color = candidate
    # Find all edges that contain the candidate vertex.
    edges = (edge for edge in MAP if vertex in edge)
    # Find all neighbors to the candidate vertex.
    neighbors = (source if source != vertex else destination for source, destination in edges)
    # Return True if any neighbor has already been colored the same color as the candidate vertex.
    return any(neighbor < len(P) and P[neighbor][1] == color for neighbor in neighbors)


def accept(P):
    return len(P) == len(MAP)


def first(candidate):
    vertex, color = candidate
    return None if vertex >= len(MAP) - 1 else (vertex + 1, RED)


def next(candidate):
    vertex, color = candidate
    return None if color >= GREEN else (vertex, color + 1)


def add(P, candidate):
    P.append(candidate)


def remove(P):
    P.pop()


def output(P):
    print(P)


if __name__ == '__main__':
    backtrack([0, RED], first, next, reject, accept, add, remove, output)