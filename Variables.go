package main

import ("fmt")

//this is introduction to the variable and constants and some other data types
//To introduce a variable you must define its type or Behaviour or infer it. Like
func main() {

	//the standard form of declaring a variable can also be used outside the Function.

	var firstVariable string = "string Variable"

	// := this is the sign to use a short method to define the Data Type.
	//short hand Method can only be use inside a function
	shorthand_Variable := "Also a String Variable"

	//we can also print an variable which is not Aasign but output is showing
	// Empty space
	var Empty_string string
	var Empty_Integer int
	var Empty_boolean bool

	//An integer can be declared(must define the type while standard form)

	var first_integer int = 80

	//  and thier is no need of declaring "int" key word.

	shorthand_Integer := 90

	// just printin the output Line by Line:
	fmt.Println(firstVariable)
	fmt.Println(shorthand_Variable)
	fmt.Println(Empty_string)
	fmt.Println(Empty_Integer)
	fmt.Println(Empty_boolean)
	fmt.Println(first_integer)
	fmt.Println(shorthand_Integer)

}
