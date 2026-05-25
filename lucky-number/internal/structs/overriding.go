package main

import "fmt"

type Engine struct {
	Name  string
	Model string
	power int
}

type Car struct {
	Engine
	Color string
}

func (e Engine) powerUP() {
	e.power += 2
}

func (c Car) powerUP() {
	c.power++
}

func testOverriding() {
	car1 := Car{
		Engine: Engine{
			Name:  "V8",
			Model: "V8",
			power: 1,
		},
		Color: "red",
	}

	car1.powerUP()
	fmt.Println(car1)

}
