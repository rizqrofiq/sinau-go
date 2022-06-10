package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// Each declaration has the general form: var name [type] = [expression]
	// If the type is omitted, it is determined by the initializer expression, which:
	// 0 for numeric types, true for booleans, "" from strings, and nil for pointers, maps, and channels.

	var name string
	fmt.Println(name) // it will print empty string, because name is not initialized yet & the zero value of string is ""

	// it is possible to declare and initialize a set of variables in a single declaration.
	var i, j, k int
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(k))

	// ommiting the type allows declaration of multiple variables of different types
	var b, f, s = true, 2.3, "four"
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(f), reflect.TypeOf(s))

	// short variable declaration
	// may be used to declare and initialize local variables
	// it takes the form: name := expression
	animal := "cat"
	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Println(animal, freq, t)

	// as with var declarations, multiple variables can be declared and initialized in a single statement
	// if some of them were already declared in the same lexical block, then the short variable
	// declaration acts like an assignment to those variables
	i, w := 1, "dua"
	fmt.Println(i, w)

	// a short variable declaration must be declare at least one new variable, so the the code below will not compile
	// i, w := 2, "tiga"

	// another way to create new variables is to use the built-in function new(T).
	// This returns a pointer to a newly allocated zero value of type T.

	n := new(int)
	fmt.Println(reflect.TypeOf(n), *n) // *int 0
	*n = 5
	fmt.Println(*n) // 5
}
