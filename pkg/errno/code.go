package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrClient     = &Errno{Code: 20004, Message: "Refresh Token came from the wrong client."}
	ErrUrl        = &Errno{Code: 20005, Message: "Error Url."}
	ErrUID        = &Errno{Code: 20006, Message: "Request with illegal uid."}

	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUserEmail         = &Errno{Code: 20105, Message: "The email address has been registered."}
	ErrUserPhone         = &Errno{Code: 20106, Message: "The phone number has been registered."}
	ErrUpdatePref        = &Errno{Code: 20107, Message: "Failed to update user preference."}
	ErrParamKey          = &Errno{Code: 20108, Message: "Get with illegal key from gin context"}
	ErrRole              = &Errno{Code: 20109, Message: "The user level is insufficient."}
	ErrLogoutRToken      = &Errno{Code: 20110, Message: "Logout refresh token error."}
	ErrRTokenInvalid     = &Errno{Code: 20111, Message: "The refresh token was invalid."}
)
