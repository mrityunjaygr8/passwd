package utils

import "fmt"

var NOT_FOUND = fmt.Errorf("NOT-FOUND")             // 404
var BAD_REQUEST = fmt.Errorf("BAD-REQUEST")         // 400
var UNAUTHENTICATED = fmt.Errorf("UNAUTHENTICATED") // 403
var UNAUTHORISED = fmt.Errorf("UNAUTHORISED")       // 401
