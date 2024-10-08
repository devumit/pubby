package puppy

import (
	"github.com/devumit/dog"
)

func Bark() string {
	return "Woof!"
}
func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark()) + " " + dog.WhenGrownUp(Bark())
}

func BigBarksNewVersion() string {
	return dog.WhenGrownUp(Bark()) + " " + dog.WhenGrownUp(Bark())
}
