package main

// importing package
import (
	"errors"
	"fmt"
)

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
	} else {
		fmt.Println("else case")
	}
	// calling the func
	printday("tuesday")

}

// switch cases
func printday(day string) {
	switch day {
	case "monday":
		fmt.Println("its monday")
	case "tuesday":
		fmt.Println("its tuesday")
	case "wednesday":
		fmt.Println("its wednesday")
	case "thursday":
		fmt.Println("its thursday")
	case "friday":
		fmt.Println("its friday")
	default:
		fmt.Println("its an unknown day")
	}

	// arrays

	var arr = [2]int{1, 2}
	fmt.Println(arr)

	var arr2 = [7]int{3, 4, 4, 4, 6, 7, 8}
	fmt.Println(arr2[1])

	// changing the value
	arr2[1] = 5
	fmt.Println(arr2[1])

	arr3 := [5]int{3, 4} //partially initialized [3,4,0,0,0]
	fmt.Println(arr3)

	arr4 := [5]int{3, 4, 4, 4, 4} //fully initialized [3,4,4,4,4]
	fmt.Println(arr4)

	arr5 := [5]int{} //not initialized [0,0,0,0,0]
	fmt.Println(arr5)

	arr6 := [3]int{1: 50} //specific initializing [0,50,0]
	fmt.Println(arr6)

	// len(arr6) -- to get length of array

	// slices
	// array is fixed size but slices are more commonly used because of its flexibility.
	nameArr := [3]string{"john", "doe"}  //using a size makes it an array
	nameSlice := []string{"john", "doe"} //dont using a size in [] makes it a slice
	fmt.Println(nameSlice, nameArr)

	// appending data to slice
	nameSlice = append(nameSlice, "doe")
	fmt.Println(nameSlice)
	// we cannot use append method on array because of its fixed size

	// check the capacity of the slice
	fmt.Println(cap(nameSlice))

	// passing to functions

	// ----
	// arrays are passed ny value, meaning a copy is made.changes made to the array doesnt affect the original

	// ----
	// slices are passed by reference, meaning changes to the slice inside a function affect the original data

	// slicing an array
	array := [5]int{5, 4, 23, 12, 67}

	mySlice := array[1:2] //from element at 1 index till the 2(element at index 2 is not included)

	// output : [4]
	fmt.Println(mySlice)

	// loops
	// go only have a for loop for iterating

	// traditional loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterating through array /slice
	nums := []int{3, 4, 5, 6}

	// i gets the index and num get the value at the index
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// pointers
	// lets create a variable
	var num int = 10

	// lets store the memory address of the num in a pointer variable
	var numPointer *int = &num

	fmt.Println("value of num : ", num)
	fmt.Println("pointer address : ", numPointer)

	// updating the value of num using the pointer
	*numPointer = 20
	fmt.Println("updated value of num : ", num)

	// structs
	// we can declare a structure using TYPE ans STRUCT keywords

	// for example we can create a person struct with name and age properties

	type Person struct {
		name string
		age  int8
	}

	var person = Person{name: "john", age: 22}

	fmt.Println(person)

	// changing the value of name
	person.name = "ajeesh"

	fmt.Println(person)

	// we can access the values of each using this
	var username string = person.name

	fmt.Println(username)

	// passing struct in functions
	// struct User and printUser function is defined outside the main function
	var userDetails User

	userDetails.name = "ajeesh"
	userDetails.age = 22
	userDetails.isMarried = false

	printUser(userDetails)

	// error handling
	// getting the result and err form the divideFunc
	// in go err is returned like a normal value
	result, err := divideFunc(5, 0)

	//checking if it got any error using the err!= nil
	// if so printing the err and returning
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Result: ", result)

	
}

// user type
type User struct {
	name      string
	age       int8
	isMarried bool
}

// function which take argument in User struct
func printUser(user User) {

	fmt.Println("name : ", user.name)
	fmt.Println("age : ", user.age)
	fmt.Println("married : ", user.isMarried)

}

// divide function which have a not divisible by zero check
func divideFunc(a, b float64) (float64, error) {
	if b == 0 {
		// returning 0 and a custom err message
		return 0, errors.New("not divisible by zero ")
	}
	// retun result,nil
	return a / b, nil
}
