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

# def backtrack(P, C, R): push C R to stacks
#     if reject(P, C, R):
#         return
#     P.append(R)
#     if accept(P):
#         output(P)
#     s = first(C)
#     while s is not None:
#         backtrack(P, *s)
#         s = next(*s)
#     # P.pop()
#     # if more(R):
#     #     backtrack(P, *next(C, R))
#
# 1
# reject, candidate = pop from stack, goto 1
# accept output()
# push candidate to stack
# candidate = first(candidate)
# goto 1

# def go():
#     root = 1, N
#     candidate = root
#     stack = list()
#     structure = list()

#     while candidate is not None:
#         if reject(structure, *candidate):
#             print(1, stack, structure)
#             candidate = None if not stack else stack.pop()
#             continue
#         structure.append(candidate)
#         if accept(structure):
#             output(structure)
#             structure.pop()
#             print(2, stack)
#             candidate = None if not stack else stack.pop()
#             break
#         if candidate[1] > 1:
#             stack.append([candidate[0], candidate[1] - 1])
#         s = first(*candidate)
#         while s is not None:
#             stack.append(s)
#             s = next(*s)
#         print(3, stack)
#         candidate = None if not stack else stack.pop()
        # print(candidate, 3)


def go():
    root = first(0, N + 1)
    # candidate = first(*root)
    stack = list()
    structure = list([root])


    while root is not None:
        print("data:    ", structure, root)
        candidate = first(*root)
        while candidate is not None:
            if reject(structure, *candidate):
                candidate = next(*candidate)
                continue
            structure.append(candidate)
            if accept(structure):
                output(structure)
                return
            stack.append(root)
            root = candidate
            break
        else:

            # print("data:    ", structure)
            # print("stack:   ", stack)
            # print("root:    ", root)

            structure.pop()
            root = next(*stack.pop())
            structure[-1] = root
            # continue


        


    # while candidate is not None:
    #     if reject(structure, *candidate):
    #         print(1, stack, structure)
    #         candidate = None if not stack else stack.pop()
    #         continue
    #     structure.append(candidate)
    #     if accept(structure):
    #         output(structure)
    #         structure.pop()
    #         print(2, stack)
    #         candidate = None if not stack else stack.pop()
    #         break
    #     if candidate[1] > 1:
    #         stack.append([candidate[0], candidate[1] - 1])
    #     s = first(*candidate)
    #     while s is not None:
    #         stack.append(s)
    #         s = next(*s)
    #     print(3, stack)
    #     candidate = None if not stack else stack.pop()
    #     print(candidate, 3)

# while candidate is not None:
#     # stack.append(candidate)
#     if reject(structure, *candidate):
#         # stack.pop()
#         # structure.pop()
#         candidate = None if not structure else structure.pop()
#         continue
#     if accept(structure):
#         output(structure)
#     candidate = next(*candidate)
#     stack.append(candidate)



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
    column = 0
    for _, row in P:
        column += 1
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


go()
