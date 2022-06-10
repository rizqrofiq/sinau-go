// Each file that stores a Go program is always begin with a package decalaration that says what package the program belongs to.
package main

// Four major kinds of declarations: var, const, func, and type.
import "fmt"

const boilingF = 212.0 
// the above is package-level declaration (as is main)
// it's visible throughout the entire package.

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	// the above is a declaration that are local to the main function.
	// its only visible within the function in which they are declared.
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}

