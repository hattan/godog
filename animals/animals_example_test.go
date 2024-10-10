package animals

import (
	"testing"

	"gotest.tools/assert"
)

// Test creating a new dog with valid values
func TestAnimalDisplay(t *testing.T) {
	// arrange
	dog, err := NewDog("Fido", 5, Corgi)

	//act
	ShowAnimal(dog)

	// assert
	assert.NilError(t, err, "Expected no error, got %v", err)
	assert.Equal(t, "Fido", dog.Name, "Expected name to be 'Fido', got %s", dog.Name)
	assert.Equal(t, 5, dog.Age, "Expected age to be 5, got %d", dog.Age)
	assert.Equal(t, Corgi, dog.Breed, "Expected breed to be Corgi, got %v", dog.Breed)
}
