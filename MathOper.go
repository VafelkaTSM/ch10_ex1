// This library implements mathematical operations
// The best website about mathematics https://www.mathsisfun.com/numbers/addition.html
package MathOper

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// A function that adds two int numbers and outputs their result as an int
func Add[T Number](a, b T) T {
	return a + b
}
