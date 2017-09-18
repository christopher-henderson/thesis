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

        children = [direction for direction in [west, north, east, south] if self.inbounds(*direction)]
        if len(children) > 0:
            # Select the first inbounds child and construct it with its siblings.
            return MazeProblem(children[0][0], children[0][1], children[1:])
        return None

    def next(self):
        if len(self.siblings) is 0:
            return None
        # Construct the nearest sibling with the rest of its siblings.
        return MazeProblem(self.siblings[0][0], self.siblings[0][1], self.siblings[1:])

    @classmethod
    def reject(cls, P, candidate):
        # It's been seen before or it's a wall.
        return candidate in cls.VISITED or cls.MAZE[candidate.R][candidate.C] is 0

    @classmethod
    def accept(cls, P):
        # The last node is a node on the boundary of the maze, thus an exit.
        return cls.is_a_goal(P[-1].R, P[-1].C)

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

    @classmethod
    def is_a_goal(cls, R, C):
        # The node is a 1 and it is on the border of the maze, thus it
        # can be used as either an entrance or an exit.
        return (cls.MAZE[R][C] is 1 and (
                    R is 0 or R is len(MazeProblem.MAZE) - 1 or
                    C is 0 or C is len(MazeProblem.MAZE[R]) - 1
                    )
                )

    @staticmethod
    def output(P):
        print(P)


    # Pretty printing.
    def __str__(self):
        return str([self.R, self.C])

    def __repr__(self):
        return str(self)

    # These two are just because I wanted throw whole objects into the VISITED hash set.
    def __hash__(self):
        return hash(tuple([self.R, self.C]))

    def __eq__(self, other):
        if type(other) == MazeProblem:
            return self.R == other.R and self.C == other.C
        return self.R == other[0] and self.C == other[1]


if __name__ == '__main__':
    MazeProblem.MAZE = MazeProblem.SIMPLE
    # Construct a list of all entrances and exits. An entrance/exit
    # is any 1 node that is on the edge of the maze.
    entrances = [
        [R, C]
        for R in range(len(MazeProblem.MAZE))
        for C in range(len(MazeProblem.MAZE[R]))
        if MazeProblem.is_a_goal(R, C)
    ]
    start = MazeProblem(entrances[0][0], entrances[0][1], entrances[1:])
    backtrack(
        start, 
        MazeProblem.first, 
        MazeProblem.next, 
        MazeProblem.reject, 
        MazeProblem.accept, 
        MazeProblem.add, 
        MazeProblem.remove, 
        MazeProblem.output)