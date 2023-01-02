package requests

type SetCustomer struct {
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
}
