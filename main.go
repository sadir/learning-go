package main

import (
  "mymath"
  "fmt"
)

/*
Some global variables
*/
var vname1, vname2, vname3 = 1, 2, 3

/* A few constants */
const Pi = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "sadir_"

func main() {
  // Some local variables, which you must use to avoid compilation errors
  vname4, vname5, vname6 := 4, 5, 6
  fmt.Printf("Local variables are vname4 with value, %v\n", vname4)
  fmt.Printf("vname5 with value, %v\n", vname5)
  fmt.Printf("and vname6 with value, %v\n", vname6)
  fmt.Printf("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))

  //arrays are typed based on their length and the data type within them.
  var arr [10]int
  //integer arrays default to 0 values.
  fmt.Printf("The first element is %d\n", arr[0]) //0
  fmt.Printf("The last element is %d\n", arr[9]) //also 0

  //you can define arrays inline
  arr2 := [10]int{1, 2, 3}
  fmt.Printf("The first element in the second array is %d\n", arr2[0]) //1
  fmt.Printf("The third element in the second array is %d\n", arr2[2]) //3
  fmt.Printf("The last element in the second array is %d\n", arr2[9]) //still 0 if not defined

  //you can let Go calculate the array length based on the input with '...'
  arr3 := [...]int{1, 2, 3}
  fmt.Printf("The third array is length %d\n", len(arr3)) //3

  //slices are a bit like dynamic arrays
  slice := []byte {'a', 'b', 'c', 'd'}

  //you can define slices from arrays or other slices
  arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  var slice1 []int //it's variable in length so no length needed in the square brackets
  slice2 := arr4[0:2] // [1, 2] but not 3 because it's exclusive. Cap is 10.
  slice3 := arr4[0:3] // [1,2,3] to get all three you have to use the length of the array. Cap is still 10.
  slice4 := arr4[:3] // [1,2,3] leaving out the starting index will use the arrays first item aka 0 index. Cap once again 10.
  slice5 := arr4[:] // [1,2..10] leaving out the other will use the array length. You can leave out both to reference the whole array.
  slice6 := slice4[0:10] /* [1,2..10] when you create a slice from a larger slice or array,
  you can still reference the extra capacity ('cap') not included by default in the slice. */

  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice)) //4
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice1)) //0
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice2)) //2
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice3)) //3
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice4)) //3
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice5)) //10
  fmt.Printf("I still have to use the slices & arrays if I define them %d\n", len(slice6)) //10

  //Append has some interesting functionality
  fmt.Printf("Before the append we have this in the array: %v\n", arr4) //[1,2,3,11,4..10]
  fmt.Printf("And this in the slice:  %v\n", slice4) //[1,2,3,11,4..10]
  slice7 := append(slice4, 11)
  fmt.Printf("Appending to slices adds the value on the end %v\n", slice7) //[1,2,3,11]
  fmt.Printf("but it doesn't change the original slice as that is a pointer to the array %v\n", slice4) //[1,2,3]
  fmt.Printf("but it does modify the original array that the slices pointed to! %v\n", arr4) //[1,2,3,11,4..10]
}
