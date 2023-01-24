package main

import (
	"fmt"
)

type shapeFactory interface {
	draw()
}
type circle struct {
}

type squre struct {
}

func (cir *circle) draw() {
	fmt.Println("drwaing circle!")
}
func (sq *squre) draw() {
	fmt.Println("drawing square!")
}

func shapeFactoryObject(shape string) shapeFactory {
	var shapeObj shapeFactory
	if shape == "circle" {
		shapeObj = &circle{}
	} else {
		shapeObj = &squre{}
	}
	return shapeObj
}
func main() {
	shapeFactory := shapeFactoryObject("circle")
	shapeFactory.draw()
	shapeFactory = shapeFactoryObject("squre")
	shapeFactory.draw()
}
