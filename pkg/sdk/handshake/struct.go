package handshake

import "fmt"

//SDKConfig
type SDKConfig struct {
	IPPort  string
	Version string
}

func (this *SDKConfig) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	return nil
}

type PennyDropReq struct {
	IFSC          string
	AccountNumber string
	MobileNumber  string
}

type PennyDropRes struct {
	Response        string `json:"response"`
	AlertMessage    string `json:"alertmsg"`
	AccountNumber   string `json:"acno"`
	Mobile          string `json:"mobile"`
	IFSC            string `json:"ifsc"`
	Branch          string `json:"branch"`
	BankName        string `json:"bank"`
	BeneficiaryName string `json:"beneficiaryName"`
}

type GetBeneficiaryReq struct {
	Mobile        string
	IFSC          string
	AccountNumber string
	MerchantID    int
}
type GetBeneficiaryResponse struct {
	AccountNumber   string `json:"acno"`
	Mobile          string `json:"mobile"`
	IFSC            string `json:"ifsc"`
	Branch          string `json:"branch"`
	BankName        string `json:"bank"`
	BankCode        string `json:"bank_code"`
	BeneficiaryName string `json:"beneficiaryName"`
}

type GetCategoryTree struct {
	Response string             `json:"response"`
	Data     []CategoryTreeData `json:"data"`
}

type CategoryTreeData struct {
	CategoryName  string            `json:"cat_name"`
	SubCategories []SubCategoryData `json:"sub_categories"`
}

type SubCategoryData struct {
	CategoryName string `json:"cat_name"`
}
