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

func (breed DogBreed) String() string {
	key, ok := getBreedKeyByValue(Breed_value, breed)
	if !ok {
		slog.Error(fmt.Sprintf("Error DogBreedToString. %d is not a valid breed", breed))
		return "unknown"
	}
	return key
}

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

func (breed DogBreed) IsMutt() bool {
	return breed == Mutt
}

func (breed DogBreed) MarshalJSON() ([]byte, error) {
	return json.Marshal(breed.String())
}

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

func FromString(s string) (DogBreed, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := Breed_value[s]
	if !ok {
		return DogBreed(0), fmt.Errorf("%q is not a valid dog breed", s)
	}
	return DogBreed(value), nil
}
