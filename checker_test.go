package checker

import "testing"

func TestAssert_IsTrue_NilReturned(t *testing.T) {
	err := assert(1 > 0, "this message is not shown")
	if err != nil {
		t.Error("Error when condition is true")
	}
}

func TestAssert_IsFalse_ErrorReturned(t *testing.T) {
	err := assert(0 > 1, "0 is not less then 1")
	if err == nil {
		t.Error("Error is not returned then condition is false")
	}
}

func TestAssert_PassErrorText_ErrorShouldHaveTheSameText(t *testing.T) {
	errorMessage := "Error Message"
	err := assert(0 > 1, errorMessage)
	if err.Error() != errorMessage {
		t.Error("Error messages are different")
	}
}

func TestRequire_IsFalse_PanicCalled(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic haven't been called")
		}
	}()

	require(0 > 1, "error message")
}

func TestRequire_IsTrue_PanicNotCalled(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Panic called, but shouldn't")
		}
	}()

	require(1 > 0, "error message")
}

func TestRequire_PassErrorText_PanicShouldHaveTheSameText(t *testing.T) {
	message := "error message"
	defer func(msg string) {
		r := recover()
		if r == nil {
			t.Error("Panic haven't been called")
		} else if r.(error).Error() != msg {
			t.Error("Error messages are different")
		}
	}(message)

	require(0 > 1, message)
}
