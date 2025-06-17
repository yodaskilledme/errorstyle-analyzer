package a

import (
    "domain"
)

func errorNotLast() (error, int) { // want "error must be last"
    return nil, 0
}

func valid() error {
    const op = "valid"
    return domain.Error{Op: op, Message: "test"}
}

func noOpConst() (int, error) {
    return 0, domain.Error{Op: "foo", Message: "test"} // want "const value must be used"
}

func invalidOpValue() error {
    const op = "invalid_OpValue" // want "operation must be `invalidOpValue` not `invalid_OpValue`"
    return domain.Error{Op: op, Message: "bar"}
}

func wrongOpConstName() error {
    const operation = "wrongOpConstName" // want "operation constant must be named `op` not `operation`"
    return domain.Error{Op: operation, Message: "bar"}
}

func returnsOpError() error {
    const op = "returnsOpError"
    return domain.OpError(op, nil)
}

func returnsFunc() (int, error) {
    return noOpConst()
}
