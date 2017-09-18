from backtrack import backtrack

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

    @staticmethod
    def output(P):
        print(P)

    def __str__(self):
        return str([self.R, self.C])

    def __repr__(self):
        return str(self)


if __name__ == '__main__':
    backtrack(MazeProblem(0, 0, []), MazeProblem.first, MazeProblem.next, MazeProblem.reject, MazeProblem.accept, MazeProblem.add, MazeProblem.remove, MazeProblem.output)