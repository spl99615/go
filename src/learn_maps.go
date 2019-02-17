package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Radhika": "08/30/1981",
		"Surya":   "01/10/1982",
		"Shreya":  "09/20/2014",
	}

	fmt.Println(birthdays)

	delete(birthdays, "Surya")

	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Surya"] = 37
	ages["Radhika"] = 36
	ages["Shreya"] = 4

	fmt.Println(ages)
}
