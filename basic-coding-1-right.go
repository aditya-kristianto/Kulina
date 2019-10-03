package main

import "fmt"

type Animal interface {
	sound() string
	getSound()
}

type AbstractAnimal struct{ Animal }

func (a AbstractAnimal) getSound() {
	fmt.Println(a.sound())
}

type Cat struct{ AbstractAnimal }

func NewCat() *Cat {
	cat := Cat{AbstractAnimal{}}
	cat.AbstractAnimal.Animal = cat

	return &cat
}

func (c Cat) sound() string {
	return "meow"
}

type Dog struct{ AbstractAnimal }

func NewDog() *Dog {
	dog := Dog{AbstractAnimal{}}
	dog.AbstractAnimal.Animal = dog

	return &dog
}

func (d Dog) sound() string {
	return "wooogh"
}

type Mouse struct{ AbstractAnimal }

func NewMouse() *Mouse {
	mouse := Mouse{AbstractAnimal{}}
	mouse.AbstractAnimal.Animal = mouse

	return &mouse
}

func (m Mouse) sound() string {
	return "ciiit"
}

func main() {
	c := NewCat()
	c.getSound()

	d := NewDog()
	d.getSound()

	m := NewMouse()
	m.getSound()
}
