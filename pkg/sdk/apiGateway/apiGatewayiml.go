package apiGateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//Accessor
type Accessor struct {
	IPPort          string
	Version         string
	clientvalidator clientvalidator.ClientValidator
}

func (this Accessor) StartTxn(request StartTxnReq,
	clientName string) (*StartTxnRes, error) {

	var serviceRes StartTxnRes
	hmacReq := make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "collect/startTxn")
	hmacReq["merchantId"] = strconv.Itoa(request.MerchantID)
	hmacReq["orderId"] = request.OrderID
	hmacReq["amount"] = request.Amount
	hmacReq["mid"] = request.Mid
	serviceReq, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(serviceReq))
	c := context.TODO()
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err, hmac))
	}
	req.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	req.Header.Add(HeaderMID, request.Mid)
	fmt.Printf("http Req %v url %v", req, url)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if 200 != res.StatusCode {
		return nil, fmt.Errorf(fmt.Sprintf("body %s status code %v", body, res.StatusCode))
	}
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	logger.Info("OrderQR StartTxn Res For OrderID:", request.OrderID, serviceRes)
	if serviceRes.ResponseCode == "401" {
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &serviceRes, nil
}
func (this Accessor) CheckTxn(request CheckTxnReq,
	clientName string) (*CheckTxnRes, error) {

	var serviceRes CheckTxnRes
	hmacReq := make(map[string]interface{}, 0)
	hmacReq["orderId"] = request.OrderID
	url := fmt.Sprintf("%s/%s", this.IPPort, "collect/checkTxn?orderId="+request.OrderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}
	req.Header.Add(HeaderMID, request.Mid)
	req.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	logger.Info("OrderQR CheckTxn Res For OrderID:", request.OrderID, serviceRes)
	if serviceRes.ResponseCode == "401" {
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &serviceRes, nil
}
