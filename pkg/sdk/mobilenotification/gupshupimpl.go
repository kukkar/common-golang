package mobilenotification

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/kukkar/common-golang/pkg/logger"
)

//GupShupAccessor
type GupShupAccessor struct {
	IPPort   string
	Version  string
	UserName string
	Password string
}

func (this GupShupAccessor) SendMessage(number string, message string) (*SendOTPRes, error) {

	//parsedURL := fmt.Sprintf("%s?dest=%s&msg=%s&uname=%s&pwd=%s&send=BHRTPE", this.IPPort, number, url.QueryEscape(merchant), this.UserName, this.Password)
	parsedURL := this.IPPort
	method := "POST"
	var output SendOTPRes
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("method", "sendMessage")
	_ = writer.WriteField("version", "1.0")
	_ = writer.WriteField("msg_type", "text")
	_ = writer.WriteField("userid", this.UserName)
	_ = writer.WriteField("password", this.Password)
	_ = writer.WriteField("auth_scheme", "PLAIN")
	_ = writer.WriteField("send_to", number)
	_ = writer.WriteField("msg", url.QueryEscape(message))
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, parsedURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(fmt.Sprintf("%s", body), "success") {
		return nil, fmt.Errorf(fmt.Sprintf("gupshup return with error %s", body))
	}
	logger.Logger.Info(fmt.Sprintf("gupshup response %s mobile %v ", body, number))
	output.Status = true
	output.Resposne = string(body)
	return &output, nil
}
