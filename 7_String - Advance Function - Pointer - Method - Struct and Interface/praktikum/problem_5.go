package main

import "fmt"

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input ", i+1, " Strudent's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	println("\n\nAverage Score Students is :", int(a.Avarage()))
	scoreMin, nameMin := a.Min()
	println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
	scoreMax, nameMax := a.Max()
	println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

}

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var count int
	for _, score := range s.score {
		count += score
	}

	return float64(count / len(s.score))
}

func (s Student) Min() (min int, name string) {
	var students = map[int]string{}

	min = s.score[0]
	for i, name := range s.name {
		students[s.score[i]] = name
		if s.score[i] < min {
			min = s.score[i]
		}
	}

	a, x := students[min]
	if x {
		name = a
	}

	return
}

func (s Student) Max() (max int, name string) {
	var students = map[int]string{}

	max = s.score[0]
	for i, name := range s.name {
		students[s.score[i]] = name
		if s.score[i] > max {
			max = s.score[i]
		}
	}

	a, x := students[max]
	if x {
		name = a
	}

	return
}
