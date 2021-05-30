package swipe

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Config
type Config struct {
	IPPort          string
	Version         string
	ClientValidator clientvalidator.ClientValidator
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	if this.ClientValidator == nil {
		return fmt.Errorf("client validator required")
	}
	return nil
}

type CreateAccountReq struct {
	MerchantID int `json:"merchantId"`
}

type CreateAccountRes struct {
	Success    bool   `json:"success"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}
