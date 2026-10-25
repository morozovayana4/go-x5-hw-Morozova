package main

import (
	"errors"
	"fmt"
)

var ErrRead = errors.New("read failed")

type ReadError struct {
	File string
	Err  error
}

func (e *ReadError) Error() string {
	return fmt.Sprintf("cannot read %s: %v", e.File, e.Err)
}

func (e *ReadError) Unwrap() error {
	return e.Err
}

func (e *ReadError) Is(target error) bool {
	return target == ErrRead
}

func readData() {
	panic("invalid data")
}

func readSafely() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in readSafely:", r)
		}
	}()

	readData()
}

func processData() {
	readData()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()
	cause := errors.New("file not found")
	err := &ReadError{
		File: "data.txt",
		Err:  cause,
	}
	fmt.Println(err)
	fmt.Println(errors.Unwrap(err))
	fmt.Println(errors.Is(err, ErrRead))
	fmt.Println(errors.Is(err, cause))
	readSafely()
	processData()
}
