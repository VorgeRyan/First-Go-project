//coded using sublime text 3

package main

import (
	"fmt"
)

func main() {
	fmt.Println("My First Time Coding On Go!")
}

func main() {
	var x int = 5     //or you can do x := 5 to save space
	var y int = 7     //or you can do x := 7 to save space
	var sum int = x + y

	fmt.Println(x)
}

func main() {
	var a [5]int
	a[2] = 7

	fmt.Println(a)
}

func sqrt(x floar64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}