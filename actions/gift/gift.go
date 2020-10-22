package gift

import (
	"fmt"

	"github.com/pagely/chargebee-go"
	"github.com/pagely/chargebee-go/models/gift"
)

func Create(params *gift.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/gifts/%v", id), nil)
}
func List(params *gift.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/gifts"), params)
}
func Claim(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/claim", id), nil)
}
func Cancel(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/cancel", id), nil)
}
func UpdateGift(id string, params *gift.UpdateGiftRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/gifts/%v/update_gift", id), params)
}
