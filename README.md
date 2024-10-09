# godog

This is a simple go module that is used in a the [godog workshop](https://github.com/hattan/godogworkshop)

The workshop is meant to be an easy onboarding and training to help software engineers quickly upskill on the go programming language.

## Examples

```go
package main

import (
    "github.com/hattan/godog/animals"
)

func main() {
    dog, err := animals.NewDog("Fido", 2, animals.Corgi)
    if err != nil {
        panic(err)
    }

    dog.Display()
}

```
