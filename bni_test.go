package bni

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BniTestSuite struct {
	suite.Suite
	client Client
}

func TestBniTestSuite(t *testing.T) {
	suite.Run(t, new(BniTestSuite))
}

func (b *BniTestSuite) SetupSuite() {
	b.client = NewClient()
	b.client.BaseURL = "https://bni-staging-hostname"
	b.client.ClientID = "bni-staging-client-id"
	b.client.ClientSecret = "bni-staging-client-secret"
}

func (b *BniTestSuite) TestCreateBillingSuccess() {
	core := CoreGateway{
		Client: b.client,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	br := BillingRequest{
		Type:         "createbilling",
		ClientID:     "bni-staging-client-id",
		TrxID:        fmt.Sprint(r1.Intn(1000000)),
		TrxAmount:    "10000",
		BillingType:  BillTypeFixed,
		CustomerName: "fulan",
	}

	resp, respError, err := core.CreateBilling(br)
	assert.NotEqual(b.T(), BillingResponse{}, resp)
	assert.Equal(b.T(), StatusSuccess, resp.Status)
	assert.Equal(b.T(), ResponseError{}, respError)
	assert.Equal(b.T(), nil, err)
}

func (b *BniTestSuite) TestCreateBillingFail() {
	core := CoreGateway{
		Client: b.client,
	}

	br := BillingRequest{
		Type:         "createbilling",
		ClientID:     "bni-staging-client-id",
		TrxID:        "1",
		TrxAmount:    "10000",
		BillingType:  BillTypeFixed,
		CustomerName: "fulan",
	}

	// run it twice to make it sure duplicate
	_, _, _ = core.CreateBilling(br)
	resp, respError, err := core.CreateBilling(br)

	assert.Equal(b.T(), BillingResponse{}, resp)
	assert.NotEqual(b.T(), StatusSuccess, respError.Status)
	assert.NotEqual(b.T(), ResponseError{}, respError)
	assert.Equal(b.T(), nil, err)
}
