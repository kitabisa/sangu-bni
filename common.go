package bni

// types of invoice/billing
const (
	// BillTypeOpen Open payment: invoice/billing can be paid multiple times as long as it is still active.
	BillTypeOpen = "o"

	// BillTypeFixed Fixed payment: invoice/billing should be paid with exactly the same amount as requested.
	BillTypeFixed = "c"

	// BillTypePartial Installment/partial payment: invoice/billing can be paid multiple times as long as paid amount is less than the requested amount and still active.
	BillTypePartial = "i"

	// BillTypeMin Minimum payment: invoice/billing can be paid with greater than or equal to the requested amount.
	BillTypeMin = "m"

	// BillTypeOpenMin Open minimum payment: invoice/billing can be paid greater than or equal to the requested amount multiple times as long as it is still active.
	BillTypeOpenMin = "n"

	// BillTypeOpenMax Open maximum payment: invoice/billing can be paid less than or equal to the requested amount multiple times as long as it is still active.
	BillTypeOpenMax = "x"
)

// status code
const (
	// StatusSuccess Success
	StatusSuccess = "000" 
	
	// StatusInvalidParameter Incomplete/invalid Parameter(s).
	StatusInvalidParameter = "001" 
	
	// StatusNotAllowed IP address not allowed or wrong Client ID.
	StatusNotAllowed = "002" 
	
	// StatusSvcNotFound Service not found.
	StatusSvcNotFound = "004" 
	
	// StatusSvcNotDefined Service not defined.
	StatusSvcNotDefined = "005" 
	
	// StatusInvalidVANo Invalid VA Number.
	StatusInvalidVANo = "006" 
	
	// StatusTechFailure Technical Failure.
	StatusTechFailure = "008" 
	
	// StatusUnexpectedErr Unexpected Error.
	StatusUnexpectedErr = "009" 
	
	// StatusRTO Request Timeout.
	StatusRTO = "010" 
	
	// StatusBillNotMatch Billing type does not match billing amount.
	StatusBillNotMatch = "011" 
	
	// StatusInvalidExpireTime Invalid expiry date/time.
	StatusInvalidExpireTime = "012" 
	
	// StatusInvalidAmount IDR currency cannot have billing amount with decimal fraction.
	StatusInvalidAmount = "013" 
	
	// StatusBillingNotFound Billing not found.
	StatusBillingNotFound = "101" 
	
	// StatusVAInUse VA Number is in use.
	StatusVAInUse = "102" 
	
	// StatusExpired Billing has been expired.
	StatusExpired = "103" 
	
	// StatusDuplicate Duplicate Billing ID.
	StatusDuplicate = "105" 
	
	// StatusInvalidAmountIsFixed Amount can not be changed.
	StatusInvalidAmountIsFixed = "107" 
	
	// StatusFailSendSMS Failed to send SMS Payment.
	StatusFailSendSMS = "200" 
	
	// StatusSMSOnlyForFixed SMS Payment can only be used with Fixed Payment.
	StatusSMSOnlyForFixed = "201" 
	
	// StatusSystemOffline System is temporarily offline.
	StatusSystemOffline = "997" 
	
	// StatusContentTypeNotSet "Content-Type" header not defined as it should be.
	StatusContentTypeNotSet = "998" 
	
	// StatusInternalErr Internal Error.
	StatusInternalErr = "999" 
	
)