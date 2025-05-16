package puppy

import (
	"fmt"

	"github.com/commitium/dog"
)

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof woof woof"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}


func From11() {
	fmt.Println("I'm from version 1.1.0")
}