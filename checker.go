package checker

import "errors"

func assert(condition bool, message string) error {
	if !condition {
		return errors.New(message)
	}
	return nil
}

func require(condition bool, message string) {
	if !condition {
		err := errors.New(message)
		panic(err)
	}
}
