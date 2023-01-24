package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type A struct {
	str string
}

var singleInstance *A

func (a *A) getAValue() string {
	return a.str
}

func (a *A) setAValue(str string) {
	a.str = str
}

func getInstance() *A {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating singleton object!")
				singleInstance = &A{}
			})
	} else {
		fmt.Println("object is already created")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 10; i++ {
		obj := getInstance()
		var s string
		fmt.Println("Enter a string")
		fmt.Scan(&s)
		obj.setAValue(s)
		fmt.Println(obj.getAValue())
	}
}
