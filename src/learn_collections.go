package main

import "fmt"

func main() {
	ages := map[string]int{
		"Surya":   37,
		"Radhika": 36,
		"Shreya":  4,
	}

	for name, age := range ages {
		fmt.Println(name, "is of age", age)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("we are counting", i)
	}

	a := 1
	for a <= 10 {
		fmt.Println("we are counting a value", a)
		a++
	}

	myarr := [3]string{"Surya", "Radhika", "Shreya"}
	for index, value := range myarr {
		fmt.Println(value, "is at index", index)
	}
}
