#include <ctime>
#include <iostream>
#include <vector>

using namespace std;

int N = 8;
int wins = 0;

class Queen {
public:
	int column;
	int row;
	Queen(int r, int c) {
		this->row = r;
		this->column = c;
		this->current = nullptr;
		this->prev = nullptr;
	}
	Queen * next() {
		if (this->current == nullptr) {
			this->current = new Queen(1, this->column + 1);
		}
		if (this->prev != nullptr) {
			delete this->prev;
		}
		if (this->current->row > N || this->current->column > N) {
			delete this->current;
			return nullptr;
		}
		this->prev = this->current;
		this->current = new Queen(this->prev->row + 1, this->prev->column);
		return this->prev;
	}
private:
	Queen * current;
	Queen * prev;
};

bool reject(Queen * queen, std::vector<Queen*> solution) {
	int column = queen->column;
	int row = queen->row;
	for (int i = 0; i < solution.size(); i++) {
		int r = solution.at(i)->row;
		int c = solution.at(i)->column;
		if ((row == r) || (column == c) || (row + column == r + c) || (row - column == r - c)) {
			return true;
		}
	}
	return false;
}

bool accept(std::vector<Queen*> solution) {
	return solution.size() == N;
}

void output(std::vector<Queen*> solution) {
	for (int i = 0; i < solution.size(); i++) {
		cout << "(" << solution.at(i)->column << ", " << solution.at(i)->row << "), ";
	}
	cout << endl;
}

void backtrack(Queen * FCG, int size) {
	Queen * root = new Queen(FCG->row, FCG->column);
	Queen * candidate;
	std::vector<Queen *> solution;
	std::vector<Queen *> stack;
	while (1) {
		candidate = root->next();
		if (candidate == nullptr) {
			if (stack.empty()) {
				cout << "dun" << endl;
				break;
			}
			// printf("%s\n", "buleteing root");
			// delete root;
			solution.pop_back();
			root = stack.back();
			stack.pop_back();
			continue;
		}
		if (reject(candidate, solution)) {
			continue;
		}
		solution.push_back(candidate);
		if (accept(solution)) {
			// output(solution);
			wins++;
			solution.pop_back();
			continue;
		}
		stack.push_back(root);
		root = candidate;
	}
}

int main(int argc, char const *argv[]) {
	std::clock_t start;
    start = std::clock();
	Queen FCG(0, 0);
	backtrack(&FCG, N);
	cout << wins << endl;
	std::cout << "Time: " << (std::clock() - start) / (double)(CLOCKS_PER_SEC) << " s" << std::endl;
	return 0;
}