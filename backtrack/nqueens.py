from backtrack import backtrack

N = 4


def first(candidate):
    C, R = candidate
    # Generate the first extension of candidate c.
    # That is, if we have a placed queen at 1, 1 (C, R) then this should return 2, 1.
    return None if C >= N else (C + 1, 1)


def next(candidate):
    C, R = candidate
    # generate the next alternative extension of a candidate, after the extension s.
    # So if first generated 2, 1 then this should produce 2, 2.
    return None if R >= N else (C, R + 1)


def accept(P):
    # Return true if c is a solution of P, and false otherwise.
    # The latest queen has been verified safe, so if we've placed N queens, we win.
    return len(P) == N

def reject(P, candidate):
    C, R = candidate
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


if __name__ == '__main__':
    backtrack((1, 1), first, next, reject, accept, add, remove, output)