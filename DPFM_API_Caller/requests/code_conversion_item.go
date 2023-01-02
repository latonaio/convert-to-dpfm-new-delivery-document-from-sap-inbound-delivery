package requests

type CodeConversionItem struct {
	SAPDeliveryDocumentItem string  `json:"SAPDeliveryDocumentItem"`
	DeliveryDocumentItem    int     `json:"DeliveryDocumentItem"`
	Product                 *string `json:"Product"`
}
