package main

import (
	"errors"
	"fmt"
	"time"
)

type TimeError struct {
	Time time.Time
	Text string
}

func (te TimeError) Error() string {
	return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}

func NewTimeError(text string) TimeError {
	return TimeError{
		Time: time.Now(),
		Text: text,
	}
}

func testFunc(i int) error {
	if i == 0 {
		return NewTimeError(`параметр в testFunc = 0`)
	}

	return nil
}

func main() {
	if err := testFunc(0); err != nil {
		fmt.Println(err)
	}

	if err := testFunc(0); err != nil {
		if v, ok := err.(TimeError); ok {
			fmt.Println(v.Time, v.Text)
		} else {
			fmt.Println(err)
		}
	}

	if err := testFunc(0); err != nil {
		var te TimeError
		if ok := errors.As(err, &te); ok {
			fmt.Println(te.Time, te.Text)
		} else {
			fmt.Println(err)
		}
	}
}
