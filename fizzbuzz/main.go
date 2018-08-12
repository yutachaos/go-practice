package main

import "fmt"

func fizzBuzz(number int) string{
		if number%3 == 0 && number % 5 == 0 {
			return "FizzBuzz"
		} else if number%3 == 0 {
			return "Fizz"
		} else if number%5 == 0 {
			return "Buzz"
		} else {
			return fmt.Sprintf("%v", number)
		}

}

func main() {
	for i := 1; i <= 100 ; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
