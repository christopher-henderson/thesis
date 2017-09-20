# Don't...pay too much attentinon to this, because I didn't spend a whole lot of time on it.
# This is a terrible way to solve the Fibonacci sequence, but it does add evidence towards the idea that
# if you can represent your problem as a graph, then this function can drive the solution.

from backtrack import backtrack

N = 6
mod = N % 2

memo = {
	0: 1,
	1: 1
}

def first(candidate):
	return None if candidate is 1 or candidate is 0 else candidate - 1


def next(candidate):
	if candidate is N:
		return None
	if candidate % 2 is mod:
		return None
	return candidate - 1


def reject(P, candidate):
	return False


def accept(P):
	if len(P) is not 0 and P[-1] in memo:
		return True
	return False


def add(P, candidate):
	P.append(candidate)


def remove(P):
	P.pop()


def output(P):
	for val in P[-2::-1]:
		memo[val] = memo[val - 1] + memo[val - 2]
	print(memo[N])


if __name__ == '__main__':
	backtrack(N, first, next, reject, accept, add, remove, output)