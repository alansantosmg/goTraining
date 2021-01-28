package sumint

import "testing"

// TestSum test sum operations
func TestSum(t *testing.T) {

	//Verify msg function, prints expected and result values
	verifyMsg := func(t *testing.T, expected, result int) {
		t.Helper()
		if result != expected {
			t.Errorf("\nExpected: '%d',\nResult: '%d'", expected, result)
		}
	}

	//Test case has 2 params: description and test function.
	t.Run("Basic sum test", func(t *testing.T) {
		expected := sum(1, 1)
		result := 2
		verifyMsg(t, expected, result)
	})

	t.Run("Basic sum test 2", func(t *testing.T) {
		expected := sum(-1, 1)
		result := 0
		verifyMsg(t, expected, result)
	})

	t.Run("Basic sum test 2", func(t *testing.T) {
		expected := sum(-1, -1)
		result := -2
		verifyMsg(t, expected, result)
	})

}
