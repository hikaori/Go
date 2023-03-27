package main

import (
	"fmt"
)


func main() {
    mybill := newBill("kaori's bill")
    fmt.Println(mybill.format())
}