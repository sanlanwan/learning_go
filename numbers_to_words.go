// Write a program to express a 2 digit number in number system wordings

package main

import (
	"fmt"
)

var teen_words [7]string
var ty_words []string

func init_vars() {
	teen_words = [7]string{"Thir", "Four", "Fif", "Six", "Seven", "Eigh", "Nin"}
	ty_words = append(append(ty_words, "Twen"), teen_words[:]...)
}

func main() {
	fmt.Print("Enter a 2 digit number: ")
	var number int
	result := "Invalid number"
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println("Invalid input; not int;")
		return
	}
	if !_validate_input(number) {
		return
	}
	init_vars()
	if number >= 0 && number <= 9 {
		result = get_word_of_single_digit(number)
	} else if number <= 99 {
		result = get_word_for_two_digit(number)
	}
	fmt.Printf("Answer for %d is %s", number, result)
}

func _validate_input(input int) bool {
	if _validate_number_of_digits(input) == false {
		fmt.Println("Only 2 digit number allowed")
		return false
	}
	return true
}

func _validate_number_of_digits(input int) bool {
	if input > 99 || input < 0 {
		return false
	}
	return true
}

func get_word_of_single_digit(number int) string {
	words := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	for _number := 0; _number <= len(words); _number++ {
		if _number == number {
			return words[_number]
		}
	}
	return "Number less than 0"
}

func _get_word_for_teens(number int) string {
	words := []string{"Thir", "Four", "Fif", "Six", "Seven", "Eigh", "Nin"}
	for _number := 0; _number <= len(words); _number++ {
		if _number+3 == number-10 {
			return words[_number] + "teen"
		}
	}
	return "Invalid number"
}

func _get_word_for_beyond_teens(number int) string {
	words := []string{"Twen", "Thir", "Four", "Fif", "Six", "Seven", "Eigh", "Nin"}
	for _number_var := 0; _number_var <= len(words); _number_var++ {
		if (_number_var + 2) == number/10 {
			if (number % 10) == 0 {
				return words[_number_var] + "ty"
			} else {
				return words[_number_var] + "ty" + get_word_of_single_digit(number%10)
			}
		}
	}
	return "Invalid number"
}

func get_word_for_two_digit(number int) string {
	if number >= 10 && number <= 12 {
		words := []string{"Ten", "Eleven", "Twelve"}
		return words[number-10]
	} else if number > 12 && number <= 19 {
		return _get_word_for_teens(number)
	} else if number <= 99 {
		return _get_word_for_beyond_teens(number)
	}
	return "Number greater than 99"
}
