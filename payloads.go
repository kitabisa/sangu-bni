package bni

// Request any request that will be sent to BNI
type Request struct {
	ClientID string `json:"client_id"`
	Data     string `json:"data"`
}

// ResponseError any error response that received from BNI
type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// BillingRequest billing data that used for request on create and update
type BillingRequest struct {
	TrxType         string `json:"type" valid:"required"`
	ClientID        string `json:"client_id" valid:"required"`
	TrxID           string `json:"trx_id" valid:"required"`
	TrxAmount       string `json:"trx_amount" valid:"required"`
	CustomerName    string `json:"customer_name" valid:"required"`
	CustomerEmail   string `json:"customer_email"`
	CustomerPhone   string `json:"customer_phone"`
	VirtualAccount  string `json:"virtual_account"`
	DatetimeExpired string `json:"datetime_expired"`
	Description     string `json:"description"`
}

// BillingCreateRequest data for create invoice/billing request
type BillingCreateRequest struct {
	BillingRequest
	BillingType string `json:"billing_type" valid:"required"`
}

// BillingCreateResponse create invoice/billing response to client
type BillingCreateResponse struct {
	Status string `json:"status"`
	BillingCreateData
}

// BillingCreateData data for create invoice/billing response
type BillingCreateData struct {
	TrxID          string `json:"trx_id"`
	VirtualAccount string `json:"virtual_account"`
}

// BillingInquiryRequest request payload for inquiry invoice/billing response
type BillingInquiryRequest struct {
	TrxType  string `json:"type"`
	ClientID string `json:"client_id"`
	TrxID    string `json:"trx_id"`
}

// BillingInquiryResponse response data for inquiry invoice/billing response
type BillingInquiryResponse struct {
	Status string `json:"status"`
	BillingInquiryData
}

// BillingInquiryData detail of billing data
type BillingInquiryData struct {
	ClientID                   string `json:"client_id"`
	TrxID                      string `json:"trx_id"`
	TrxAmount                  string `json:"trx_amount"`
	VirtualAccount             string `json:"virtual_account"`
	CustomerName               string `json:"customer_name"`
	CustomerEmail              string `json:"customer_email"`
	CustomerPhone              string `json:"customer_phone"`
	DatetimeCreatedIso8601     string `json:"datetime_created_iso8601"`
	DatetimeExpiredIso8601     string `json:"datetime_expired_iso8601"`
	DatetimeLastUpdatedIso8601 string `json:"datetime_last_updated_iso8601"`
	Description                string `json:"description"`
	VaStatus                   string `json:"va_status"`
	PaymentAmount              string `json:"payment_amount"`
	PaymentNtb                 string `json:"payment_ntb"`
	BillingType                string `json:"Billing_type"`
	DatetimePaymentIso8601     string `json:"datetime_payment_iso8601"`
}
