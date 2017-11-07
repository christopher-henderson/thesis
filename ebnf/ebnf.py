from backtrack import backtrack


class SL(object):

	def __init__(self, tokens):
		self.tokens = tokens
		self.result = []
		self.succeded = False

	@staticmethod
	def __reject__(_, self):
		if len(self.tokens) is 0:
			return False
		s = S(self.tokens)
		backtrack(s)
		if not s.succeded:
			return True
		sl = SL(self.tokens)
		backtrack(sl)
		self.result.extend(sl.result)
		return sl.succeded
	
	@staticmethod
	def __accept__(solution):
		return True

	def __child__(self):
		return self

	def __sibling__(self):
		return None

class S(object):

	def __init__(self, tokens):
		self.tokens = tokens
		self.succeded = False
		self.consumed = []
		self.result = []

	@staticmethod
	def __reject__(_, self):
		for e in (t(self.tokens) for t in Exp.EXP):
			backtrack(e)
			if not e.succeded or len(self.tokens) is 0 or self.tokens[0] != ';':
				self.tokens.extend(e.consumed)
				continue
			self.consumed.extend(e.consumed)
			self.consumed.append(self.tokens.pop(0))
			self.result.extend(e.result)
			self.succeded = True
			return False
		return True

	def __accept__(self):
		return True

	def __output__(self):
		pass

	def __child__(self):
		return self

	def __sibling__(self):
		return None

class Int(object):

	def __init__(self, tokens):
		self.tokens = tokens
		self.succeded = False
		self.consumed = []
		self.result = [Int]

	@staticmethod
	def __reject__(_, self):
		try:
			int(self.tokens[0])
			self.consumed = self.tokens.pop(0)
			self.succeded = True
			return False
		except:
			return True

	def __accept__(self):
		return True

	def __child__(self):
		return self

	def __sibling__(self):
		return None

class Add(object):

	def __init__(self, tokens):
		self.tokens = tokens
		self.succeded = False
		self.consumed = []
		self.result = [Add]

	@staticmethod
	def __reject__(_, self):
		if len(tokens) < 3:
			return True
		left = Int(self.tokens)
		backtrack(left)
		if not left.succeded:
			return True
		self.consumed.extend(left.consumed)
		if self.tokens[0] != '+':
			return True
		print("IT's WORKING")
		self.consumed.append(self.tokens.pop(0))
		right = Int(self.tokens)
		backtrack(right)
		if not right.succeded:
			return True
		self.consumed.extend(right.consumed)
		self.succeded = True
		return False

	def __accept__(self):
		return True

	def __child__(self):
		return self

	def __sibling__(self):
		return None

class Exp(object):

	EXP = [Int, Add]

	def __init__(self, tokens):
		self.tokens = tokens
		self.consumed = []
		self.result = []
		self.succeded = False

	@staticmethod
	def __reject__(_, self):
		for exp in self.EXP:
			e = exp(self.tokens)
			backtrack(e)
			if e.succeded:
				if len(self.tokens) is not 0 and self.tokens[0] == ';':
					self.consumed.extend(e.consumed)
					self.consumed.append(self.tokens.pop(0))
					self.result.extend(e.result)
					self.succeded = True
					return False
				self.tokens = i.consumed + self.tokens
		return True

	def __accept__(self):
		return True

	def __child__(self):
		return self

	def __sibling__(self):
		return None

tokens = ['1', '+', '1', ';']
sl = SL(tokens)
backtrack(sl)
print(sl.result, sl.tokens)