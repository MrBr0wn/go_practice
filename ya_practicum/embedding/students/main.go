package main

import (
	"fmt"
)

type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Year of birth: %d", p.Name, p.Year)
}

func (p Person) Print() {
	fmt.Println(p)
}

type Faculty string

func (f Faculty) String() string {
	return string(f)
}

func (f Faculty) Print() {
	fmt.Println(f)
}

type Student struct {
	Person
	Faculty
	Group string
}

func NewStudent(name, faculty, group string, year int) Student {
	return Student{
		Person:  NewPerson(name, year),
		Faculty: Faculty(faculty),
		Group:   group,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("%s, Faculty: %s, Group: %s", s.Person, s.Faculty, s.Group)
}

func main() {
	s := NewStudent("John Doe", "MTF", "701", 1980)
	// s.Print() // will be error because Faculty and Person has some method

	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Faculty, s.Group)
}
