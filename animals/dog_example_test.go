package animals

import (
	"encoding/json"
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
	assert.Equal(t, Corgi, dog.Breed, "Expected breed to be Corgi, got %v", dog.Breed)
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

// Test that the dog's breed is un-marshalled correctly
func TestDogBreedUnmarshalJSON(t *testing.T) {
	// arrange
	expectedName := "new test dog"
	expectedAge := 7
	expectedBreed := Frenchie
	var jsonContent = `{
		"name" : "new test dog",
		"age" : 7,
		"breed": "frenchie"
	}`
	// act
	var dog *Dog
	bytes := []byte(jsonContent)
	json.Unmarshal(bytes, &dog)

	// assert
	assert.Equal(t, expectedName, dog.Name)
	assert.Equal(t, expectedAge, dog.Age)
	assert.Equal(t, expectedBreed, dog.Breed)
}
