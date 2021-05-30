package ifscfinder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IFSCFinder struct {
	URL string
}

func (this IFSCFinder) GetIFSC(ifscCode string) (*IFSCRes, error) {

	var serviceRes IFSCRes
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", this.URL, ifscCode)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf(fmt.Sprintf("not found %v", body))
	}
	fmt.Printf("body %s", body)
	err = json.Unmarshal(body, &serviceRes)
	if err != nil {
		return nil, err
	}
	return &serviceRes, nil
}
