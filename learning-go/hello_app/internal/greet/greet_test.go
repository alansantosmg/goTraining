/*
/*
The test filename: some-name_test.go  (the _test fragment is mandadory)
The test file needs to be hosted at tested package DIR
Run go test in the package dir
*/

package greet //It will be the same as package to be tested

import "testing" // Import the testing package

func TestOla(t *testing.T) {
	// Function name needs to start with  Test (first letter uppercase
	// Function mame examples: TestOla TestMyalternatives TestSomething
	// For testing always use param:  t *testing.T

	verifyCorrectMsg := func(t *testing.T, result, expected string) {

		// Errorf is formated error with variables
		t.Helper()
		if result != expected {
			t.Errorf("\nResult: '%s', \nExpected: '%s'", result, expected)
		}
	}

	t.Run("Says Hello to people", func(t *testing.T) {
		result := Hello("Alan", "") // tested function
		expected := "Hello Alan"    // result expected for tested function
		verifyCorrectMsg(t, result, expected)
	})

	t.Run("Says Hello Dude if nothing passed by", func(t *testing.T) {
		result := Hello("", "")  // tested function
		expected := "Hello Dude" // result expected for tested function
		verifyCorrectMsg(t, result, expected)

	})

	t.Run("Says spanish", func(t *testing.T) {
		result := Hello("Elodie", "espanhol") // tested function
		expected := "Hola Elodie"             // result expected for tested function
		verifyCorrectMsg(t, result, expected)

	})

	t.Run("Says french", func(t *testing.T) {
		result := Hello("Elodie", "francês") // tested function
		expected := "Bonjour Elodie"         // result expected for tested function
		verifyCorrectMsg(t, result, expected)

	})

}
