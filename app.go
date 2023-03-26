package main

import "fmt"

func main() {
	type igor_st struct {
		name   string
		age    int
		mybool bool
	}

	c := igor_st{name: "igor", age: 34, mybool: false}

	fmt.Println(c.age, c.name, c.mybool)

}
