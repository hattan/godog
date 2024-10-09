package animals

type DogBreed int

const (
	Corgi    DogBreed = 0
	Frenchie DogBreed = 1
	Poodle   DogBreed = 2
	Mutt     DogBreed = 3
)

func (breed DogBreed) String() string {
	switch breed {
	case Corgi:
		return "Corgi"
	case Frenchie:
		return "Frenchie"
	case Poodle:
		return "Poodle"
	case Mutt:
		return "Mutt"
	default:
		return "unknown"
	}
}

func FromString(breed string) DogBreed {
	switch breed {
	case "Corgi":
		return Corgi
	case "Frenchie":
		return Frenchie
	case "Poodle":
		return Poodle
	case "Mutt":
		return Mutt
	default:
		return Corgi
	}
}
func (breed DogBreed) IsMutt() bool {
	return breed == Mutt
}
