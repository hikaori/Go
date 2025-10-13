package main

import (
	"fmt"
)

// like a TS interface shape, but concrete
type User struct {
	Name string
	Age  int
}

// VALUE receiver: gets a copy (does NOT mutate caller)
func (u User) Anniversary() {
	u.Age++ // modifies the copy only
}

// POINTER receiver: works on the original (mutates caller)
func (u *User) Birthday() {
	if u == nil {
		return
	}
	u.Age++
}

func (u *User) Rename(newName string) {
	if u == nil {
		return
	}
	u.Name = newName
}

func main() {
	alice := User{Name: "Alice", Age: 29}

	// value receiver -> no change
	alice.Anniversary()
	fmt.Println("after Anniversary (value):", alice) // Age still 29

	alice.Birthday()
	fmt.Println("after Birthday (value):", alice) // Age still 30

	alice.Rename("Alicia")
	fmt.Println("after Rename (pointer):", alice) // Name "Alicia"

	// explicit pointer usage
	ptr := &alice
	ptr.Birthday()
	fmt.Println("after `ptr.Birthday` Print alice:", alice) // Age 31
	fmt.Println("after `ptr.Birthday` Print ptr:", ptr)     // Age 31

	demoPointerParams()

	demoJSON()

	mike := User{Name: "Mike", Age: 20}
	fmt.Println(mike.Label())

	pBen := NewUser("Ben", 10)

	pBen.Birthday()
	pBen.Rename("Benji")
	// fmt.Sprintf("after Ben renamed: %s, age(%d)", pBen.Name, ptr.Age)
	fmt.Println("Ben age:", pBen.Age)
	fmt.Println("after Ben renamed:", pBen.Name)

}
