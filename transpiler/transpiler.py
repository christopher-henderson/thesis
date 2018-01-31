# coding=utf8
# the above tag defines encoding for this document and is for Python 2.x compatibility

import re

# With the handy help of:
# 	https://regex101.com/r/ObsKny/1/

regex = re.compile(r"backtrack\s+on\s+(?P<ROOT>.*):\n\s+reject\(\s*(?P<REJECT_PARAM_SOLUTION>.*)\s*,\s*(?P<REJECT_PARAM_CANDIDATE>.*)\s*\):\n(?P<REJECT_BLOCK>(?:.|\n)*)accept\(\s*(?P<ACCEPT_PARAM_SOLUTION>.*)\s*\):\n(?P<ACCEPT_BLOCK>(?:.|\n)*)children\(\s*(?P<CHILDREN_PARAM_PARENT>.*)\s*\):\n(?P<CHILDREN_BLOCK>(?:.|\n)*)\n?")

class BacktrackFunction(object):

	INDENT = '\t' # deal with it, nerd

	def __init__(self, name, body, *params):
		self.name = name
		self.body = self.outdent_body(body)
		self.params = params

	@staticmethod
	def outdent_body(body):
		if len(body) is 0:
			return body
		lines = body.split("\n")
		indent = BacktrackFunction.get_indent(lines[0])
		for i, l in enumerate(lines):
			lines[i] = '{I}{L}'.format(I=BacktrackFunction.INDENT, L=l.replace(indent, "", 1))
		return "\n".join(lines)


	@staticmethod
	def get_indent(first_line):
		if len(first_line) is 0:
			return 0
		i = 0
		whitespace = first_line[0]
		for c in first_line:
			if c == whitespace:
				i += 1
			else:
				break
		return "".join(whitespace for _ in range(i))

	def __str__(self):
		return '''\
def {N}({PARAMS}, **globals, **locals):
{BODY}
'''.format(N=self.name, PARAMS=", ".join(self.params), BODY=self.body)


class BacktrackModule(object):

	def __init__(self, root, start, end):
		self.start = start
		self.end = end
		self.root = None
		self.functions = list()

	def add(self, *functions):
		self.functions.extend(functions)


_start_bactrack = re.compile(r"\s*backtrack\s+on.*")
def start_backtrack(l):
	return bool(_start_bactrack.match(l))

_whitespace = re.compile(r"\s")
def get_indent(l):
		i = 0
		for c in l:
			if bool(_whitespace.match(c)):
				i += 1
			else :
				return i
		return i


def build_module(backtrack, start, file):
	print('-------------BUILDING BACKTRACK MODULE-----------')
	indent = get_indent(backtrack)
	statement = [backtrack]
	final = ""
	consumed = 0
	for l in file:
		consumed += 1
		if get_indent(l) <= indent:
			final = l
			break
		else:
			statement.append(l)
	end = start + consumed - 1
	statement = "".join(statement)
	match = [_ for _ in regex.finditer(statement)][0]
	module = BacktrackModule(match.group("ROOT"), start, end)
	module.add(*[
		BacktrackFunction("reject", match.group("REJECT_BLOCK"), match.group("REJECT_PARAM_SOLUTION"), match.group("REJECT_PARAM_CANDIDATE")),
		BacktrackFunction("accept", match.group("ACCEPT_BLOCK"), match.group("ACCEPT_PARAM_SOLUTION")),
		BacktrackFunction("children", match.group("CHILDREN_BLOCK"), match.group("CHILDREN_PARAM_PARENT"))
	])
	for function in module.functions:
		print(function)
	return final, consumed


def main():
	m = list()
	with open("nqueens.py", "r") as file:
		lineNum = 0
		for line in file:
			lineNum += 1
			if not start_backtrack(line):
				m.append(line)
				continue
			else:
				left_over, consumed = build_module(line, lineNum, file)
				lineNum += consumed
				while start_backtrack(left_over):
					left_over, consumed = build_module(backtrack, start, file)
					lineNum += consumed
				m.append(left_over)


if __name__ == "__main__":
	main()