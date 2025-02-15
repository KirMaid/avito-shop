package http

import "errors"

const StatusError = "errors"

var ErrInvalidAccessToken = errors.New("invalid auth token")
var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrUsernameNotFoundInContext = errors.New("username not found in context")
var ErrUsernameInvalidFormat = errors.New("username invalid format")
