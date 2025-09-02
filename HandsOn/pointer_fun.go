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
	fmt.Println("after addYearCopy:" , bob.Age)

	addYearInPlace(&bob)
	fmt.Println("after addYearInPlace:" , bob.Age)
}