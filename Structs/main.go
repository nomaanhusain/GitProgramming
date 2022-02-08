package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	//nomaan := person{firstName: "Nomaan", lastName: "Husain", age: 22} //one way to do this
	var nomaan person
	nomaan.firstName = "Nomaan"
	nomaan.lastName = "Husain"
	nomaan.age = 22
	jim := person{
		firstName: "Jim",
		lastName:  "Pearson",
		age:       28,
		contact: contactInfo{
			email:   "jim@xyz.com",
			pinCode: 24568,
		},
	}
	fmt.Println(nomaan)
	// %+v is used to print all fields and data of a struct
	//fmt.Printf("%+v", nomaan)
	fmt.Println()
	fmt.Printf("%+v", jim)

	//Go is a pass by value language (except for refrence data types eg. slice,map,channel,pointers,function)
	//and hence the original struct does not get updated.
	//We have to use pointer to pass as refrence
	//& gives the memory adress of jim
	jimPointer := &jim
	jimPointer.updateFirstName("Nomaaannn")
	fmt.Println()
	fmt.Printf("%+v", jim)
	//Or, Both are valid ways, this uses automatic type conversion
	nomaan.updateFirstName("Jimmy")
	fmt.Println()
	fmt.Printf("%+v", nomaan)
}

// *{type} it means that we are working with a pointer to a person
func (pointerToPerson *person) updateFirstName(newFName string) {
	// *{vaiable_name} returns the value stored in the memory address of the pointer passed
	(*pointerToPerson).firstName = newFName
}
