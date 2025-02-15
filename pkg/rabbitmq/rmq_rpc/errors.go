package rmqrpc

import "errors"

var ErrTimeout = errors.New("timeout")
var ErrInternalServer = errors.New("internal server error")
var ErrBadHandler = errors.New("unregistered handler")

const Success = "success"
