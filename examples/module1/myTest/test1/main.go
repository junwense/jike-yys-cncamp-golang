package main

import "fmt"

func main() {
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	for index, _ := range mySlice {
		switch index {
		case 2:
			mySlice[index] = "smart"
		case 4:
			mySlice[index] = "strong"
		default:
			break
		}
	}

	fmt.Printf("mySlice %+v\n", mySlice)
}
