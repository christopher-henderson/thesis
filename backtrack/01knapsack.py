from backtrack import backtrack


class KnapSack(object):


	WEIGHT = 50
	ITEMS = ((10, 60), (20, 100), (30, 120))

	MEMO = {}

	def __init__(self, item_number):
		self.item_number = item_number	

	def first(self):
		return KnapSack(0)

	def next(self):
		if self.item_number >= len(self.ITEMS) - 1:
			return None
		return KnapSack(self.item_number + 1)

	@classmethod
	def reject(cls, P, candidate):
		weight = cls.total_weight(P)
		if weight + candidate.weight > cls.WEIGHT or weight in cls.MEMO:
			# It doesn't fill up the whole bag, but it IS a solution.
			cls.output(P)
			return True
		return False

	@classmethod
	def accept(cls, P):
		return cls.total_weight(P) == cls.WEIGHT

	@staticmethod
	def add(P, candidate):
		P.append(tuple([candidate.weight, candidate.value]))

	@staticmethod
	def remove(P):
		P.pop()

	@classmethod
	def output(cls, P):
		weight = cls.total_weight(P)
		if weight in cls.MEMO:
			if cls.total_value(P) > cls.total_value(cls.MEMO[weight]):
				cls.MEMO[weight] = tuple(P)
		else:
			cls.MEMO[weight] = tuple(P)

	# Three helpers
	@staticmethod
	def total_weight(P):
		return sum(item[0] for item in P)

	@staticmethod
	def total_value(P):
		return sum(item[1] for item in P)		

	@property
	def weight(self):
		return self.ITEMS[self.item_number][0]

	@property
	def value(self):
		return self.ITEMS[self.item_number][1]


if __name__ == '__main__':
	backtrack(KnapSack(0),
		KnapSack.first,
		KnapSack.next,
		KnapSack.reject,
		KnapSack.accept,
		KnapSack.add,
		KnapSack.remove,
		KnapSack.output)
	print("Max total is: {}\nUsing the solution: {}".format(KnapSack.total_value(KnapSack.MEMO[KnapSack.WEIGHT]), KnapSack.MEMO[KnapSack.WEIGHT]))