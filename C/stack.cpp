#include <iostream>
#include <vector>
using namespace std;

template <class T>
class Stack {
public:
	Stack(int);
	~Stack();
	int len() {return this->pointer;};
	void push(T obj);
	T pop();
private:
	T * stack;
	int max;
	int pointer;
};

template<class T>
Stack<T>::Stack(int max) {
	this->max = max;
	this->pointer = 0;
	this->stack = new T[max];
}

template<class T>
Stack<T>::~Stack() {
	delete this->stack;
}

template<class T>
void Stack<T>::push(T obj) {
	this->stack[this->pointer] = obj;
	this->pointer++;
}

template<class T>
T Stack<T>::pop() {
	this->pointer--;
	return this->stack[this->pointer];
}

int main(int argc, char const *argv[]) {
	std::vector<char> v;
}