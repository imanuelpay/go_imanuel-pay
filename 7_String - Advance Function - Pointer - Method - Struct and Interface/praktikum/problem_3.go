package main

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	println(a, b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
