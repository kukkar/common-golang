package lending

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Config
type Config struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	ClientValidator clientvalidator.ClientValidator
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	return nil
}

type allowBankAccountChangeRes struct {
	ResponseCode string                     `json:"responseCode"`
	Status       string                     `json:"status"`
	Success      bool                       `json:"success"`
	Data         allowBankAccountChangeData `json:"data"`
}

type allowBankAccountChangeData struct {
	Message           string `json:"message"`
	BankAccountChange bool   `json:"bankAccountChange"`
}
