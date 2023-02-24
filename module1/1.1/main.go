package main

import (
	"fmt"
)


func main()  {
	myArray := [5]string{"I", "am", "stupid", "and","weak"}
	fmt.Println(myArray)
	for index, value := range myArray {
		if value == "stupid" {
			myArray[index] = "smart"
		} else if value == "weak" {
			myArray[index] = "strong"
		} else {
			myArray[index] = value
		}
	}
	fmt.Println(myArray)
}