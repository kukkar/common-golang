package loyalty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (this Accessor) GetLoyaltyPoints(merchantID int, clientName string) (*GetLoyaltyPointsRes, error) {

	var serviceRes serviceResGetLoyaltyPoints
	var serviceReq serviceReqGetLoyalPoints
	hmacReq := make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s/%s", this.IPPort, RouteCommonPrefix, RouteGetPoints)

	serviceReq = serviceReqGetLoyalPoints{
		MerchantID: strconv.Itoa(merchantID),
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return nil, err
	}

	hmacReq["merchant_id"] = strconv.Itoa(merchantID)
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
	httpReq.Header.Add(headerHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(headerClientName, clientName)
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
		var errObject loyaltyServiceErrRes
		err = json.Unmarshal(body, &errObject)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}

	return &GetLoyaltyPointsRes{
		Points: serviceRes.Points,
	}, nil
}

func (this Accessor) RedeemLoyaltyPoints(req RedeemPointsReq, clientName string) error {
	var serviceRes serviceResRedeemPoints
	var serviceReq serviceReqRedeemPoints
	hmacReq := make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s/%s", this.IPPort, RouteCommonPrefix, RouteRedeemPoints)

	serviceReq = serviceReqRedeemPoints{
		MerchantID: strconv.Itoa(req.MerchantID),
		TxnID:      req.TxnID,
		Category:   req.Category,
		TxnType:    req.TxnType,
		Points:     req.Points,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return err
	}

	hmacReq["merchant_id"] = strconv.Itoa(req.MerchantID)
	hmacReq["txn_id"] = req.TxnID
	hmacReq["txn_type"] = req.TxnType
	hmacReq["category"] = req.Category
	hmacReq["points"] = req.Points

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
	httpReq.Header.Add(headerHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(headerClientName, clientName)
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
		var errObject loyaltyServiceErrRes
		err = json.Unmarshal(body, &errObject)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return err
	}
	if !serviceRes.Status {
		return fmt.Errorf("unable to redeem points loyalty return with error %v", serviceRes)
	}
	return nil
}
