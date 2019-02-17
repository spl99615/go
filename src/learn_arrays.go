package main

import "fmt"

func main() {
	//var myarr []string
	myarr := make([]string, 1)
	myarr[0] = "Surya Lanka"
	myarr = append(myarr, "Radhika")
	myarr = append(myarr, "Shreya Lanka")

	fmt.Println(myarr)
}
