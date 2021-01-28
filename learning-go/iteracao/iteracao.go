// Package iteracao defines a for fucntion
// that repeat a string 5 times
package iteracao

// Repeat does reapeat a single string n times
func Repeat(myString string, repeatTimes int) string {
	var result string
	for len(result) < repeatTimes { // len(myvar) gets the lenght of myvar
		result += myString
	}
	return result
}
