package otpLess

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Config
type Config struct {
	IPPort          string
	Version         string
	ClientId        string
	ClientSecret    string
	ClientValidator clientvalidator.ClientValidator
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	if this.ClientId == "" {
		return fmt.Errorf("ClientId can not be empty")
	}
	if this.ClientSecret == "" {
		return fmt.Errorf("ClientSecret can not be empty")
	}
	if this.ClientValidator == nil {
		return fmt.Errorf("client validator required")
	}
	return nil
}

type CreateIntentRes struct {
	Status string `json:"status"`
	Intent string `json:"url"`
}

type VerifyTokenOtplLessReq struct {
	Token string `json:"token"`
}

type VerifyWTTokenRes struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
