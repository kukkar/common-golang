package payout

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

type reqClubMember struct {
	MerchantID int `json:"merchant_id"`
}

type serviceResClubMember struct {
	Success  bool   `json:"success"`
	Message  string `json:"msg"`
	Eligible bool   `json:"eligibile"`
}

type errorResClubMember struct {
	ResponseCode string `json:"responseCode"`
	Message      string `json:"responseMessage"`
	Status       string `json:"status"`
}
