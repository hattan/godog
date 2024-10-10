package animals

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

// Test creating a new dogbreed with valid values
func TestDogBreedFromString(t *testing.T) {
	// arrange
	breeds := []string{"corgi", "frenchie", "poodle", "mutt"}

	//act
	for _, breed := range breeds {
		breed = strings.ToLower(breed)
		dogBreed, _ := FromString(breed)
		expected := strings.ToLower(dogBreed.String())
		assert.Equal(t, breed, expected)
	}
}

// Test invalid dogbreed string returns corgi (default)
func TestInvalidDogBreedFromStringReturnsCorgi(t *testing.T) {
	_, err := FromString("notavalidbreed")
	assert.Error(t, err, "\"notavalidbreed\" is not a valid dog breed")
}

// Test that and invalid DogBreed returns unknown
func TestInvalidDogBreedToStringReturnsUnknown(t *testing.T) {
	breed := DogBreed(20)
	breedStr := breed.String()

	assert.Equal(t, "unknown", breedStr)
}

// Test that IsMutt returns true for Mutt
func TestIsMutt(t *testing.T) {
	dog, _ := NewDog("Mutty", 5, Mutt)
	assert.Assert(t, dog.IsMutt())
}
