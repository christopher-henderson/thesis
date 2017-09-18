from backtrack import backtrack
import copy

RED = 0
BLUE = 1
GREEN = 2

MAP = [(0, 1), (1, 2), (2, 3), (3, 0)]


def reject(PP, N):
    nodeN, colorN = N
    P = copy.deepcopy(PP)
    P.append([nodeN, colorN])
    for node, color in P:
        for source, destination in MAP:
            if source >= len(P) or destination >= len(P):
                continue
            if node == source and P[destination][1] == color:
                return True
            if node == destination and P[source][1] == color:
                return True
    return False


def accept(P):
    return len(P) == 4


def first(candidate):
    node, color = candidate
    return (node + 1, RED)


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