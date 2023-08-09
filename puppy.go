package puppy

import (
	"github.com/ManaponS/dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof wOOf wooF"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func Ver1() string {
	return "1.0.0"
}
