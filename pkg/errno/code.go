package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "Param error"}
	ErrDataIsNotExist   = &Errno{Code: 10004, Message: "该数据不存在"}

	ErrValidation    = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase      = &Errno{Code: 20002, Message: "Database error."}
	ErrToken         = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrGetUploadFile = &Errno{Code: 20004, Message: "Error get file while geting file."}
	ErrUploadingFile = &Errno{Code: 20005, Message: "Error get file while uploading file."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	// course errors
	ErrCourseNotFound   = &Errno{Code: 20201, Message: "The course was not found."}
	ErrCourseCreateFail = &Errno{Code: 20202, Message: "The course create fail."}

	// video errors
	ErrVideoNotFound   = &Errno{Code: 20401, Message: "The video create fail."}
	ErrVideoCreateFail = &Errno{Code: 20402, Message: "The video create fail."}

	// topic errors
	ErrNoRightEdit = &Errno{Code: 20501, Message: "你没有编辑该文章的权限哦~"}
)
