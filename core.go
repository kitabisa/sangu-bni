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
func (gateway *CoreGateway) CreateBilling(br BillingCreateRequest) (resp BillingCreateResponse, respError ResponseError, err error) {
	// set type to create
	br.TrxType = TypeCreate

	respError, err = gateway.processing(br, &resp.BillingCreateData)
	if err == nil && respError.Status == "" {
		resp.Status = StatusSuccess
	}

	return
}

// InquiryBilling function to call Inquiry Invoice/Billing
func (gateway *CoreGateway) InquiryBilling(br BillingInquiryRequest) (resp BillingInquiryResponse, respError ResponseError, err error) {
	// set type to inquiry
	br.TrxType = TypeInquiry

	respError, err = gateway.processing(br, &resp.BillingInquiryData)
	if err == nil && respError.Status == "" {
		resp.Status = StatusSuccess
	}

	return
}

// UpdateBilling function to call Update Transaction
func (gateway *CoreGateway) UpdateBilling(br BillingRequest) (resp BillingCreateResponse, respError ResponseError, err error) {
	// set type to update
	br.TrxType = TypeUpdate

	respError, err = gateway.processing(br, &resp.BillingCreateData)
	if err == nil && respError.Status == "" {
		resp.Status = StatusSuccess
	}

	return
}

// general processing request/response to bni api
func (gateway *CoreGateway) processing(request interface{}, respData interface{}) (respError ResponseError, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	brByte, err := json.Marshal(request)
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

		json.Unmarshal([]byte(dataRespStr), &respData)
	} else {
		respError.Status = res["status"].(string)
		respError.Message = res["message"].(string)
	}

	return
}
