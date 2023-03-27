package main

import "fmt"

type bill struct{
    name string
    items map[string] float64
    tip float64
}

func newBill(name string) bill {
	b := bill {
		name:name,
		items:map[string]float64{"pie":5.99,"cake":3.99},
		tip:0,
	}

	return b
}

// receiver function
func (b *bill)format()string{
	fs := "bill breakdown: \n"
	var total float64 = 0

	// list items
	for k,v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":" , v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)

	return fs
}