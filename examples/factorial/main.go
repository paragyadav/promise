package main

import (
	"fmt"

	"github.com/paragyadav/promise"
)

func findFactorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * findFactorial(n-1)
}

func main() {
	var factorial = promise.Resolve(findFactorial(5))
	result, error := factorial.Await()

	fmt.Println("factorial result : ", result)
	fmt.Println("factorial error : ", error)
}
