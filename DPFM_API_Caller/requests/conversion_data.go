package requests

type ConversionData struct {
	SAPDeliveryDocument     string `json:"SAPDeliveryDocument"`
	DeliveryDocument        int    `json:"DeliveryDocument"`
	SAPDeliveryDocumentItem string `json:"SAPDeliveryDocumentItem"`
	DeliveryDocumentItem    int    `json:"DeliveryDocumentItem"`
}
