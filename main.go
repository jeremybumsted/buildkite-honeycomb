package main

import(
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	x := Add(1, 2)
	y := Subtract(5,2)

	r := fmt.Sprint(x, y)

	fmt.Println(r)
}

// Add is our function that sums two integers
func Add(x, y int) (res int) {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}

func Multiply(x, y int) (res int) {
	return x * y
}

func Divide(x, y int) (res int) {
	return x / y
}

// RandomBoolean returns a random boolean value
func RandomBoolean() bool {
	number := rand.Intn(10) + 1  // Generates a number between 1 and 10
	return number < 7           // Returns true for numbers 1-6, false for 7-10
}
