package animals

import (
	"errors"
	"fmt"
	"log/slog"
)

type Dog struct {
	Name string
	Age  int
	DogBreed
}

func NewDog(name string, age int, breed DogBreed) (*Dog, error) {
	if age <= 0 {
		return nil, errors.New("age must be greater than zero")
	}
	dog := Dog{
		Name:     name,
		Age:      age,
		DogBreed: breed,
	}

	slog.Info(fmt.Sprintf("New Dog Created %s", dog.Name))
	return &dog, nil
}

func (d *Dog) Display() {
	fmt.Printf("Dog\n Name:%s\n Age:%d\n Breed:%s\n", d.Name, d.Age, d.DogBreed)
}

func (d *Dog) SetAge(age int) {
	d.Age = age
}
