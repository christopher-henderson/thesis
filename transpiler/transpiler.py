# coding=utf8
# the above tag defines encoding for this document and is for Python 2.x compatibility

import re

# https://regex101.com/r/ObsKny/1/

regex = r"backtrack\s+on\s+(.*):\n\s+reject\(\s*(.*)\s*,\s*(.*)\s*\):\n((?:.|\n)*)accept\(\s*(.*)\s*\):\n((?:.|\n)*)children\(\s*(.*)\s*\):\n((?:.|\n)*)\n"

test_str = ("backtrack on (1, 1):\n"
	"	reject(solution, candidate):\n"
	"        column, row = candidate\n"
	"        return any(\n"
	"            row == r or\n"
	"            column == c or\n"
	"            row + column == r + c or \n"
	"            row - column == r - c for r, c in solution)\n"
	"    accept(solution):\n"
	"        return len(solution) == N\n"
	"    children(parent):\n"
	"        column, _ = parent\n"
	"        for row in range(1, N + 1):\n"
	"            yield tuple([column, row])\n"
	"print(\"LOL\")")

matches = re.finditer(regex, test_str)

for matchNum, match in enumerate(matches):
    matchNum = matchNum + 1
    
    print ("Match {matchNum} was found at {start}-{end}: {match}".format(matchNum = matchNum, start = match.start(), end = match.end(), match = match.group()))
    
    for groupNum in range(0, len(match.groups())):
        groupNum = groupNum + 1
        
        print ("Group {groupNum} found at {start}-{end}: {group}".format(groupNum = groupNum, start = match.start(groupNum), end = match.end(groupNum), group = match.group(groupNum)))

# Note: for Python 2.7 compatibility, use ur"" to prefix the regex and u"" to prefix the test string and substitution.
