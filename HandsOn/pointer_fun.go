package main

import "fmt"

// value param: won't affect caller
func addYearCopy(u User){
	u.Age++
}

// pointer param: will affect caller
func addYearInPlace(u *User){
	u.Age++
}

func demoPointerParams() {
	bob := User{Name: "Bob", Age:40}
	addYearCopy(bob)
	fmt.Println("after addYearCopy:" , bob.Age) // after addYearCopy: 40
	
	// kaori := User{Name: "kaori", Age:40}
	// addYearCopy(&kaori) // this fail to build because addYearCopy is no allow to set pointer param
	// fmt.Println("after addYearCopy kaori:" , kaori.Age) 

	addYearInPlace(&bob)
	fmt.Println("after addYearInPlace:" , bob.Age) // after addYearInPlace: 41
}