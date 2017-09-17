import copy
import time
N = 4

# backtrack r = root(); c = first(); c = next() {
#     switch
#         case reject(c):
#             break
#         case accept:
#             output(c)
#             break
#         case continue:
#             break
# }

# root = (0, N  + 1)

# backtrack on root {
#     procedures {
#         reject:
#             for column, row in P:
#                 if (R == row or
#                         C == column or
#                         R + C == row + column or
#                         R - C == row - column):
#                     return True
#             return False
#     }
# }

def backtrack(root, first, next, reject, accept, add, remove):
    stack = list()
    structure = list([root])

    while root is not None: 
        candidate = first(*root)
        while candidate is not None:
            if reject(structure, *candidate):
                candidate = next(*candidate)
                continue
            add(structure, candidate)
            if accept(structure):
                output(structure)
                remove(structure)
                candidate = next(*candidate)
                continue
            stack.append(root)
            root = candidate
            break
        else:
            root = next(*root)
            remove(structure)
            while stack:
                if root is None:
                    root = next(*stack.pop())
                    remove(structure)
                    continue
                if reject(structure, *root):
                    root = next(*root)
                    continue
                break
            add(structure, root)

        





def first(C, R):
    # Generate the first extension of candidate c.
    # That is, if we have a placed queen at 1, 1 (C, R) then this should return 2, 1.
    return None if C >= N else (C + 1, 1)


def next(C, R):
    # generate the next alternative extension of a candidate, after the extension s.
    # So if first generated 2, 1 then this should produce 2, 2.
    return None if R >= N else (C, R + 1)


def accept(P):
    # Return true if c is a solution of P, and false otherwise.
    # The latest queen has been verified safe, so if we've placed N queens, we win.
    return len(P) == N

def reject(P, C, R):
    # Return true only if the partial candidate c is not worth completing.
    # c, in this case, is Column and Row together.
    # column = 0
    for column, row in P:
        # column += 1
        if (R == row or
                C == column or
                R + C == row + column or
                R - C == row - column):
            return True
    return False



def output(P):
    # use the solution c of P, as appropriate to the application.
    # unique.add(tuple(P))
    print(P)


def add(P, candidate):
    P.append(candidate)


def remove(P):
    P.pop()



RED = 0
BLUE = 1
GREEN = 2


MAP = [(0, 1), (1, 2), (2, 3), (3, 0)]

def colorReject(PP, nodeN, colorN):
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

def colorAccept(P):
    return len(P) == 4


def colorFirst(node, color):
    return (node + 1, RED)


def colorNext(node, color):
    return None if color >= GREEN else (node, color + 1)


def colorAdd(P, candidate):
    P.append(candidate)


def colorRemove(P):
    P.pop()


backtrack([0, RED], colorFirst, colorNext, colorReject, colorAccept, colorAdd, colorRemove)

# print(colorReject(
#         [[0, RED], [1, BLUE], [2, RED], [3, GREEN]]
#     ))
print("NQueens")
backtrack((1, 1), first, next, reject, accept, add, remove)
