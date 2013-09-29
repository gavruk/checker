package checker

import "errors"

func require(condition bool, message string) error {
    if !condition {
        return errors.New(message)
    }
    return nil
}
