package constants

import "net/http"

type (
	InternalCode int
	ServerCode   int
	MsgBase      string
)

const (
	Success       InternalCode = 2001 // Success
	ParamInvalid  InternalCode = 2003 // is invalid
	TimeInvalid   InternalCode = 2004
	QueryNotFound InternalCode = 2005 // query not found
	UuidInvalid   InternalCode = 2006 // query not found

	// token
	InvalidToken InternalCode = 3001
	InvalidOTP   InternalCode = 3002
	AcceptDenied InternalCode = 3003

	// otp
	CantSendEmailOtp InternalCode = 4003

	// authen
	AuthFailed InternalCode = 4005
	// Register Code
	UserHasExists InternalCode = 5001 // user has already registered

	// Err Login
	OtpNotExists     InternalCode = 6009
	UserOtpNotExists InternalCode = 6008

	// Two Factor Authentication
	TwoFactorAuthSetupFailed  InternalCode = 8001
	TwoFactorAuthVerifyFailed InternalCode = 8002

	InternalServerErr InternalCode = 9999
	DatabaseErr       InternalCode = 9998

	gw               ServerCode = 1000
	userSer          ServerCode = 1001
	MarkerDataSer    ServerCode = 1002
	notificationSer  ServerCode = 1003
	portfolioSer     ServerCode = 1004
	whaleTrackingSer ServerCode = 1005
	consumerSer      ServerCode = 1006
	orchestratorSer  ServerCode = 1007
	jobSer           ServerCode = 1008

	DatabaseNotFound MsgBase = "entity not found"
)

// message for Client
var Msg = map[InternalCode]string{
	Success:          "success",
	ParamInvalid:     "email is invalid",
	InvalidToken:     "token is invalid",
	InvalidOTP:       "Otp error",
	CantSendEmailOtp: "Failed to send email OTP",

	UserHasExists: "user has already registered",

	OtpNotExists:     "otp exists but not registered",
	UserOtpNotExists: "user otp not exists",
	AuthFailed:       "authentication failed",

	// Two Factor Authentication
	TwoFactorAuthSetupFailed:  "two factor authentication setup failed",
	TwoFactorAuthVerifyFailed: "two factor authentication verify failed",

	InternalServerErr: "internal server error",
	DatabaseErr:       "internal server error",
}

var HttpCode = map[InternalCode]int{
	Success:          http.StatusOK,
	ParamInvalid:     http.StatusBadRequest,
	InvalidToken:     http.StatusUnauthorized,
	InvalidOTP:       http.StatusUnauthorized,
	CantSendEmailOtp: http.StatusInternalServerError,

	UserHasExists: http.StatusConflict,

	OtpNotExists:     http.StatusUnauthorized,
	UserOtpNotExists: http.StatusUnauthorized,
	AuthFailed:       http.StatusUnauthorized,

	// Two Factor Authentication
	TwoFactorAuthSetupFailed:  http.StatusInternalServerError,
	TwoFactorAuthVerifyFailed: http.StatusInternalServerError,

	InternalServerErr: http.StatusInternalServerError,
	DatabaseErr:       http.StatusInternalServerError,
}
