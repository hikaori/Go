package main

import ("fmt")

// like a TS interface shape, but concrete
type User struct {
	Name string
	Age int
}

// VALUE receiver: gets a copy (does NOT mutate caller)
func (u User) Anniversary(){
	u.Age++ // modifies the copy only
}

// POINTER receiver: works on the original (mutates caller)
func (u *User) Birthday(){
	u.Age++
}


func (u *User) Rename(newName string){
	u.Name = newName
}

func main() {
	alice := User{Name: "Alice", Age:29}

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
	fmt.Println("after `ptr.Birthday` Print ptr:", ptr) // Age 31

	demoPointerParams()
}