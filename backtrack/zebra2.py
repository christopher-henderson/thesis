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

configuration = [[None for _ in range(5)] for _ in range(5)]
for index in range(5):
	configuration[index][0] = COLORS[index]
	configuration[index][1] = DRINKS[index]
	configuration[index][2] = NATIONALITIES[index]
	configuration[index][3] = PET[index]
	configuration[index][4] = CIGARETTE[index]

# print(configuration)
conf = configuration
# conf[0][0], conf[1][0], conf[2][0], conf[3][0], conf[4][0] = conf[4][0], conf[0][0], conf[1][0], conf[2][0], conf[3][0]

def _next():
	def inner():
		yield conf
		for _ in range(5):
			for _ in range(5):
				for _ in range(5):
					for _ in range(5):
						for i in range(5):
							# print(i)
							shift(4)
							yield conf
						shift(3)
					shift(2)
				shift(1)
			shift(0)
	gen = inner()
	def nexter(_):
		try:
			return next(gen)
		except:
			return None
	return nexter

nexter = _next()

def shift(val):
	conf[0][val], conf[1][val], conf[2][val], conf[3][val], conf[4][val] = conf[4][val], conf[0][val], conf[1][val], conf[2][val], conf[3][val]


color = 0
drink = 1
nationality = 2
pet = 3
cigarette = 4

def first(_):
	return nexter(_)

def accept(_):
	return True

def reject(_, candidate):
	if candidate[len(candidate) // 2][drink] != MILK:
		return True
	if candidate[0][nationality] != NORWEGIAN:
		return True
	if candidate[1][color] != BLUE:
		return True

	for index, house in enumerate(candidate):
		if house[nationality] == ENGLISH and house[color] != RED:
			print(1)
			return True
		if house[nationality] == SPANISH and house[pet] != DOG:
			print(2)
			return True
		if house[drink] == COFFEE and house[color] != GREEN:
			print(3)
			return True
		if house[nationality] == UKRAINIAN and house[drink] != TEA:
			print(4)
			return True
		if house[color] == GREEN and index is 0 or candidate[house - 1][color] != IVORY:
			print(5)
			return True
		if house[cigarette] == OLD_GOLD and house[pet] != SNAILS:
			print(6)
			return True
		if house[cigarette] == KOOLS and house[color] != YELLOW:
			print(7)
			return True
		if house[cigarette] == CHESTERFIELDS and not (
			False if index == len(candidate) - 1 else candidate[index + 1][pet] == FOX or
			False if index == 0 else candidate[index - 1][pet] == FOX):
			print(8)
			return True
		if house[cigarette] == KOOLS and not (
			False if index == len(candidate) - 1 else candidate[index + 1][pet] == HORSE or
			False if index == 0 else candidate[index - 1][pet] == HORSE):
			print(9)
			return True
		if house[cigarette] == LUCKY_STRIKE and house[drink] != ORANGE_JUICE:
			print(10)
			return True
		if house[nationality] == JAPANESE and house[cigarette] != PARLIAMENTS:
			print(11)
			return True
	return False

backtrack(1, first, nexter, reject, accept)