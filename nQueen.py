'''
Wikipedia sucks...unless it doesn't. This is the N-Queens problem implemented
using the following general-form pseudocode.

https://en.wikipedia.org/wiki/Backtracking#Pseudocode

It is getting correct answers, although I am doing something wrong and also
getting duplicates.
'''

N = 8
unique = set()


def backtrack(P, C, R):
    if reject(P, C, R):
        return
    # This append is not a part of the general form. Should it be in 'reject'? That is:
    #
    # 'reject' returns true only if the candicate c and all of its subtrees
    # are to be rejected. Otherwise it should update P to include c and return
    # false.
    #
    # That would be kinda odd, and changes the genera form shown to us but the
    # structure must be updated before making a call to accept.
    P.append(R)
    if accept(P):
        output(P)
        P.pop()
        return
    s = first(C)
    while s is not None:
        backtrack(P, *s)
        s = next(*s)
    # The general form admits this case. We could always add to the form another
    # procedure that takes in P and c and returns whether more candidate solutions
    # may exist. In this case I added the 'more' procedure.
    P.pop()
    s = next(C, R)
    if s is not None:
        backtrack(P, *s)
    # if more(R):
    #     backtrack(P, *next(C, R))


def root():
    # Return the partial candidate at the root of the search tree.
    # First spot on the board.
    return 1, 1


def reject(P, C, R):
    # Return true only if the partial candidate c is not worth completing.
    # c, in this case, is Column and Row together.
    column = 0
    for row in P:
        column += 1
        if (R == row or
                C == column or
                R + C == row + column or
                R - C == row - column):
            return True
    return False


def accept(P):
    # Return true if c is a solution of P, and false otherwise.
    # The latest queen has been verified safe, so if we've placed N queens, we win.
    return len(P) == N


def first(C):
    # Generate the first extension of candidate c.
    # That is, if we have a placed queen at 1, 1 (C, R) then this should return 2, 1.
    return None if C >= N else (C + 1, 1)


def next(C, R):
    # generate the next alternative extension of a candidate, after the extension s.
    # So if first generated 2, 1 then this should produce 2, 2.
    return None if R >= N else (C, R + 1)


def output(P):
    # use the solution c of P, as appropriate to the application.
    unique.add(tuple(P))
    print(P)


def more(R):
    return R < N


if __name__ == "__main__":
    backtrack([], *root())
    print(len(unique))
