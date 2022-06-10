package main

import (
	"fmt"
	"reflect"
)

// a pointer value is the addres of a variable
// A pointer is thus the location at which a vaule is stored

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	fmt.Println(reflect.TypeOf(p))
	*p = 2
	fmt.Println(x)
	fmt.Println(*p)

	// 	// the zero value of a pointer is nil.
	// 	// pointer is also comparable, two pointers are equal if they point to the same variable or both are nil

	fmt.Println(f(&x))
	fmt.Println(f(&x))
	fmt.Println(x)
}

// // it's perfectly safe for a function to return the address of a local variable
func f(n *int) *int {
	*n = 5
	v := 5
	return &v
}

// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

// func main() {
// 	flag.Parse()
// 	fmt.Print(strings.Join(flag.Args(), *sep))
// 	if !*n {
// 		fmt.Println()
// 	}
// }
