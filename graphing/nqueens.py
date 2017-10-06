from backtrack import backtrack



class NQueens(object):

    N = 4

    def __init__(self, column, row):
        self.column = column
        self.row = row


    def can_attack(self, other):
        return (
            other.row == self.row or
            other.column == self.column or
            other.row + other.column == self.row + self.column or
            other.row - other.column == self.row - self.column
        )


    # @ self.C >= self.N
    # def __first__(self):
    #     return None

    def __child__(self):
        # Generate the first extension of candidate c.
        # That is, if we have a placed queen at 1, 1 (C, R) then this should return 2, 1.
        return None if self.column >= self.N else NQueens(self.column + 1, 1)

    # @ self.R >= self.N
    # def __sibing__(self):
    #     return None

    def __sibling__(self):
        return None if self.row >= self.N else NQueens(self.column, self.row + 1)

    @classmethod
    def __accept__(cls, solution):
        return len(solution) == cls.N

    @staticmethod
    def __reject__(solution, candidate):
        return any(queen.can_attack(candidate) for queen in solution)

    def __str__(self):
        return str(tuple([self.column, self.row]))

    def __repr__(self):
        return str(self)


def main():
    # backtrack NQueens(1, 1)
    backtrack(NQueens(1, 1))

if __name__ == '__main__':
    main()
    