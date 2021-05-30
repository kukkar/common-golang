package kyc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (this Accessor) GetGSTDetail(merchantID int, clientName string) (*GSTDetail, error) {

	var output GSTDetail
	var serviceRes gstDetailServiceRes
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s?merchantId=%d", this.IPPort, "kyc/document/gst/details", merchantID)

	hmacReq["merchantId"] = merchantID
	httpReq, err := http.NewRequest("GET", url, nil)
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
		Timeout: 5 * time.Second,
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
	if serviceRes.StatusCode != "200" {
		return nil, fmt.Errorf("someting went wrong %s", body)
	}
	output = GSTDetail{
		GSTNumber: serviceRes.Data.GSTNumber,
		Status:    serviceRes.Data.Status,
	}
	return &output, nil
}

func (this Accessor) GetPanDetail(merchantID int, clientName string) (*PANDetail, error) {

	var output PANDetail
	var serviceRes panServiceRes
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s?merchantId=%d", this.IPPort, "internal/kyc/pancard/details", merchantID)

	hmacReq["merchantId"] = merchantID
	httpReq, err := http.NewRequest("GET", url, nil)
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
		Timeout: 5 * time.Second,
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
	if serviceRes.StatusCode != "200" {
		return nil, fmt.Errorf("someting went wrong %s", body)
	}
	output = PANDetail{
		PANNumber: serviceRes.Data.DocumentNumber,
		Status:    serviceRes.Data.Status,
	}
	return &output, nil
}

func (this Accessor) PanFetchNVerify(panCardNumber string, panName string,
	clientName string) (*PanDataVerfiy, error) {

	var output PanDataVerfiy
	var serviceRes PanFetchRes
	serviceReq := PanFetchDataReq{
		PanNumber: panCardNumber,
		PanName:   panName,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "internal/kyc/pan/name/match")

	hmacReq["pan_card_number"] = panCardNumber
	hmacReq["name_to_match"] = panName
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
	if serviceRes.StatusCode != "200" {
		return nil, fmt.Errorf("someting went wrong %s", body)
	}
	output = PanDataVerfiy{
		NameMatchPercent:  serviceRes.Data.NameMatchPercent,
		PanCardHolderName: serviceRes.Data.PanCardHolderName,
		ProofType:         serviceRes.Data.ProofType,
	}
	return &output, nil
}

func (this Accessor) SubmitKycDoc(req SubmitKycDocReq, clientName string) error {

	var serviceRes submitKycDocServiceRes
	serviceReq := submitKycDocServiceReq{
		ProofNumber:     req.ProofNumber,
		DocType:         req.DocType,
		MerchantID:      req.MerchantID,
		MerchantStoreID: req.MerchantStoreID,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return err
	}
	var hmacReq = make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "internal/kyc/proof")

	hmacReq["merchant_id"] = strconv.Itoa(req.MerchantID)
	if req.MerchantStoreID != nil {
		hmacReq["merchant_store_id"] = strconv.Itoa(*req.MerchantStoreID)
	} else {
		hmacReq["merchant_store_id"] = nil
	}

	hmacReq["proof_number"] = req.ProofNumber
	hmacReq["app_name"] = req.DocType
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
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{
		Timeout: PanVerifyTimeout * time.Second,
	}
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
	if serviceRes.StatusCode != "200" {
		return fmt.Errorf("someting went wrong %s", body)
	}
	return nil
}
