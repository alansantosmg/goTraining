package main

import (
	"errors"
	"fmt"
)

func validateNumber(i int) error {

	if err := validateFizz(i); err != nil {
		validateBuzz(i)

		if err := validateBuzz(i); err != nil {
			return err
		} else {
			return err
		}
	} else {
		return err
	}

}

func validateBuzz(i int) error {
	if i%5 == 0 {
		return errors.New("Buzz")
	} else {
		return nil
	}
}

func validateFizz(i int) error {
	if i%3 == 0 {
		return errors.New("Fizz")
	} else {
		return nil
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		if err := validateNumber(i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i)
		}
	}
}
