package errors

import "errors"

var DbConnectionError error = errors.New("Failed to Connecion Mysql DB")