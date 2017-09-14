import copy

N = 8
NUM_FOUND = 0
derp = list()


def backtrack(P, C, R):
    if reject(P, C, R):
        return
    # Not a part of the general form. Should it be in 'reject'?
    # That would be kinda odd, but it must appended before a call to output.
    P.append([C, R])
    if accept(P):
        output(P)
    s = first(C)
    while s is not None:
        # Explicitly passing by value (deepcopy a list) is another addition of
        # mine. Is this mandatory?
        backtrack(copy.deepcopy(P), *s)
        s = next(*s)
    # The general form admits this case, although there can be a way around it.
    if R < N:
        P.pop()
        backtrack(copy.deepcopy(P), C, R + 1)


def root():
    # Return the partial candidate at the root of the search tree.
    # First spot on the board.
    return 1, 1


def reject(P, C, R):
    # Return true only if the partial candidate c is not worth completing.
    # c, in this case, is Column and Row together.
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
    global NUM_FOUND
    NUM_FOUND += 1
    derp.append(P)
    print(P)


if __name__ == "__main__":
    backtrack([], *root())
    dupes = 0
    for index, i in enumerate(derp):
        for j in derp[index + 1::]:
            if i == j:
                dupes += 1
    print(dupes)
    print("Found solutions: {NUM}".format(NUM=NUM_FOUND))
