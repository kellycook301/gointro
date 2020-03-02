package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func inc(x *int) {
	*x++
}

func main() {
	p := person{name: "Kelly Cook,", age: 27}
	fmt.Println(p)

	i := 7
	inc(&i)
	fmt.Println(i)

	fmt.Println("Hello World")
	x := 387
	y := 8379
	sum := x + y
	fmt.Println(sum)

	if x > 7 {
		fmt.Println("More than 7")
	}

	a := []int{5, 3, 2, 7, 0}
	a = append(a, 13)

	fmt.Println(a)

	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["circle"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value", value)
	}

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
