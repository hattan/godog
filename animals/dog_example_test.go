package animals

import (
	"testing"

	"gotest.tools/assert"
)

// Test creating a new dog with valid values
func TestNewDog(t *testing.T) {
	// act
	dog, err := NewDog("Fido", 5, Corgi)

	// assert
	assert.NilError(t, err, "Expected no error, got %v", err)
	assert.Equal(t, "Fido", dog.Name, "Expected name to be 'Fido', got %s", dog.Name)
	assert.Equal(t, 5, dog.Age, "Expected age to be 5, got %d", dog.Age)
	assert.Equal(t, Corgi, dog.DogBreed, "Expected breed to be Corgi, got %v", dog.DogBreed)
}

// Test creating a new dog with invalid age
func TestNewDogInvalidAge(t *testing.T) {
	// act
	_, err := NewDog("Max", -1, Frenchie)

	// assert
	assert.Error(t, err, "age must be greater than zero")
}

// Test that the dog's age can be updated
func TestSetAge(t *testing.T) {
	// arrange
	dog, _ := NewDog("Gaston", 4, Frenchie)

	//act
	dog.SetAge(5)

	//assert
	assert.Equal(t, 5, dog.Age)
}

// Test that the dog's information is displayed correctly
func TestDogDisplay(t *testing.T) {
	// arrange
	dog, _ := NewDog("Gaston", 4, Frenchie)

	// act
	dog.Display()
}
