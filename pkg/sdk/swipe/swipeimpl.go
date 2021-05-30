package swipe

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
	clientvalidator clientvalidator.ClientValidator
}

func (this Accessor) CreateAccount(request CreateAccountReq,
	clientName string) (*CreateAccountRes, error) {

	var serviceRes CreateAccountRes
	hmacReq := make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s", this.IPPort, "internal/merchant/createAccount")
	hmacReq["merchantId"] = strconv.Itoa(request.MerchantID)
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
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	req.Header.Add(HeaderHash, fmt.Sprintf("%s", hmac))
	req.Header.Add(HeaderClientName, clientName)
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

	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	if serviceRes.StatusCode != "200" {
		return nil, fmt.Errorf("%s %w", body, ErrorServiceFailed)
	}
	return &serviceRes, nil
}
