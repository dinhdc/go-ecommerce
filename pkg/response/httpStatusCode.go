package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrInvalidToken     = 30001 // token is invalid
	// register code
	ErrCodeUserHasAlreadyExist = 50001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:             "success",
	ErrCodeParamInvalid:        "Email is invalid",
	ErrInvalidToken:            "Invalid token",
	ErrCodeUserHasAlreadyExist: "user has already register",
}
