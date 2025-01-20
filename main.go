package main

// importing package
import "fmt"

// entry point
func main() {
	// fmt - formating package which have function println similar to console.log in js/ts
	fmt.Println("go")

	// variables and data types

	// string
	// int (int,int8,int 16,int32,int64 - same for unsigned prefix-uint)
	// float(float32,float64)
	// bool (true,false)
	// pointer type -uintptr

	var name string = "go"
	var age int = 22
	var isTrue bool
	var sample float32 = 3
	var unusedVar string = "i am unused variable" //this line will throw an error because in go we need to use the variables that we declared.

	// const sum string -----
	// this line will throw error because we need to initialize const variables with a value when declariing -(not applicable for normal VARs)

	var ting int //if we do like this it will be initialized with 0 if type is int, empty string if string type,false if bool type

	fmt.Println(ting)
	fmt.Println(unusedVar)
	fmt.Println(sample)

	// it will return a formted concatenated string (fmt.Sprintf)
	variable := fmt.Sprintf("Hi %s,i am %d years old", name, age)

	// output : Hi go,i am 22 years old
	fmt.Println(variable)

	// if statement
	if isTrue {
		// this wont print because isTrue is false
		fmt.Println("inside false if statement")
	}

	isTrue = true

	if isTrue {
		// this will print because we changes the isTrue to be trueƒ
		fmt.Println("inside true if statement")
	}
}
