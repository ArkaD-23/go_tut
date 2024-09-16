package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Time type: %T\n", currentTime)

	formattedTime := currentTime.Format("15:04:05 PM 02/01/2006 Monday")
	fmt.Println("Formatted time:", formattedTime)

	//corverting date format to time format
	dateStr := "16/09/2024"
	layoutStr := "02/01/2006"
	convertedStr, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Converted time:", convertedStr)
}
