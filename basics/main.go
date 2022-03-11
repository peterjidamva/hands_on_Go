package main

import (
	"fmt"
	"math"
	rvsStr "mypackgs"
)

//structure
type Student struct {
	Name  string
	regNo string
	age   int
	sex   string

	// CalculateReg();

}

func CalculateReg(s Student) int {

	value := len(s.Name)

	return value
}

//using a package
func tellingFacts(word string) string {
	return word
}

func main() {
	//structure data filling

	student1 := Student{Name:"PETER Jidamva" ,regNo:"2000-04-1234",age:21,sex:"M" }
	
	//variable declarations
	var uname = "DRIFT"
	var time = 12
	const pi = 3.14
	var array [3]string

	//arrays
	array[0] = "drift"
	array[1] = "wood"
	array[2] = "jidamva"

	number := []int{1, 3, 4, 5, 6, 7, 8, 9, 10}

	//also we can declare variable as

	value := 6.55
	
	fmt.Println(CalculateReg(student1))

	fmt.Println(array)
	fmt.Println(number[1:4])
	fmt.Println(tellingFacts("drift"))
	fmt.Println(rvsStr.Reverse("Hello"))
	fmt.Println(uname, time, pi, value)
	fmt.Println("Hello world")
	fmt.Println(math.Floor(value))
	fmt.Println(math.Sqrt(value))

	x := 3
	y := 8
	//control statements
	if x <= y {
		fmt.Printf("%d is less than %d", x, y)

	} else {
		fmt.Printf("%d is less than %d", y, x)

	}

	switch x {

	case 3:
		fmt.Println("it is 3")
	case 5:
		fmt.Println("it is 5")
	default:
		fmt.Println("we dont know it")
	}

	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//maps

	email := make(map[string]string)
	email["jidamva"] = "jidamva@gmail.com"
	email["drift"] = "drift@gmail.com"
	email["name"] = "name@gmail.com"

	fmt.Println(email)
	fmt.Println(len(email))
	delete(email, "drift")
	fmt.Println(email["name"])

	//declaring and initializing a map
	school := map[int]string{
		1: "money",
		2: "fame",
		3: "Power"}

	fmt.Println(school)

}
