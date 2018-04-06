'''
Wikipedia sucks...unless it doesn't. This is the N-Queens problem implemented
using the following general-form pseudocode.

https://en.wikipedia.org/wiki/Backtracking#Pseudocode

It is getting correct answers, although I am doing something wrong and also
getting duplicates.
'''

N = 8


def backtrack(P, root, accept, reject, first, next, output):
    if reject(P, root):
        return
    P.append(root)
    if accept(P):
        output(P)
        P.pop()
        return
    s = first(root)
    while s is not None:
        backtrack(P, s, accept, reject, first, next, output)
        s = next(s)
    P.pop()
    s = next(root)
    if s is not None:
        backtrack(P, s, accept, reject, first, next, output)


def root():
    # Return the partial candidate at the root of the search tree.
    # First spot on the board.
    return 1, 1


def reject(P, candidate):
    C, R = candidate
    for column, row in P:
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


def first(root):
    # Generate the first extension of candidate c.
    # That is, if we have a placed queen at 1, 1 (C, R) then this should return 2, 1.
    return None if root[0] >= N else (root[0] + 1, 1)


def next(child):
    # generate the next alternative extension of a candidate, after the extension s.
    # So if first generated 2, 1 then this should produce 2, 2.
    return None if child[1] >= N else (child[0], child[1] + 1)

def output(P):
    print(P)


def more(R):
    return R < N


if __name__ == "__main__":
    backtrack([], root(), accept, reject, first, next, output)
    print(derp)
