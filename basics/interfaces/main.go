package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

type Dog struct{}

func (p Person) Speak() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func (d Dog) Speak() string {
	return "woof woof"
}

func GetSpeech(s Speaker) string {
	return s.Speak()
}

func main() {
	p := Person{
		Name: "Ars",
	}

	d := Dog{}

	fmt.Println(GetSpeech(p))
	fmt.Println(GetSpeech(d))
}
