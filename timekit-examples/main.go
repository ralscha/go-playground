package main

import (
	"fmt"
	"github.com/bartmika/timekit"
	"log"
	"time"
)

func main() {
	startOfYearDate := timekit.FirstDayOfThisYear(time.Now)
	fmt.Println(startOfYearDate)

	jsTime := int64(1643082322380)
	goTime := timekit.ParseJavaScriptTime(jsTime)
	fmt.Println(goTime)

	loc := time.Local
	start := time.Date(2022, 1, 7, 1, 0, 0, 0, loc)                   // Jan 7th 2022
	end := time.Date(2022, 1, 10, 1, 0, 0, 0, loc)                    // Jan 10th 2022
	dts := timekit.RangeFromTimeStepper(start, end, 0, 0, 1, 0, 0, 0) // Step by day.
	fmt.Println(dts)

	loc = time.UTC                                 // closure can be used if necessary
	start = time.Date(2022, 1, 7, 1, 0, 0, 0, loc) // Jan 7th 2022
	end = time.Date(2022, 1, 10, 1, 0, 0, 0, loc)  // Jan 10th 2022
	ts := timekit.NewTimeStepper(start, end, 0, 0, 1, 0, 0, 0)

	var actual time.Time
	running := true
	for running {
		// Get the value we are on in the timestepper.
		actual = ts.Get()

		log.Println(actual) // For debugging purposes only.

		// Run our timestepper to get our next value.
		ts.Next()

		running = ts.Done() == false
	}
}
