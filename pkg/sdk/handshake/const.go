package handshake

import "errors"

const (
	ServiceErrorAccountVerficationLimit       = "Account verification limit execeed."
	ServiceErrorAlreadyCompletedVerifyAccount = "You have already completed acountverification process with other account! Please try again."
	ServiceErrorAccoutnAlreadyRegisterd       = "this account already linked with another mobile number"
	ServiceErrorVerificationLimitExceeded     = "Account verification limit execeed."
	ServiceErrorWrongIFSCCode                 = "Wrong IFSC code"
	ServiceErrorWrongBankDetail               = "Sorry. Your bank account could not be verified at this moment. Please try again or use different bank account."
)

var ErrorLimitExceed = errors.New("Account verification limit execeed.")
var ErrorAccountVerificaitonAlreadyDone = errors.New("You have already completed acountverification process with other account! Please try again.")
var ErrorAccountAlreadyRegistered = errors.New("this account already linked with another mobile number")
var ErrorWrongDetail = errors.New("Wrong BankDetails")
var ErrorServiceFailed = errors.New("penny drop service return with error")
