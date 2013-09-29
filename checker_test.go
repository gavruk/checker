package checker

import "testing"

func TestRequire_IsTrue_NilReturned(t *testing.T) {
    err := require(1 > 0, "this message is not shown")
    if err != nil {
        t.Error("Error when condition is true")
    }
}

func TestRequire_IsFalse_ErrorReturned(t *testing.T) {
    err := require(0 > 1, "0 is not less then 1")
    if err == nil {
        t.Error("Error is not returned then condition is false")
    }
}

func TestRequire_PassErrorText_ErrorShouldHaveTheSameText(t *testing.T) {
    errorMessage := "Error Message"
    err := require(0 > 1, errorMessage)
    if err.Error() != errorMessage {
        t.Error("Error messages are different")
    }
}
