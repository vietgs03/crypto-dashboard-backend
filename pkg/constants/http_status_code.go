package constants

type (
	InternalCode int
	ServerCode   int
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
)

// message for Client
var Msg = map[InternalCode]string{
	Success:          "success",
	ParamInvalid:     "Email is invalid",
	InvalidToken:     "token is invalid",
	InvalidOTP:       "Otp error",
	CantSendEmailOtp: "Failed to send email OTP",

	UserHasExists: "user has already registered",

	OtpNotExists:     "OTP exists but not registered",
	UserOtpNotExists: "User OTP not exists",
	AuthFailed:       "Authentication failed",

	// Two Factor Authentication
	TwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	TwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",

	InternalServerErr: "Internal Server Error",
	DatabaseErr:       "Internal Server Error",
}
