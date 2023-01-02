package requests

type SAPInboundDeliveryHeaderPartner struct {
	AddressID       *string `json:"AddressID"`
	ContactPerson   *string `json:"ContactPerson"`
	Customer        *string `json:"Customer"`
	PartnerFunction string  `json:"PartnerFunction"`
	Personnel       *string `json:"Personnel"`
	SDDocument      *string `json:"SDDocument"`
	SDDocumentItem  *string `json:"SDDocumentItem"`
	Supplier        *string `json:"Supplier"`
}
