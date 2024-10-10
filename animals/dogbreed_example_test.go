package animals

import (
	"encoding/json"
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
	assert.Assert(t, dog.Breed.IsMutt())
}

// Test that the dog's breed is un-marshalled correctly
func TestDogBreedInvalidBreedName(t *testing.T) {
	// arrange
	var jsonContent = `{
		"name" : "new test dog",
		"age" : 7,
		"breed": "FakeBreed"
	}`
	// act
	var dog *Dog
	bytes := []byte(jsonContent)
	err := json.Unmarshal(bytes, &dog)

	// assert
	assert.Error(t, err, "\"fakebreed\" is not a valid dog breed")
}

func TestDogBreedInvalidBreedValue(t *testing.T) {
	// arrange
	var jsonContent = `{
		"name" : "new test dog",
		"age" : 7,
		"breed": 2
	}`
	// act
	var dog *Dog
	bytes := []byte(jsonContent)
	err := json.Unmarshal(bytes, &dog)

	// assert
	assert.Error(t, err, "json: cannot unmarshal number into Go struct field Dog.breed of type string")
}

func TestDogBreedMarshalJson(t *testing.T) {
	// arrange
	dogBreed := Corgi
	// act
	bytes, _ := json.Marshal(dogBreed)

	str := string(bytes)
	// assert
	assert.Equal(t, "\"corgi\"", str)
}
