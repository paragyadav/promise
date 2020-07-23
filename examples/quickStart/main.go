package main

import (
	"errors"
	"fmt"

	"github.com/paragyadav/promise"
)

func main() {
	var p = promise.New(func(resolve func(interface{}), reject func(error)) {
		// Do something asynchronously.
		const sum = 2 + 2

		// If your work was successful call resolve() passing the result.
		if sum == 4 {
			resolve(sum)
			return
		}

		// If you encountered an error call reject() passing the error.
		if sum != 4 {
			reject(errors.New("2 + 2 doesnt't equal 4"))
			return
		}
	}).
		// You may continue working with the result of
		// a previous async operation.
		Then(func(data interface{}) interface{} {
			fmt.Println("The result is:", data)
			return data.(int) + 1
		},
			//   To consume an error and return a new value for next then operations.
			func(err error) interface{} {
				fmt.Println("The error is: ", err)
				return "recovered from error."
				// To throw a new error after consuming the previous error.
				// return promise.Reject(errors.New("math: square root of negative number"))
			}).
		// Handlers can be added even after the success or failure of the asynchronous operation.
		// Multiple handlers may be added by calling .Then or .Catch several times,
		// to be executed independently in insertion order.
		Then(func(data interface{}) interface{} {
			fmt.Println("The new result is:", data)
			return "successfully executed"
		}, nil).
		Catch(func(error error) error {
			fmt.Println("Error during execution:", error.Error())
			return nil
		})

		// Since handlers are executed asynchronously you can wait for them
		// Or you can add a finally block which internally takes care of Await.

	pResult, pError := p.Await()
	// result, error := p.Finally(func() {
	// 	fmt.Println("All is Well.")
	// })

	fmt.Println("promise result : ", pResult)
	fmt.Println("promise error : ", pError)
}
