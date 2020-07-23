# PROMISE

## About
Promises library for Golang. 

Inspired by [JS Promises](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) and [chebyrash/promise](https://github.com/chebyrash/promise)

## Install

    $ go get -u github.com/paragyadav/promise

## Quick Start
```go
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
	// You may continue working with the result of a previous async operation.
	Then(func(data interface{}) interface{} {
		fmt.Println("The result is:", data)
		return data.(int) + 1
	    },
		// To consume an error and return a new value for next then operations.
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
```

## Methods

### New

To create a new promise.

Function Signature : New(executor func(resolve func(interface{}), reject func(error))) *Promise

### Then

To chain on the return value of a promise.

Function Signature : Then(onFulfilled func(data interface{}) interface{}, onRejected func(data interface{}) interface{}) *Promise

### Catch

To catch on the errors returned during promise execution.

Function Signature : Catch(rejection func(err error) error) *Promise

### Finally

When the promise is settled, i.e either fulfilled or rejected, the specified callback function is executed. This provides a way for code to be run whether the promise was fulfilled successfully or rejected once the Promise has been dealt with. 

Function Signature : Finally(onFinally func() (interface{}, error))

### Resolve

Returns a new Promise that is resolved with the given value. If the value is a thenable (i.e. has a then method), the returned promise will "follow" that thenable, adopting its eventual state; otherwise the returned promise will be fulfilled with the value.

Function Signature : Resolve(resolution interface{}) *Promise

Example:
```go
var p1 = promise.Resolve("Hello, World")
result, _ := p1.Await()
fmt.Println(result)
// Hello, World
```

### Reject

Returns a new Promise that is rejected with the given reason.

Function Signature : Reject(err error) *Promise

Example:
```go
var p1 = promise.Reject(errors.New("bad error"))
_, err := p1.Await()
fmt.Println(err)
// bad error
```

### Await

To Wait on an already created promise.
Returns the result and error value pair.

Function Signature : Await() (interface{}, error)

Example:
```go
var p1 = promise.Resolve("Hello, World")
result, error := p1.Await()
fmt.Println("Result : ",result)
fmt.Println("Error : ",error)
// Result : Hello, World
```

