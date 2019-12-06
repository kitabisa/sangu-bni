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

// BillingRequest data for create invoice/billing request
type BillingRequest struct {
	Type            string `json:"type" valid:"required"`
	ClientID        string `json:"client_id" valid:"required"`
	TrxID           string `json:"trx_id" valid:"required"`
	TrxAmount       string `json:"trx_amount" valid:"required"`
	BillingType     string `json:"billing_type" valid:"required"`
	CustomerName    string `json:"customer_name" valid:"required"`
	CustomerEmail   string `json:"customer_email"`
	CustomerPhone   string `json:"customer_phone"`
	VirtualAccount  string `json:"virtual_account"`
	DatetimeExpired string `json:"datetime_expired"`
	Description     string `json:"description"`
}

// BillingResponse create invoice/billing response to client
type BillingResponse struct {
	Status string `json:"status"`
	BillingData
}

// BillingData data for create invoice/billing response
type BillingData struct {
	TrxID          string `json:"trx_id"`
	VirtualAccount string `json:"virtual_account"`
}
