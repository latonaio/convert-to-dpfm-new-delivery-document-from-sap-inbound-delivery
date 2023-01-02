package requests

type SetCustomerHeader struct {
	CustomerForBillToParty *string `json:"CustomerForBillToParty"`
	CustomerForPayer       *string `json:"CustomerForPayer"`
}
