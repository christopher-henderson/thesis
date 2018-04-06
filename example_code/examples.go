package main

func first(q Queen) {
	return Queen{q.Column + 1, q.Row}
}

func next(q Queen) {
	return Queen{q.Column, q.Row + 1}
}

func main() {

}
