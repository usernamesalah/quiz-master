package constants

import (
	"strconv"
)

type MessageEnum int32

const (
	Error_UNKOWN_ERROR MessageEnum = iota + 1
	Error_INVALID_LOGIN_INFO
	Error_USER_ALREADY_EXIST
	Error_REGISTER_FAILED
	Error_UNABLE_TO_CREATE_OTP
	Error_INVALID_OTP_CODE
	Error_USER_BLOCKED
	Error_USER_NOT_AUTHORIZED
)

var MessageEnum_MessageName = map[int32]string{
	1: "Something went wrong",
	2: "Invalid Username or Password",
	3: "User already exist",
	4: "Registration failed",
	5: "Can not create OTP for your email, try with another email",
	6: "Invalid OTP Code",
	7: "This user is has been blocked, try another username",
	8: "You Are Not Authorized to access this page",
}

// String : return message string from enum
func (x MessageEnum) String() string {
	return enumToStr(MessageEnum_MessageName, int32(x))
}

func (x MessageEnum) Int() int {
	return int(x)
}

func enumToStr(m map[int32]string, v int32) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}
