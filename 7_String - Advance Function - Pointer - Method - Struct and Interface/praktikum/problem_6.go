package main

import "fmt"

var cipher = make(map[int]string)

func main() {
	var menu int
	var s = Student{}
	var c Chiper = &s

	for i := 0; i < 13; i++ {
		cipher[65+i] = string(byte(90 - i))
		cipher[90-i] = string(byte(65 + i))
		cipher[97+i] = string(byte(122 - i))
		cipher[122-i] = string(byte(97 + i))
	}

	print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		print("\nInput Student's Name: ")
		fmt.Scan(&s.name)
		print("\nEncode of Student's Name " + s.name + " is : " + c.Encode())
	} else if menu == 2 {
		print("\nInput Student's Decode Name: ")
		fmt.Scan(&s.nameEncode)
		print("\nDecode of Student's Name " + s.nameEncode + " is : " + c.Decode())
	} else {
		println("Wrong input name menu!")
	}
}

type Student struct {
	name       string
	nameEncode string
	score      string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	var nameEncode = ""
	name := (*s).name

	for _, value := range name {
		char, isExists := cipher[int(value)]
		if isExists {
			nameEncode += char
		}
	}

	return nameEncode
}

func (s *Student) Decode() string {
	var nameDecode = ""
	nameEncode := (*s).nameEncode

	for _, value := range nameEncode {
		char, isExists := cipher[int(value)]
		if isExists {
			nameDecode += char
		}
	}

	return nameDecode
}
