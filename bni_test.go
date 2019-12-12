package bni

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BniTestSuite struct {
	suite.Suite
	client Client
}

type credentials struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
}

func TestBniTestSuite(t *testing.T) {
	suite.Run(t, new(BniTestSuite))
}

func (b *BniTestSuite) SetupSuite() {
	theToml, err := ioutil.ReadFile("credentials.toml")
	if err != nil {
		b.T().Log(err)
		b.T().FailNow()
	}

	var cred credentials
	if _, err := toml.Decode(string(theToml), &cred); err != nil {
		b.T().Log(err)
		b.T().FailNow()
	}

	b.client = NewClient()
	b.client.BaseURL = cred.BaseUrl
	b.client.ClientID = cred.ClientId
	b.client.ClientSecret = cred.ClientSecret
}

func (b *BniTestSuite) TestCreateBillingSuccess() {
	core := CoreGateway{
		Client: b.client,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	br := BillingCreateRequest{
		Type:         TypeCreate,
		ClientID:     b.client.ClientID,
		TrxID:        fmt.Sprint(r1.Intn(1000000)),
		TrxAmount:    "10000",
		BillingType:  BillTypeFixed,
		CustomerName: "fulan",
	}

	resp, respError, err := core.CreateBilling(br)
	assert.NotEqual(b.T(), BillingCreateResponse{}, resp)
	assert.Equal(b.T(), StatusSuccess, resp.Status)
	assert.Equal(b.T(), ResponseError{}, respError)
	assert.Equal(b.T(), nil, err)
}

func (b *BniTestSuite) TestCreateBillingFail() {
	core := CoreGateway{
		Client: b.client,
	}

	br := BillingCreateRequest{
		Type:         TypeCreate,
		ClientID:     b.client.ClientID,
		TrxID:        "1",
		TrxAmount:    "10000",
		BillingType:  BillTypeFixed,
		CustomerName: "fulan",
	}

	// run it twice to make it sure duplicate
	_, _, _ = core.CreateBilling(br)
	resp, respError, err := core.CreateBilling(br)

	assert.Equal(b.T(), BillingCreateResponse{}, resp)
	assert.NotEqual(b.T(), StatusSuccess, respError.Status)
	assert.NotEqual(b.T(), ResponseError{}, respError)
	assert.Equal(b.T(), nil, err)
}

func (b *BniTestSuite) TestInquiryBillingSuccess() {
	core := CoreGateway{
		Client: b.client,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	br := BillingCreateRequest{
		Type:         TypeCreate,
		ClientID:     b.client.ClientID,
		TrxID:        fmt.Sprint(r1.Intn(1000000)),
		TrxAmount:    "10000",
		BillingType:  BillTypeFixed,
		CustomerName: "fulan",
	}

	_, respCreateError, err := core.CreateBilling(br)
	if err != nil || !cmp.Equal(respCreateError, ResponseError{}) {
		b.T().Log(err)
		b.T().Log(respCreateError)
		b.T().FailNow()
	}

	bdr := BillingDetailRequest{
		Type:     TypeInquiry,
		ClientID: b.client.ClientID,
		TrxID:    br.TrxID,
	}

	resp, respError, err := core.InquiryBilling(bdr)
	assert.NotEqual(b.T(), BillingDetailResponse{}, resp)
	assert.Equal(b.T(), resp.TrxID, bdr.TrxID)
	assert.Equal(b.T(), StatusSuccess, resp.Status)
	assert.Equal(b.T(), ResponseError{}, respError)
	assert.Equal(b.T(), nil, err)
}

func (b *BniTestSuite) TestInquiryBillingFail() {
	core := CoreGateway{
		Client: b.client,
	}

	bdr := BillingDetailRequest{
		Type:     TypeInquiry,
		ClientID: b.client.ClientID,
		TrxID:    "a",
	}

	resp, respError, err := core.InquiryBilling(bdr)
	assert.Equal(b.T(), BillingDetailResponse{}, resp)
	assert.Equal(b.T(), respError.Status, StatusBillingNotFound)
	assert.Equal(b.T(), nil, err)
}
