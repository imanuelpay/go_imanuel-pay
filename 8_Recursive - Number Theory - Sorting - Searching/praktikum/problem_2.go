package main

func main() {
	println(ficonacci(0))
	println(ficonacci(2))
	println(ficonacci(9))
	println(ficonacci(10))
	println(ficonacci(12))
}

func ficonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 || number == 2 {
		return 1
	} else {
		return ficonacci(number-1) + ficonacci(number-2)
	}
}
