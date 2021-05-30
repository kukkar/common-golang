package handshake

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/kukkar/common-golang/pkg/logger"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
}

func (this Accessor) PennyDrop(pennyDropReq PennyDropReq) (*PennyDropRes, error) {

	uniqueID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	uID := strings.ReplaceAll(uniqueID.String(), "-", "")
	if len(uID) > 30 {
		uID = uID[:30]
	}
	payload := strings.NewReader(
		fmt.Sprintf("ifsc=%s&acno=%s&mobile=%s&uniqueno=%s&verifybankaccount",
			pennyDropReq.IFSC, pennyDropReq.AccountNumber,
			pennyDropReq.MobileNumber, uID))
	logger.Logger.Info(fmt.Sprintf("hand shake payload penny drop %s", payload))
	client := &http.Client{}
	req, err := http.NewRequest("POST", this.IPPort, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	logger.Logger.Info(fmt.Sprintf("body %s", body))
	var output PennyDropRes
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	if output.Response == "failed" {
		if output.AlertMessage == ServiceErrorWrongIFSCCode || output.AlertMessage == ServiceErrorWrongBankDetail {
			return nil, ErrorWrongDetail
		} else if output.AlertMessage == ServiceErrorVerificationLimitExceeded {
			return nil, ErrorLimitExceed
		} else if output.AlertMessage == ServiceErrorAccoutnAlreadyRegisterd {
			return nil, ErrorAccountAlreadyRegistered
		}
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &output, nil
}

func (this Accessor) GetBeneficiary(getBeneficiaryReq GetBeneficiaryReq) (*GetBeneficiaryResponse, error) {

	payload := strings.NewReader(
		fmt.Sprintf("ifsc=%s&acno=%s&mobile=%s",
			getBeneficiaryReq.IFSC, getBeneficiaryReq.AccountNumber,
			getBeneficiaryReq.Mobile))
	logger.Logger.Info(fmt.Sprintf("hand shake payload penny drop %v", payload))
	client := &http.Client{}
	req, err := http.NewRequest("POST", this.IPPort, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	var output GetBeneficiaryResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (this Accessor) GetCategoryTree() (*GetCategoryTree, error) {

	payload := strings.NewReader(
		fmt.Sprintf("key=%s", "newcategory"))
	logger.Logger.Info(fmt.Sprintf("hand shake payload penny drop %v", payload))
	client := &http.Client{}
	req, err := http.NewRequest("POST", this.IPPort, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	var output GetCategoryTree
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
