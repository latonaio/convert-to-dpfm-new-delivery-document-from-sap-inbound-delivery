package requests

type CodeConversionHeader struct {
	SAPDeliveryDocument string `json:"SAPDeliveryDocument"`
	DeliveryDocument    int    `json:"DeliveryDocument"`
	Buyer               *int   `json:"Buyer"`
	DeliverToParty      *int   `json:"DeliverToParty"`
	BillToParty         *int   `json:"BillToParty"`
	Payer               *int   `json:"Payer"`
}
