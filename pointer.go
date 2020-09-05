package main

func add5Value(count int) {
	count += 5
}

func add5Point(count *int) {
	*count += 5
}

func main() {
	var count int
	add5Value(count)
	add5Point(&count)
}

/*
main 0
main 0xc0000100b0

add5Value 0xc0000100b8
add5Value before 5

main 0

add5Point 0xc000006030
add5Point 5

main 5
main 0xc0000100b0
*/
