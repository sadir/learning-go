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

  //you can let go calculate the array length based on the input with ...
  arr3 := [...]int{1, 2, 3}
  fmt.Printf("The third array is length %d\n", len(arr3)) //3
}
