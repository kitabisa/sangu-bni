package bni

import (
	"encoding/json"
	"io"
	"strings"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v *map[string]interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseURL + path

	return gateway.Client.Call(method, path, header, body, v)
}

// CreateBilling function to call Create Invoice/Billing
func (gateway *CoreGateway) CreateBilling(br BillingRequest) (resp BillingResponse, respError ResponseError, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	brByte, err := json.Marshal(br)
	if err != nil {
		return
	}

	data := Encrypt(string(brByte), gateway.Client.ClientID, gateway.Client.ClientSecret)

	req := Request{
		ClientID: gateway.Client.ClientID,
		Data:     data,
	}

	reqByte, err := json.Marshal(req)
	if err != nil {
		return
	}

	var res map[string]interface{}

	err = gateway.Call("POST", "", headers, strings.NewReader(string(reqByte)), &res)
	if err != nil {
		return
	}

	if res["status"].(string) == StatusSuccess {
		dataRespStr, errDecrypt := Decrypt(res["data"].(string), gateway.Client.ClientID, gateway.Client.ClientSecret)
		if errDecrypt != nil {
			err = errDecrypt
			return
		}

		var dataResp BillingData

		json.Unmarshal([]byte(dataRespStr), &dataResp)
		resp = BillingResponse{
			res["status"].(string),
			dataResp,
		}
	} else {
		respError.Status = res["status"].(string)
		respError.Message = res["message"].(string)
	}

	return
}
