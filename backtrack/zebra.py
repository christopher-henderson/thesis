from __future__ import division
from backtrack import backtrack
from itertools import product

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
t = 0

class House(object):

	COLORS = {RED, BLUE, GREEN, YELLOW, IVORY}
	DRINKS = {WATER, COFFEE, TEA, MILK, ORANGE_JUICE}
	NATIONALITIES = {ENGLISH, SPANISH, UKRAINIAN, NORWEGIAN, JAPANESE}
	PET = {ZEBRA, DOG, SNAILS, FOX, HORSE}
	CIGARETTE = {OLD_GOLD, KOOLS, CHESTERFIELDS, LUCKY_STRIKE, PARLIAMENTS}

	@classmethod
	def init_houses(cls, num):
		first = previous = House()
		for _ in range(num - 1):
			current = House()
			previous.right = current
			current.left = previous
			previous = current
		first.set_configurations_options()
		return first.next()

	def __init__(self):
		self.color = None
		self.drink = None
		self.nationality = None
		self.pet = None
		self.cigarette = None

		self.left = None
		self.right = None

	@staticmethod
	def reject(s, candidate):
		solution = s + [candidate]
		reject =  (
			any(house.nationality == ENGLISH and house.color != RED for house in solution) or
			any(house.nationality == SPANISH and house.pet != DOG for house in solution) or
			any(house.drink == COFFEE and house.color != GREEN for house in solution) or
			any(house.nationality == UKRAINIAN and house.drink != TEA for house in solution) or
			any(house.color == GREEN and (house.left is None or house.left.color != IVORY) for house in solution) or
			any(house.cigarette == OLD_GOLD and house.pet != SNAILS for house in solution) or
			any(house.cigarette == KOOLS and house.color != YELLOW for house in solution) or
			len(solution) >= 3 and solution[len(solution) // 2].drink != MILK or
			solution[0].nationality != NORWEGIAN or
			any(house.cigarette == CHESTERFIELDS and (
				(house.left is not None and house.left.pet != FOX) and
				(house.right is not None and house.right.pet != FOX)
				) for house in solution
			) or
			any(house.cigarette == KOOLS and (
				(house.left is not None and house.left.pet != HORSE) and
				(house.right is not None and house.right.pet != HORSE)
				) for house in solution
			) or
			any(house.cigarette == LUCKY_STRIKE and house.drink != ORANGE_JUICE for house in solution) or
			any(house.nationality == JAPANESE and house.cigarette != PARLIAMENTS for house in solution) or
			any(house.nationality == NORWEGIAN and (
				(house.left is not None and house.left.color != BLUE) and
				(house.right is not None and house.right.color != BLUE)
				) for house in solution
			)
		)
		return reject

	@classmethod
	def accept(cls, solution):
		print(len(solution))
		global t
		t += 1
		return len(solution) == 5
		# return (
		# 	len(solution) == 5 and
		# 	any(house.nationality == ENGLISH and house.color == RED for house in solution) and
		# 	any(house.nationality == SPANISH and house.pet == DOG for house in solution) and
		# 	any(house.drink == COFFEE and house.color == GREEN for house in solution) and
		# 	any(house.nationality == UKRAINIAN and house.drink == TEA for house in solution) and
		# 	any(house.color == GREEN and house.left is not None and house.left.color == IVORY for house in solution) and
		# 	any(hocuse.cigarette == OLD_GOLD and house.pet == SNAILS for house in solution) and
		# 	any(house.cigarette == KOOLS and house.color == YELLOW for house in solution) and
		# 	solution[len(solution) // 2].drink == MILK and
		# 	solution[0].nationality == NORWEGIAN and
		# 	any(house.cigarette == CHESTERFIELDS and (
		# 		house.left is not None and house.left.pet == FOX or
		# 		house.right is not None and house.right.pet == FOX
		# 		) for house in solution
		# 	) and
		# 	any(house.cigarette == KOOLS and (
		# 		house.left is not None and house.left.pet == HORSE or
		# 		house.right is not None and house.right.pet == HORSE
		# 		) for house in solution
		# 	) and
		# 	any(house.cigarette == LUCKY_STRIKE and house.drink == ORANGE_JUICE for house in solution) and
		# 	any(house.nationality == JAPANESE and house.cigarette == PARLIAMENTS for house in solution) and
		# 	any(house.nationality == NORWEGIAN and (
		# 		house.left is not None and house.left.color == BLUE or
		# 		house.right is not None and house.right.color == BLUE
		# 		) for house in solution
		# 	)
		# )

	def first(self):
		if self.right is None:
			return None
		self.right.set_configurations_options()
		return self.right.next()

	def next(self):
		self.replace_options()
		try:
			configuration = next(self.options)
		except StopIteration:
			return None		
		self.color = configuration[0]
		self.drink = configuration[1]
		self.nationality = configuration[2]
		self.pet = configuration[3]
		self.cigarette = configuration[4]

		self.COLORS.remove(self.color)
		self.DRINKS.remove(self.drink)
		self.NATIONALITIES.remove(self.nationality)
		self.PET.remove(self.pet)
		self.CIGARETTE.remove(self.cigarette)
		return self

	@classmethod
	def remove(cls, solution):
		house = solution.pop()
		house.replace_options()

	def replace_options(self):
		if self.color is not None:
			self.COLORS.add(self.color)
			self.DRINKS.add(self.drink)
			self.NATIONALITIES.add(self.nationality)
			self.PET.add(self.pet)
			self.CIGARETTE.add(self.cigarette)
		self.color = None
		self.drink = None
		self.nationality = None
		self.pet = None
		self.cigarette = None

	def set_configurations_options(self):
		self.options = product(self.COLORS, self.DRINKS, self.NATIONALITIES, self.PET, self.CIGARETTE)

	def __str__(self):
		return str(tuple([self.color, self.drink, self.nationality, self.pet, self.cigarette]))

	def __repr__(self):
		return str(self)


first = House.init_houses(5)
first.next()
backtrack(first, House.first, House.next, House.reject, House.accept)
print(t)