package services

import "errors"

var ErrorNotFound = errors.New("not found")
var UserBadRequest = errors.New("bad request")
var ServerUnavailable = errors.New("service unavailable")
