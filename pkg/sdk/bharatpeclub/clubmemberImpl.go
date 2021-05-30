package payout

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

func (this Accessor) ClubMember(merchantID int, clientName string) (bool, error) {

	var serviceRes serviceResClubMember
	var serviceReq reqClubMember
	hmacReq := make(map[string]interface{}, 0)
	url := fmt.Sprintf("%s/%s/%s", this.IPPort, routeBharatpeClud, routeClubMember)

	serviceReq = reqClubMember{
		MerchantID: merchantID,
	}
	j, err := json.Marshal(serviceReq)
	if err != nil {
		return false, err
	}

	hmacReq["merchant_id"] = strconv.Itoa(merchantID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return false, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateBase64HMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return false, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add(headerHash, fmt.Sprintf("%s", hmac))
	httpReq.Header.Add(headerClientName, clientName)
	fmt.Printf("http Req %v url %v", httpReq, url)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if 200 != resp.StatusCode {
		var errObject errorResClubMember
		err = json.Unmarshal(body, &errObject)
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("%s", body)
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return false, err
	}
	return serviceRes.Eligible, nil
}
