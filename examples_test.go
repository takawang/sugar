package sugar_test

import (
	"testing"
	"github.com/marksalpeter/sugar"
)

type Struct struct {
	Field string
}

func Example() {
	
	sugar.New(&testing.T{}).

	Must("this must be true or t.FailNow() will be called", func (_ sugar.Log) bool {
		return true
	}).
	
	Warn("this won't cause a failure, it just prints a warning", func (_ sugar.Log) bool {
		return true
	}).
	
	Assert("this must be true or t.Fail() will be called!", func (_ sugar.Log) bool {
		return false
	}).
	
	Assert("and finally this is a full demonstration of the logger", func (log sugar.Log) bool {
		
		log("by default, %s works like fmt.Printf", "sugar.Log")
		
		log("but, if you just pass in structs, it will print them with their field names")
		log(&Struct{ Field: "1" }, &Struct{ Field: "2" }, &Struct{ Field: "3" })
		
		nestedLogger := sugar.NewLogger()
		nestedLogger.Log("finally, its possible to nest logs by createing a new logger")
		log(nestedLogger)
				
		return true
	})
	
	// Output: 
	// PASS	    5.042µs	this must be true or t.FailNow() will be called
	// PASS	    1.111µs	this won't cause a failure, it just prints a warning
	// FAIL	    1.187µs	this must be true or t.Fail() will be called!
	// PASS	   13.006µs	and finally this is a full demonstration of the logger
	//  ┠ by default, sugar.Log works like fmt.Printf
	//  ┠ but, if you just pass in structs, it will print them with their field names
	//  ┠ &{Field:1}
	//  ┠ &{Field:2}
	//  ┠ &{Field:3}
	//  ┖  ┖ finally, its possible to nest logs by createing a new logger
	
}

func ExampleSugar() {
	
	sugar.New(&testing.T{}).

	Title("welcome to the sugar example").
	
	Must("this must be true or t.FailNow() will be called", func (_ sugar.Log) bool {
		return true
	}).
	
	Assert("this must be true or t.Fail() will be called!", func (_ sugar.Log) bool {
		return false
	}).
	
	Warn("this won't cause a failure, it just prints a warning", func (_ sugar.Log) bool {
		return false
	})
	
	// Output:
	// >>> welcome to the sugar example <<<
	// PASS	    1.231µs	this must be true or t.FailNow() will be called
	// FAIL	    1.078µs	this must be true or t.Fail() will be called!
	// WARN	    1.005µs	this won't cause a failure, it just prints a warning
}

func ExampleLogger() {
	
	sugar.New(&testing.T{}).

	Assert("this is a full demonstration of the logger", func (log sugar.Log) bool {
		
		log("by default, %s works like fmt.Printf", "sugar.Log")
		
		log("but, if you just pass in structs, it will print them with their field names")
		log(&Struct{ Field: "1" }, &Struct{ Field: "2" }, &Struct{ Field: "3" })
		
		nestedLogger := sugar.NewLogger()
		nestedLogger.Log("finally, its possible to nest logs by createing a new logger")
		log(nestedLogger)
				
		return true
	})
	
	// Output:
	// PASS	   13.006µs	and finally this is a full demonstration of the logger
	//  ┠ by default, sugar.Log works like fmt.Printf
	//  ┠ but, if you just pass in structs, it will print them with their field names
	//  ┠ &{Field:1}
	//  ┠ &{Field:2}
	//  ┠ &{Field:3}
	//  ┖  ┖ finally, its possible to nest logs by createing a new logger
}