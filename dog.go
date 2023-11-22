package animals

import (
	"errors"
	"fmt"
)

type Dog struct {
	Name string
	age  int
	DogBreed
}

func NewDog(name string, age int, breed DogBreed) (*Dog, error) {
	if age < 0 {
		return nil, errors.New("age cannot be less than zero")
	}

	dog := Dog{
		Name:     name,
		age:      age,
		DogBreed: breed,
	}

	logger.Info(fmt.Sprintf("New Dog Created %s", dog.Name))
	return &dog, nil
}

func (dog *Dog) Display() {
	fmt.Printf("Dog\n Name: %s\n Age: %d\n Breed:%s\n", dog.Name, dog.age, dog.DogBreed)
}

func (dog *Dog) SetAge(age int) {
	dog.age = age
}
