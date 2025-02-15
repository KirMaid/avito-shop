package usecases

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")
var ErrFailedToGetUser = errors.New("failed to get user")
