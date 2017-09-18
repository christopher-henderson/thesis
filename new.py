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

# def backtrack(root, first, next, reject, accept, add, remove):
#     stack = list()
#     structure = list([root])

#     while root is not None: 
#         candidate = first(*root)
#         while candidate is not None:
#             if reject(structure, *candidate):
#                 candidate = next(*candidate)
#                 continue
#             add(structure, candidate)
#             if accept(structure):
#                 output(structure)
#                 remove(structure)
#                 candidate = next(*candidate)
#                 continue
#             stack.append(root)
#             root = candidate
#             break
#         else:
#             root = next(*root)
#             remove(structure)
#             while stack:
#                 if root is None:
#                     root = next(*stack.pop())
#                     remove(structure)
#                     continue
#                 if reject(structure, *root):
#                     root = next(*root)
#                     continue
#                 break
#             add(structure, root)


def backtrack(root, first, next, reject, accept, add, remove):
    stack = list()
    structure = list([root])

    while root is not None: 
        candidate = first(root)
        while candidate is not None:
            if reject(structure, candidate):
                candidate = next(candidate)
                continue
            add(structure, candidate)
            if accept(structure):
                output(structure)
                remove(structure)
                candidate = next(candidate)
                continue
            stack.append(root)
            root = candidate
            break
        else:
            root = next(root)
            remove(structure)
            while stack:
                if root is None:
                    root = next(stack.pop())
                    remove(structure)
                    continue
                if reject(structure, root):
                    root = next(root)
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
    print(1, P)


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


# backtrack([0, RED], colorFirst, colorNext, colorReject, colorAccept, colorAdd, colorRemove)

# print(colorReject(
#         [[0, RED], [1, BLUE], [2, RED], [3, GREEN]]
#     ))
# print("NQueens")
# backtrack((1, 1), first, next, reject, accept, add, remove)




class MazeProblem(object):

    MAZE = [
        [1, 0, 0, 0],
        [1, 1, 1, 1],
        [0, 1, 0, 1],
        [1, 1, 1, 1]
    ]

    VISITED = set()
    VISITED.add((0, 0))

    def __init__(self, R, C, siblings):
        self.R = R
        self.C = C
        self.siblings = siblings

    def first(self):
        add = lambda a, b: a + b
        west = tuple(map(add, [0, -1], [self.R, self.C]))
        north = tuple(map(add, [-1, 0], [self.R, self.C]))
        east = tuple(map(add, [0, 1], [self.R, self.C]))
        south = tuple(map(add, [1, 0], [self.R, self.C]))

        options = [option for option in [west, north, east, south] if self.inbounds(*option)]
        if len(options) > 0:
            return MazeProblem(options[0][0], options[0][1], options[1:])
        return None

    def next(self):
        if len(self.siblings) is 0:
            return None
        return MazeProblem(self.siblings[0][0], self.siblings[0][1], self.siblings[1:])

    def __hash__(self):
        return hash(tuple([self.R, self.C]))

    def __eq__(self, other):
        if type(other) == MazeProblem:
            return self.R == other.R and self.C == other.C
        return self.R == other[0] and self.C == other[1]

    @classmethod
    def reject(cls, P, candidate):
        return (candidate in cls.VISITED or 
                not cls.inbounds(candidate.R, candidate.C) or
                cls.MAZE[candidate.R][candidate.C] is 0
            )

    @staticmethod
    def accept(P):
        return P[-1] == (3, 3)

    @classmethod
    def inbounds(cls, R, C):
        return R >= 0 and C >=0 and R < len(cls.MAZE) and C < len(cls.MAZE[0])

    @classmethod
    def add(cls, P, candidate):
        cls.VISITED.add(candidate)
        P.append(candidate)


    @classmethod
    def remove(cls, P):
        candidate = P.pop()
        cls.VISITED.remove(candidate)


    def __str__(self):
        return str([self.R, self.C])


    def __repr__(self):
        return str(self)

backtrack(MazeProblem(0, 0, []), MazeProblem.first, MazeProblem.next, MazeProblem.reject, MazeProblem.accept, MazeProblem.add, MazeProblem.remove)
