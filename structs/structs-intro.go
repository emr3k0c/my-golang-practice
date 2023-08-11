package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	// Also I can reassign this by
	alex.firstName = "Emre"
	alex.lastName = "Koc"
	jim := person{
		firstName: "jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "info@gmail.com",
			zipCode: 7600,
		},
	}
	fmt.Println(jim)

}
