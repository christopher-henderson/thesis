from __future__ import division
import time
from backtrack import backtrack
from itertools import product, permutations

RED = 'red'
BLUE = 'blue'
GREEN = 'green'
YELLOW = 'yellow'
IVORY = 'ivory'

WATER = 'water'
COFFEE = 'coffee'
TEA = 'tea'
MILK = 'milk'
ORANGE_JUICE = 'orange juice'

ENGLISH = 'English'
SPANISH = 'Spanish'
UKRAINIAN = 'Ukrainian'
NORWEGIAN = 'Norwegian'
JAPANESE = 'Japanese'

ZEBRA = 'zebra'
DOG = 'dog'
SNAILS = 'snails'
FOX = 'fox'
HORSE = 'horse'

OLD_GOLD = 'Old Gold'
KOOLS = 'Kools'
CHESTERFIELDS = 'Chesterfields'
LUCKY_STRIKE = 'Lucky Strike'
PARLIAMENTS = 'Parliaments'

COLORS = [RED, BLUE, GREEN, YELLOW, IVORY]
DRINKS = [WATER, COFFEE, TEA, MILK, ORANGE_JUICE]
NATIONALITIES = [ENGLISH, SPANISH, UKRAINIAN, NORWEGIAN, JAPANESE]
PET = [ZEBRA, DOG, SNAILS, FOX, HORSE]
CIGARETTE = [OLD_GOLD, KOOLS, CHESTERFIELDS, LUCKY_STRIKE, PARLIAMENTS]

ATTRS = [COLORS, DRINKS, NATIONALITIES, PET, CIGARETTE]

# configuration = [[None for _ in range(5)] for _ in range(5)]
# for index in range(5):
# 	configuration[index][0] = COLORS[index]
# 	configuration[index][1] = DRINKS[index]
# 	configuration[index][2] = NATIONALITIES[index]
# 	configuration[index][3] = PET[index]
# 	configuration[index][4] = CIGARETTE[index]


# print(configuration)
# conf = configuration
# conf[0][0], conf[1][0], conf[2][0], conf[3][0], conf[4][0] = conf[4][0], conf[0][0], conf[1][0], conf[2][0], conf[3][0]

color = 0
drink = 1
nationality = 2
pet = 3
cigarette = 4

class Next(object):

	def __init__(self):
		# self.conf = [list(_) for _ in zip(COLORS, DRINKS, NATIONALITIES, PET, CIGARETTE)]
		self.conf = [[None for _ in range(5)] for _ in range(5)]
		self.next = self.permutations(color)

	def __call__(self, *args, **kwargs):
		try:
			return next(self.next)
		except StopIteration:
			return None

	# def _next(self):
	# 	yield self.conf
	# 	for a in range(5):
	# 		for b in range(5):
	# 			for c in range(5):
	# 				for d in range(5):
	# 					for e in range(5):
	# 						for _ in range(5):
	# 							self.shift(cigarette)
	# 							yield self.conf
	# 						self.swap(e, (e + 1) % 5, cigarette)
	# 						yield self.conf
	# 					self.shift(pet)
	# 				self.shift(nationality)
	# 			self.shift(drink)
	# 		self.shift(color)


	def permutations(self, attr):
		if attr >= 5:
			yield self.conf
		else:
			for configuration in permutations(ATTRS[attr], 5):
				self.conf[0][attr], self.conf[1][attr], self.conf[2][attr], self.conf[3][attr], self.conf[4][attr] = configuration
				yield from self.permutations(attr + 1)
		# for combination of this attr:
		# 	for every combination of subattrs:
		# 		yield self.conf


	def shift(self, val):
		self.conf[0][val], self.conf[1][val], self.conf[2][val], self.conf[3][val], self.conf[4][val] = self.conf[4][val], self.conf[0][val], self.conf[1][val], self.conf[2][val], self.conf[3][val]

	def swap(self, a, b, attr):
		self.conf[a][attr], self.conf[b][attr] = self.conf[b][attr], self.conf[a][attr]



def accept(solution):
	return True

t = 0
def reject(_, candidate):
	if candidate[len(candidate) // 2][drink] != MILK:
		# print(-2)
		return True
	if candidate[0][nationality] != NORWEGIAN:
		# print(-1)
		return True
	if candidate[1][color] != BLUE:
		# print(0)
		return True

	for index, house in enumerate(candidate):
		if house[nationality] == ENGLISH and house[color] != RED:
			# print(1)
			return True
		if house[nationality] == SPANISH and house[pet] != DOG:
			# print(2)
			return True
		if house[drink] == COFFEE and house[color] != GREEN:
			# print(3)
			return True
		if house[nationality] == UKRAINIAN and house[drink] != TEA:
			# print(4)
			return True
		if house[color] == GREEN and index is 0 or candidate[index - 1][color] != IVORY:
			# print(5)
			return True
		if house[cigarette] == OLD_GOLD and house[pet] != SNAILS:
			# print(6)
			return True
		if house[cigarette] == KOOLS and house[color] != YELLOW:
			# print(7)
			return True
		if house[cigarette] == CHESTERFIELDS and not (
			False if index == len(candidate) - 1 else candidate[index + 1][pet] == FOX or
			False if index == 0 else candidate[index - 1][pet] == FOX):
			# print(8)
			return True
		if house[cigarette] == KOOLS and not (
			False if index == len(candidate) - 1 else candidate[index + 1][pet] == HORSE or
			False if index == 0 else candidate[index - 1][pet] == HORSE):
			# print(9)
			return True
		if house[cigarette] == LUCKY_STRIKE and house[drink] != ORANGE_JUICE:
			# print(10)
			return True
		if house[nationality] == JAPANESE and house[cigarette] != PARLIAMENTS:
			# print(11)
			return True
	return False

n = Next()

backtrack(1, n, n, reject, accept)