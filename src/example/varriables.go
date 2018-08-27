package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//var x int
	//fmt.Println(x)
	//	it will be initialize as 0;

	//var a int = 5
	//var b int = 7
	//var sum int = a + b

	//we can do it without var but like this

	//c := 6
	//d := 7

	//fmt.Println("c is", c ,"and d is", d)

	//fmt.Println("sum is =", sum)

	//if c > 7 {
	//	fmt.Println("WTC")
	//} else if d < 8 {
	//	fmt.Println("nothing")
	//}	else{
	//	fmt.Println("nothing common")
	//}

	//	array
	//	var ab [5]int
	//	ab[2] = 7
	//	fmt.Println(ab)

	//	for initialization we always use :=

	//	unspecified length array called slice :)
	// surly go will made me crazy
	//	a := []int{4,2,3,6,6}
	//	a[5] = 6
	//fmt.Println(a[2])

	//a = append(a, 13)
	//fmt.Println(a)

	//dictionary objects
	//vertices := make(map[string]int)
	//vertices["a"] = 1
	//vertices["b"] = 1
	//vertices["c"] = 1
	//fmt.Println(vertices["a"])
	//delete(vertices, "c")
	//fmt.Println(vertices)

	//loops only for loops in go
	//arr := []string{"a", "b", "c"}
	//m := make(map[string]string)
	//m["a"] = "alpha"
	//m["b"] = "beta"
	//for index, value := range arr {
	//	fmt.Println(index)
	//	fmt.Println(value)
	//}
	//
	//for key, value := range m {
	//	fmt.Println(key)
	//	fmt.Println(value)
	//}
	//fmt.Println(sum(6, 9))
	//result, err := sqrt(12)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(result)
	//}

	//struct
	//p := person{name : "Abdul Hameed", age: 27}
	//fmt.Println(p)

	//pointers

	i := 7
	//fmt.Println(i)
	//fmt.Println(&i)
	inc(&i)
	fmt.Println(i)

	fmt.Println(withoutInc(i))
}

func sum(x int, y int) int {
	return (x + y) * x
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("why you supplied a nonsense")
	}

	return math.Sqrt(x), nil
}

type person struct {
	name string
	age  int
}

func inc(x *int) {
	*x++
}

func withoutInc(x int) int {
	var y = x - 2
	return y
}
