package mobilenotification

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/kukkar/common-golang/pkg/logger"
)

//WhatsAppAccessor
type WhatsAppAccessor struct {
	IPPort   string
	Version  string
	UserName string
	Password string
}

func (this WhatsAppAccessor) SendMessage(number string, message string) (*SendOTPRes, error) {

	parsedURL := this.IPPort
	var output SendOTPRes
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("method", "SendMessage")
	_ = writer.WriteField("v", "1.1")
	_ = writer.WriteField("msg_type", "text")
	_ = writer.WriteField("userid", this.UserName)
	_ = writer.WriteField("password", this.Password)
	_ = writer.WriteField("send_to", number)
	_ = writer.WriteField("msg", message)
	_ = writer.WriteField("format", "text")

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
		return nil, fmt.Errorf(fmt.Sprintf("whatsapp return with error %s", body))
	}
	logger.Logger.Info(fmt.Sprintf("whatsapp response %s", body))
	output.Status = true
	//	output.Resposne = body
	return &output, nil
}
