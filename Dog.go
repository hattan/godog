package dog

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func NewDog(name string, age int, breed string) Dog {
	return Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}
}

func (d Dog) Bark() string {
	return "Woof!"
}
