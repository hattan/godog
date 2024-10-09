// Package animals provides a simple abstraction for representing animals.
//
// I'm so sorry this is basic and contrived. I'm not a zoologist.
package animals

// This is not used at all. I just put this here so the docs don't feel lonely
const IS_ANIMAL = true

// This looks impressive, it does nothing
var _ Animal = (*Dog)(nil)

// Animal is an interface that defines the behavior of an animal.
// It's used to demonstrate how interfaces work in Go during the workshop.
// I completely understand that this is a contrived example and that animals in real life don't have a Display method.
type Animal interface {
	Display()
}

// ShowAnimal prints the underlying data of the animal to stdout.
// Intentionally not a method with receiver, but maybe in a future update?
//   - animal - the animal to display.

func ShowAnimal(animal Animal) {
	animal.Display()
}
