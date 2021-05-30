package pennydrop

import "fmt"

//Config
type Config struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
}

func (this *Config) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	if this.SuperKey == "" {
		return fmt.Errorf("Super Key can not be empty")
	}
	return nil
}

type PennyDropReq struct {
	AccountNumber string
	IFSC          string
	ClientName    string
}

type pennyDropServiceReq struct {
	AccountNumber string `json:"accountNo"`
	IFSC          string `json:"ifsc"`
}

type pennyDropRes struct {
	ResponseCode string           `json:"responseCode"`
	Message      string           `json:"message"`
	Status       string           `json:"status"`
	Data         pennyDropResData ` json:"data"`
}

type pennyDropResData struct {
	Ifsc            string `json:"ifsc"`
	BeneficiaryName string `json:"beneficiaryName"`
	AccountNo       string `json:"accountNo"`
}

type AccountInfo struct {
	Ifsc            string
	BeneficiaryName string
	AccountNo       string
}
