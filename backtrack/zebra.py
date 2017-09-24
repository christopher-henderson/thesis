from __future__ import division
import time
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

	STATE = {
		RED: None,
		BLUE: None,
		GREEN: None,
		YELLOW: None,
		IVORY: None,

		WATER: None,
		COFFEE: None,
		TEA: None,
		MILK: None,
		ORANGE_JUICE: None,

		ENGLISH: None,
		SPANISH: None,
		UKRAINIAN: None,
		NORWEGIAN: None,
		JAPANESE: None,

		ZEBRA: None,
		DOG: None,
		SNAILS: None,
		FOX: None,
		HORSE: None,

		OLD_GOLD: None,
		KOOLS: None,
		CHESTERFIELDS: None,
		LUCKY_STRIKE: None,
		PARLIAMENTS: None,
	}

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

	@classmethod
	def reject(cls, solution, candidate):
		solution = solution + [candidate]
		return (
			cls.STATE[ENGLISH] is not None and cls.STATE[ENGLISH].color != RED or 
			cls.STATE[SPANISH] is not None and cls.STATE[SPANISH].pet != DOG or
			cls.STATE[COFFEE] is not None and cls.STATE[COFFEE].color != GREEN or
			cls.STATE[UKRAINIAN] is not None and cls.STATE[UKRAINIAN].drink != TEA or
			cls.STATE[GREEN] is not None and (cls.STATE[GREEN].left is None or cls.STATE[GREEN].left.color != IVORY) or
			cls.STATE[OLD_GOLD] is not None and cls.STATE[OLD_GOLD].pet != SNAILS or 
			cls.STATE[KOOLS] is not None and cls.STATE[KOOLS].color != YELLOW or
			len(solution) >= 3 and solution[2].drink != MILK or
			solution[0].nationality != NORWEGIAN or
			cls.STATE[CHESTERFIELDS] is not None and not (
				cls.STATE[CHESTERFIELDS].left is not None and cls.STATE[CHESTERFIELDS].left.pet == FOX or
				cls.STATE[CHESTERFIELDS].right is not None and cls.STATE[CHESTERFIELDS].right.pet == FOX
			) or
			cls.STATE[KOOLS] is not None and not (
				cls.STATE[KOOLS].left is not None and cls.STATE[KOOLS].left.pet == HORSE or
				cls.STATE[KOOLS].right is not None and cls.STATE[KOOLS].right.pet == HORSE
			) or
			cls.STATE[LUCKY_STRIKE] is not None and cls.STATE[LUCKY_STRIKE].drink != ORANGE_JUICE or
			cls.STATE[JAPANESE] is not None and cls.STATE[JAPANESE].cigarette != PARLIAMENTS or
			solution[0].right is not None and solution[0].right.color != BLUE
		)

	@classmethod
	def accept(cls, solution):
		# print(len(solution))
		global t
		t += 1
		return len(solution) == 5

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

		self.STATE[self.color] = self
		self.STATE[self.drink] = self
		self.STATE[self.nationality] = self
		self.STATE[self.pet] = self
		self.STATE[self.cigarette] = self

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
		self.STATE[self.color] = None
		self.STATE[self.drink] = None
		self.STATE[self.nationality] = None
		self.STATE[self.pet] = None
		self.STATE[self.cigarette] = None
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
start = time.time()
backtrack(first, House.first, House.next, House.reject, House.accept, single_solution=True)
print(time.time() - start)
print(t)