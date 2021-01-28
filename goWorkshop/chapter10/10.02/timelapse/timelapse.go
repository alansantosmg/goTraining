package timelapse

import (
	"strconv"
	"time"
)

// ElapsedTime uses start and and time to get time duration
// Gets start time and end time type time.
// Return a string with execution time
func ElapsedTime(start time.Time, end time.Time) string {

	// subtract end from start
	elapsed := end.Sub(start)

	// convert time to string
	hours := strconv.Itoa(int(elapsed.Hours()))
	minutes := strconv.Itoa(int(elapsed.Minutes()))
	seconds := strconv.Itoa(int(elapsed.Seconds()))

	// return execution time in a human being format
	return "The total execution time elapsed is: " + hours + "hours and " + minutes + "minutes, " + seconds + "seconds!"

}
