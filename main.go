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

type Human struct {
	name, phone  string
	age   int
}

type Employee struct {
	Human     // embedded field Human
	company, speciality, phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

func (h Human) sayHi() {
  fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.primaryContact())
}

func (e Employee) sayHi() {
  fmt.Printf("Hi, I am %s, I work for %s as a %s, you can call me on %s\n",
  e.name, e.company, e.speciality, e.primaryContact())
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) primaryContact() string {
  var phoneNumber string
  if e.phone != "" {
    phoneNumber = e.phone
  } else if e.Human.phone != "" {
    phoneNumber = e.Human.phone
  } else {
    phoneNumber = "n/a"
  }
  return phoneNumber
}

func (h Human) primaryContact() string {
  var phoneNumber string
  if h.phone != "" {
    phoneNumber = h.phone
  } else {
    phoneNumber = "n/a"
  }
  return phoneNumber
}

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
  fmt.Printf("And this in the slice:  %v\n", slice4) //[1,2,3]
  slice7 := append(slice4, 11)
  fmt.Printf("Appending to slices adds the value on the end %v\n", slice7) //[1,2,3,11]
  fmt.Printf("but it doesn't change the original slice as that is a pointer to the array %v\n", slice4) //[1,2,3]
  fmt.Printf("but it does modify the original array that the slices pointed to! %v\n", arr4) //[1,2,3,11,4..10]

  //Maps are the go equivalent of ruby's hashes
  map1 := make(map[string] int, 10) // this initializes a map for you.
  var map2 map[string] int // this creates a nil map. A map which doesn't point to an initialised map in memory.
  map1["one"] = 1
  map1["another one"] = 1
  //map2["a third one"] = 1 This won't work, because map2 doesn't point to a map in memory. You can't add to nil.

  fmt.Printf("map1 is: %v\n", map1) // map[one:1 another one:1] or map[another one:1 one:1]
  fmt.Printf("map2 is: %v\n", map2) // map[]

  delete(map1, "another one")
  fmt.Printf("After delete map1 is: %v\n", map1) // map[one:1]

  _, ok := map1["one"]
  fmt.Printf("You can use the _, ok syntax to check if a map contains a value: %v\n", ok) // true
  map1["another one"] = 1
  map1["a third one"] = 1

  for key, value := range map1 {
    fmt.Println("Key:", key, "Value:", value)
  }
  //Key: one Value: 1
  //Key: another one Value: 1
  //Key: a third one Value: 1
  // Maps are unordered and therefore these can print in any sequence

  type testInt func(int) bool // you can define types of variables to be functions with specifics inputs and outputs.

  e1 := Employee{
    Human: Human{name: "Bob", age: 34, phone: "07888888888"},
    company: "BT",
    speciality: "Designer",
    phone: "07999999999",
  }

  e2 := Employee{
    Human: Human{name: "Bobby", age: 34, phone: "07888888888"},
    company: "Monzo",
    speciality: "Engineer",
  }

  e3 := Employee{
    Human: Human{name: "Bobz", age: 34},
    company: "Myself",
    speciality: "jack of all trades",
  }

  fmt.Printf("employee 1 has a name of: %v\n", e1.name) // bob
  fmt.Printf("employee 1 has a phone of: %v\n", e1.phone) // 07999999999
  fmt.Printf("employee 1 has a personal phone of: %v\n", e1.Human.phone) // 07888888888
  fmt.Printf("employee 1 has a primary contact of: %s\n", e1.primaryContact()) // 07999999999
  e1.sayHi()

  fmt.Printf("employee 2 has a name of: %v\n", e2.name) // Bobby
  fmt.Printf("employee 2 has a phone of: %v\n", e2.phone) // ""
  fmt.Printf("employee 2 has a personal phone of: %v\n", e2.Human.phone) // 07888888888
  fmt.Printf("employee 2 has a primary contact of: %s\n", e2.primaryContact()) // 07888888888
  e2.sayHi()

  fmt.Printf("employee 3 has a name of: %v\n", e3.name) // Bobby
  fmt.Printf("employee 3 has a phone of: %v\n", e3.phone) // ""
  fmt.Printf("employee 3 has a personal phone of: %v\n", e3.Human.phone) // ""
  fmt.Printf("employee 3 has a primary contact of: %s\n", e3.primaryContact()) // n/a
  e3.sayHi()
}
