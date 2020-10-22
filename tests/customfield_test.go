package tests

import (
	"testing"

	"github.com/pagely/chargebee-go"
	"github.com/stretchr/testify/assert"
)

func TestCustomFieldExtractionFromResponse(t *testing.T) {
	data := subscriptionWithCustomfield()
	result := &chargebee.Result{}
	chargebee.UnmarshalJSON(data, result)
	s := result.Subscription
	c := result.Customer
	assert.Equal(t, "Female", s.CustomField["cf_gender"])
	assert.Equal(t, "12-10-1994", c.CustomField["cf_DOB"])
}
