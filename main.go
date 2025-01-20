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
}
