package payout

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	clientvalidator clientvalidator.ClientValidator
}

func (this Accessor) PennyDrop(req PennyDropReq, clientName string) (*PennyDropRes, error) {

	var output PennyDropRes
	var serviceRes pennyDropServiceRes
	var serviceReq pennyDropServiceReq
	var hmacReq map[string]interface{}
	url := fmt.Sprintf("%s/%s/%s", this.IPPort, pennyDropRoute, this.Version)
	copier.Copy(&serviceReq, &req)
	logger.Logger.Info(fmt.Sprintf("penny drop v3  request %v", serviceReq))
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(PennyDropHeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(PennyDropHeaderClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	if serviceRes.ResponseCode != "100" {
		if serviceRes.Message == ServiceErrorWrongIFSCCode || serviceRes.Message == ServiceErrorWrongBankDetail {
			return nil, ErrorWrongDetail
		} else if serviceRes.Message == ServiceErrorVerificationLimitExceeded {
			return nil, ErrorLimitExceed
		} else if serviceRes.Message == ServiceErrorAccoutnAlreadyRegisterd {
			return nil, ErrorAccountAlreadyRegistered
		}
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	output = PennyDropRes{
		BeneficiaryName: serviceRes.BeneficiaryData.BeneficiaryName,
		IFSC:            serviceRes.BeneficiaryData.IFSC,
		AccountNumber:   serviceRes.BeneficiaryData.AccountNumber,
	}
	return &output, nil
}

func (this Accessor) EarlySettlement(req ReqEarlySettlement,
	clientName string) ([]int, error) {

	var serviceRes serviceResEarlySettlement
	var serviceReq serviceReqEarlySettlement
	var hmacReq map[string]interface{}
	url := fmt.Sprintf("%s/%s", this.IPPort, earlyPayoutRoute)
	serviceReq = serviceReqEarlySettlement{
		MerchantID:      req.MerchantID,
		MerchantStoreID: req.MerchantStoreID,
		Amount:          req.Amount,
		Token:           req.Token,
	}
	sMID := strconv.Itoa(req.MerchantID)
	//	sAmount := fmt.Sprintf("%s", req.Amount)
	// var sMSID string
	// if req.MerchantStoreID != nil {
	// 	sMSID = strconv.Itoa((*req.MerchantStoreID))
	// }

	tmpServiceReq := tmpservReqEarlySett{
		MerchantID: sMID,
		//	Amount:     &sAmount,
		//	MerchantStoreID: &sMSID,
		Token: req.Token,
	}
	logger.Logger.Info(fmt.Sprintf("early settlement request %v %s", serviceReq, url))
	j, err := json.Marshal(tmpServiceReq)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(PennyDropHeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(PennyDropHeaderClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	fmt.Printf("service res payout %v", serviceRes)
	if serviceRes.StatusCode == "420" {
		return nil, ErrInsufficientBalance
	}
	if serviceRes.StatusCode != "200" ||
		len(serviceRes.Data.SettlementIDs) == 0 {
		return nil, fmt.Errorf("enable to submit settlement request  %v ", serviceRes)
	}
	return serviceRes.Data.SettlementIDs, nil
}
