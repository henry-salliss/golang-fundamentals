package main

import (
	"fmt"
)



func main() {  
	//typed and untyped constants
	var defaultName = "Sam" //allowed defaults to string untyped
	type myString string
	var customName myString = "Sam" //allowed typed constant is string
	customName = defaultName //not allowed cant mix typed and untyped

	// boolean constants
	const trueConst = true
    type myBool bool
    var defaultBool = trueConst //allowed
    var customBool myBool = trueConst //allowed
    defaultBool = customBool //not allowed

	// numeric constants
	const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

}