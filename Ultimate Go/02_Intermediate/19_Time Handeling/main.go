package main

import (
	"fmt"
	"time"
)

func main() {
	// Current Local time
	fmt.Println(time.Now())

	// specific time
	specificTime := time.Date(2025, time.June, 01, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific time : ", specificTime)

	// Parsed time
	// format : 2020-05-01 18:03:00 +0000 UTC yy/mm/dd, H:M:Sec std time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	parsedTime1, _ := time.Parse("06-01-02 15-04 -0700", "20-05-01 18-03 +0530") // with date and time
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	// Formatting time :
	tf := time.Now()
	fmt.Println("Formatted time : ", tf.Format("Sunday 06-01-02 04-15")) // return a formatted current time

	// time manupulations
	OneDayLater := tf.Add(time.Hour * 24)
	fmt.Println("One day later :", OneDayLater)
	fmt.Println("One day later day is :", OneDayLater.Weekday())
	fmt.Println("Rounded time : ", tf.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// tf = time.Date(2025, time.June, 8, 14, 16, 40, 00, time.UTC)

	// // convert this to the specif time zone
	// tLocal := tf.In(loc)
	// // Perform rounding
	// roundedTime := tf.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original Time (UTC) : ", tf)
	// fmt.Println("Original Time (Local) : ", tLocal)
	// fmt.Println("Rounded Time (UTC) : ", roundedTime)
	// fmt.Println("Rounded Time (Local) : ", roundedTimeLocal)

	// fmt.Println("Truncated Time :", tf.Truncate(time.Hour))

	// Working on different time zones :
	// locate, _ := time.LoadLocation("America/New_York")

	// // Convert the time to location
	// TINNY := time.Now().In(locate)
	// fmt.Println("New york time is : ", TINNY)

	// Time calculation :
	t1 := time.Date(2024, time.June, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.June, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)

	fmt.Println("Duration : ", duration)

	// Compare times
	fmt.Println("T2 is after ", t2.After(t1))
}
