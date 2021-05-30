package otpLess

type SDK interface {
	CreateIntent() (*CreateIntentRes, error)
	VerifyWTToken(token string) (*VerifyWTTokenRes, error)
}
