package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	UnknownErrCode       = gcode.New(10000, "Undefined Error", nil)          // Undefined Error
	PermsErrCode         = gcode.New(10001, "Insufficient Permissions", nil) // Insufficient Permissions
	DataNotFoundErrCode  = gcode.New(10002, "Data Not Found", nil)           // Data Not Found
	DataExistErrCode     = gcode.New(10003, "Data Already Exists", nil)      // Data Already Exists
	DataInvalidErrCode   = gcode.New(10004, "Invalid Data", nil)             // Invalid Data
	ParamsInvalidErrCode = gcode.New(10005, "Invalid Parameters", nil)       // Invalid Parameters
	AuthExpiredErrCode   = gcode.New(10006, "Authentication Expired", nil)   // Authentication Expired
	SystemErrCode        = gcode.New(10007, "System Error", nil)             // System Error
	NotLoggedInErrCode   = gcode.New(401, "Not Logged In", nil)              // Not Logged In
	UnauthorizedErrCode  = gcode.New(403, "Unauthorized", nil)               // Unauthorized
)
