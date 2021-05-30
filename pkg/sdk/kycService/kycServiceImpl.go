package kycService

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

func (this Accessor) GetPanHolderName(panCardNumber string, merchantID int, clientName string) (string, error) {

	var serviceRes GetPanHolderNameRes

	serviceReq := GetPanHolderNameReq{
		PanNumber:  panCardNumber,
		MerchantId: merchantID,
		UserType:   DefaultUserType,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return "", err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/internal/pan-verify")

	hmacReq["panNumber"] = serviceReq.PanNumber
	hmacReq["identifier"] = serviceReq.MerchantId
	hmacReq["userType"] = serviceReq.UserType
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return "", err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(HeaderClientName, clientName)
	client := &http.Client{
		Timeout: PanVerifyTimeout * time.Second,
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if 200 != resp.StatusCode {
		return "", fmt.Errorf("something went Wrong with kyc service body :-  %s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return "", err
	}
	if !serviceRes.Status {
		return "", fmt.Errorf("someting went wrong %s", body)
	}
	return serviceRes.Data.Name, nil
}
func (this Accessor) UpdateMerchantKyc(pan string, token string, clientName string) error {

	var serviceRes UpdateKYCRes

	serviceReq := UpdateKYCReq{
		PanNumber: pan,
		DocType:   KycPanType,
		Source:    DefaultUserType,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/process-kyc")

	hmacReq["docType"] = serviceReq.DocType
	hmacReq["panNumber"] = serviceReq.PanNumber
	hmacReq["source"] = serviceReq.Source

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(HeaderClientName, clientName)
	httpReq.Header.Add(HeaderToken, token)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("something went Wrong with kyc service body :-  %s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return err
	}
	return nil
}
func (this Accessor) NameMatch(name1 string, name2 string, clientName string) (*int, error) {
	var matchPercentage int
	if name1 == name2 {
		matchPercentage = 100
		return &matchPercentage, nil
	}

	var serviceRes NameMatchRes

	serviceReq := NameMatchReq{
		Name1: name1,
		Name2: name2,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/internal/name-match")

	hmacReq["name1"] = serviceReq.Name1
	hmacReq["name2"] = serviceReq.Name2
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
	httpReq.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(HeaderClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{
		Timeout: PanVerifyTimeout * time.Second,
	}
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
		return nil, fmt.Errorf("something went Wrong with kyc service body :-  %s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	if !serviceRes.Status {
		return nil, fmt.Errorf("someting went wrong %s", body)
	}
	matchPercentage = int(serviceRes.NameMatchPercentage * 100)
	return &matchPercentage, nil
}

func (this Accessor) GetPan(merchantId int, clientName string) (*GetPanData, error) {
	var serviceRes GetPanRes

	serviceReq := GetPanReq{
		MerchantId: merchantId,
		UserType:   DefaultUserType,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "api/v1/internal/pan-details")

	hmacReq["identifier"] = merchantId
	hmacReq["userType"] = DefaultUserType
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
	httpReq.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(HeaderClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{
		Timeout: PanVerifyTimeout * time.Second,
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 404 == resp.StatusCode {
		return nil, nil
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("something went Wrong with kyc service body :-  %s", body)
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	if !serviceRes.Status {
		return nil, fmt.Errorf("someting went wrong %s", body)
	}
	return &serviceRes.Data, nil
}
