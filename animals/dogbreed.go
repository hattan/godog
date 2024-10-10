package animals

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

type DogBreed int

const (
	Corgi    DogBreed = 0
	Frenchie DogBreed = 1
	Poodle   DogBreed = 2
	Mutt     DogBreed = 3
)

var Breed_value = map[string]DogBreed{
	"corgi":    Corgi,
	"frenchie": Frenchie,
	"poodle":   Poodle,
	"mutt":     Mutt,
}

// Converts a DogBreed type to a string. Used for JSON marshalling and logging
func (breed DogBreed) String() string {
	key, ok := getBreedKeyByValue(Breed_value, breed)
	if !ok {
		slog.Error(fmt.Sprintf("Error DogBreedToString. %d is not a valid breed", breed))
		return "unknown"
	}
	return key
}

// Returns a breed name from a DogBreed value by utilizing the built-in map
func getBreedKeyByValue(m map[string]DogBreed, value DogBreed) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

// IsMutt returns true if the DogBreed is Mutt
func (breed DogBreed) IsMutt() bool {
	return breed == Mutt
}

// MarshalJSON marshals the DogBreed to a JSON string
func (breed DogBreed) MarshalJSON() ([]byte, error) {
	return json.Marshal(breed.String())
}

// UnmarshalJSON unmarshals a JSON string to a DogBreed
func (s *DogBreed) UnmarshalJSON(data []byte) (err error) {
	var breed string
	if err := json.Unmarshal(data, &breed); err != nil {
		return err
	}
	if *s, err = FromString(breed); err != nil {
		return err
	}
	return nil
}

// FromString converts a string to a DogBreed
func FromString(s string) (DogBreed, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := Breed_value[s]
	if !ok {
		return DogBreed(0), fmt.Errorf("%q is not a valid dog breed", s)
	}
	return DogBreed(value), nil
}
