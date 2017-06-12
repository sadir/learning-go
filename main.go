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
  /* Some local variables, which you must use to avoid compilation errors */
  vname4, vname5, vname6 := 4, 5, 6
  fmt.Printf("Local variables are vname4 with value, %v\n", vname4)
  fmt.Printf("vname5 with value, %v\n", vname5)
  fmt.Printf("and vname6 with value, %v\n", vname6)
  fmt.Printf("Hello, world. Sqrt(2) = %v\n", mymath.Sqrt(2))
}
