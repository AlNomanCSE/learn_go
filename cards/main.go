package main

import (
	"fmt"
)

func main() {
	var myFunction func()
	myFunction = func() {
		fmt.Println("This is MyFunction")
	}
	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its an String")
		case func():
			fmt.Println("Its a Function")
		default:
			fmt.Println("This is a function")
		}

	}
	whoAmI(myFunction)
}

// func switchCaseOne(day string) {
// 	switch day {
// 	case "Monday":
// 		fmt.Println("Today is Monday")
// 	case "Tuesday":
// 		fmt.Println("Today is Tuesday.")
// 	case "Wednesday":
// 		fmt.Println("Today is Wednesday.")
// 	case "Thursday":
// 		fmt.Println("Today is Thursday.")
// 	case "Friday":
// 		fmt.Println("Today is Friday.")
// 	case "Saturday", "Sunday":
// 		fmt.Println("It's the weekend!")
// 	default:
// 		fmt.Println("Invalid day.")
// 	}
// }
// func switchCaseTwo(number int) {

// 	switch {
// 	case number > 30:
// 		fmt.Println("Number is Greater then 30")
// 	case number > 20:
// 		fmt.Println("Number is greater than 20.")
// 	case number > 10:
// 		fmt.Println("Number is greater than 10.")
// 	}
// }

// func switchCaseThree(i any) {
// 	switch i.(type) {
// 	case int:
// 		fmt.Println("Integer ")
// 	case string:
// 		fmt.Println("String")
// 	}
// }

// func switchFour() {
// 	hour := time.Now().Hour()
// 	fmt.Println(time.Now().Weekday())

// 	switch {
// 	case hour > 12:
// 		fmt.Println("Good Morning")
// 	case hour < 12:
// 		fmt.Println("Good afternoon")
// 	default:
// 		fmt.Println("Good evening")

// 	}
// }

// func switchDay() {
// 	switch time.Now().Weekday() {
// 	case time.Saturday:
// 		fmt.Println("Today is Saturday")
// 	}

// }
