from backtrack import backtrack
import copy

RED = 0
BLUE = 1
GREEN = 2

MAP = [(0, 1), (1, 2), (2, 3), (3, 0)]


def reject(P, candidate):
    node, color = candidate
    edges = (edge for edge in MAP if node in edge)
    for source, destination in edges:
        neighbor = source if source != node else destination
        if neighbor >= len(P):
            continue
        if P[neighbor][1] == color:
            return True
    return False


def accept(P):
    return len(P) == 4


def first(candidate):
    node, color = candidate
    return None if node >= 3 else (node + 1, RED)


def next(candidate):
    node, color = candidate
    return None if color >= GREEN else (node, color + 1)


def add(P, candidate):
    P.append(candidate)


def remove(P):
    P.pop()


def output(P):
    print(P)


if __name__ == '__main__':
    backtrack([0, RED], first, next, reject, accept, add, remove, output)