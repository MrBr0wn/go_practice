package pets

import (
	"errors"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// "fmt"

type Eater interface {
	Eat(amount uint8) (uint8, error)
}

type Walker interface {
	Walk() string
}

type Named interface {
	GetName() string
}

type EaterWalker interface {
	Eater
	Walker
	Named
}

type Animal struct {
	Name string
}

func (a *Animal) GetName() string {
	caser := cases.Title(language.English)
	return caser.String(a.Name)
}

type ActionError struct {
	Name   string
	Reason string
}

func (ae *ActionError) Error() string {
	return fmt.Sprintf("%s cannot perform the action: %s", ae.Name, ae.Reason)
}

func newError(msg string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}

	return errors.New(msg)
}
