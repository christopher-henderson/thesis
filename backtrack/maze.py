from backtrack import backtrack

class MazeProblem(object):


    SIMPLE = [
        [0, 0, 0, 0, 0, 0],
        [1, 1, 0, 0, 0, 0],
        [0, 1, 0, 0, 0, 0],
        [0, 1, 1, 1, 1, 0],
        [0, 1, 0, 1, 0, 0],
        [0, 0, 0, 1, 0, 0]
    ]

    CYCLE = [
        [0, 0, 0, 0, 0, 0],
        [1, 1, 1, 1, 0, 0],
        [0, 1, 0, 1, 0, 0],
        [0, 1, 1, 1, 1, 0],
        [0, 1, 0, 1, 0, 0],
        [0, 0, 0, 1, 0, 0]
    ]


    MULTIPLE_ENTRANCIES_EXITS = [
        [0, 0, 0, 0, 0, 0],
        [1, 1, 0, 0, 0, 0],
        [0, 1, 0, 0, 0, 0],
        [0, 1, 1, 1, 1, 1],
        [0, 1, 0, 1, 0, 0],
        [0, 1, 0, 1, 0, 0]
    ]

    VISITED = set()

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
        R, C = P[-1].R, P[-1].C
        return (R is 0 or R is len(MazeProblem.MAZE) - 1 or
                C is 0 or C is len(MazeProblem.MAZE[R]) - 1)

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
    MazeProblem.MAZE = MazeProblem.MULTIPLE_ENTRANCIES_EXITS
    entrances = [
        [R, C]
        for R in range(len(MazeProblem.MAZE))
        for C in range(len(MazeProblem.MAZE[R]))
        if MazeProblem.MAZE[R][C] is 1 and
        (R is 0 or R is len(MazeProblem.MAZE) - 1 or
        C is 0 or C is len(MazeProblem.MAZE[R]) - 1)
    ]
    start = MazeProblem(entrances[0][0], entrances[0][1], entrances[1:])
    backtrack(start, MazeProblem.first, MazeProblem.next, MazeProblem.reject, MazeProblem.accept, MazeProblem.add, MazeProblem.remove, MazeProblem.output)