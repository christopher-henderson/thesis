# import re
# r = re.compile("(?P<SPACES>\s*)(?:return)\s+(?P<EXP>(.*\n?)*)")

# b = 'for i in range(10):\n\tif i % 2 == 0:\n\t\treturn any(\n\t\t\t1, 2, 3\n\t\t\t)\n\telse:\n\t\treturn 1'

# def exp(e):
#     print(e)
#     p = 0
#     b = 0
#     c = 0
#     for i, char in enumerate(e):
#         print(char == '\n', char)
#         if char == '\n' and p is 0 and b is 0 and c is 0:
#             print("nugz")
#             return e[:i], e[i:]
#         if char == '(':
#             p += 1
#         elif char == ')':
#             p -= 1
#         elif char == '[':
#             b += 1
#         elif char == ']':
#             b -= 1
#         elif char == '{':
#             c += 1
#         elif char == '}':
#             c -= 1
#     return e, ""

# def repl(match):
#     e, rest = exp(match.group("EXP"))
#     return "{S}__func__return__ = {EXP} ; break{REST}".format(S=match.group("SPACES"), EXP=e, REST=rest)
#     # return match.group(0) + "__fun__return__ = "

# t = r.sub(repl, b)
# print(t)

def a(l):
    # le sigh https://stackoverflow.com/a/1278351
    # gotta do this is Python 2
    exec ("")
    locals().update(l)
    this = "is a thing"
    print(var)
    return True, locals()

def b():
    var = 5
    result, l = a(locals())
    exec("")
    locals().update(l)
    print(this)

print(b())