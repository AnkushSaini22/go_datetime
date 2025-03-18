package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()

	// Print the current date and time
	fmt.Println("Current Date & Time:", currentTime.Format("2006-01-02 15:04:05"))
}
