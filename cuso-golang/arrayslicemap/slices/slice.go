package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[:2]
	fmt.Println(a2, s2)
	fmt.Println(&a2[0], a2[0])
	fmt.Println(&s2[0], s2[0])

	fmt.Println(&a2[1], a2[1])
	fmt.Println(&s2[1], s2[1])

}
