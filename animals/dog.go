package animals

import (
	"errors"
	"fmt"
	"log/slog"
)

// Dog is a representation of a dog and contains a name, age, and breed.
// Don't use this type directly, use NewDog instead.
type Dog struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Breed DogBreed `json:"breed"`
}

// NewDog creates a new instance of Dog with the specified values.
//   - name - a unique name for this dog; common names include 'fido' and 'max'.
//   - age - the current age of the dog. Must be greater than zero.
//   - breed - the breed of the dog. Use the DogBreed constants to specify the breed.
func NewDog(name string, age int, breed DogBreed) (*Dog, error) {
	if age <= 0 {
		return nil, errors.New("age must be greater than zero")
	}
	dog := Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}

	slog.Info(fmt.Sprintf("New Dog Created %s", dog.Name))
	return &dog, nil
}

// Display prints the dog's information in a nice format. It's printed to stdout.
func (d *Dog) Display() {
	fmt.Printf("Dog\n Name:%s\n Age:%d\n Breed:%s\n", d.Name, d.Age, d.Breed)
}

// SetAge is a setter method that updates the dog's age.
// This is used in the workshop portion that discusses pointers.
//   - age - the new age of the dog. Must be greater than zero.
func (d *Dog) SetAge(age int) {
	d.Age = age
}
