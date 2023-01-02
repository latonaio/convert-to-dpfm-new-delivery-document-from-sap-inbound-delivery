package dpfm_api_processing_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MappingHeader         *MappingHeader           `json:"MappingHeader"`
	CodeConversionHeader  *CodeConversionHeader    `json:"CodeConversionHeader"`
	SetCustomerHeader     *SetCustomerHeader       `json:"CustomerHeader"`
	MappingItem           *[]MappingItem           `json:"MappingItem"`
	CodeConversionItem    *[]CodeConversionItem    `json:"CodeConversionItem"`
	MappingPartner        *[]MappingPartner        `json:"MappingPartner"`
	SetPartnerFunction    *[]SetPartnerFunction    `json:"PartnerFunction"`
	CodeConversionPartner *[]CodeConversionPartner `json:"CodeConversionPartner"`
	MappingAddress        *[]MappingAddress        `json:"MappingAddress"`
	CodeConversionAddress *[]CodeConversionAddress `json:"CodeConversionAddress"`
	ConversionData        *[]ConversionData        `json:"ConversionData"`
}

type Header struct {
	DeliveryDocument                       int       `json:"DeliveryDocument"`
	SupplyChainRelationshipID              *int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int      `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  *int      `json:"Buyer"`
	Seller                                 *int      `json:"Seller"`
	DeliverToParty                         *int      `json:"DeliverToParty"`
	DeliverFromParty                       *int      `json:"DeliverFromParty"`
	DeliverToPlant                         *string   `json:"DeliverToPlant"`
	DeliverFromPlant                       *string   `json:"DeliverFromPlant"`
	BillToParty                            *int      `json:"BillToParty"`
	BillFromParty                          *int      `json:"BillFromParty"`
	BillToCountry                          *string   `json:"BillToCountry"`
	BillFromCountry                        *string   `json:"BillFromCountry"`
	Payer                                  *int      `json:"Payer"`
	Payee                                  *int      `json:"Payee"`
	IsExportImport                         *bool     `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string   `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string   `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int      `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int      `json:"ReferenceDocumentItem"`
	OrderID                                *int      `json:"OrderID"`
	OrderItem                              *int      `json:"OrderItem"`
	ContractType                           *string   `json:"ContractType"`
	OrderValidityStartDate                 *string   `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string   `json:"OrderValidityEndDate"`
	DocumentDate                           *string   `json:"DocumentDate"`
	PlannedGoodsIssueDate                  *string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string   `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string   `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool     `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string   `json:"HeaderDeliveryStatus"`
	CreationDate                           *string   `json:"CreationDate"`
	CreationTime                           *string   `json:"CreationTime"`
	LastChangeDate                         *string   `json:"LastChangeDate"`
	LastChangeTime                         *string   `json:"LastChangeTime"`
	GoodsIssueOrReceiptSlipNumber          *string   `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string   `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string   `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool     `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32  `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32  `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string   `json:"HeaderWeightUnit"`
	Incoterms                              *string   `json:"Incoterms"`
	TransactionCurrency                    *string   `json:"TransactionCurrency"`
	StockIsFullyConfirmed                  *float32  `json:"StockIsFullyConfirmed"`
	HeaderDeliveryBlockStatus              *float32  `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *float32  `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *float32  `json:"HeaderReceivingBlockStatus"`
	Partner                                []Partner `json:"Partner"`
	Item                                   []Item    `json:"Item"`
	Address                                []Address `json:"Address"`
}

type Partner struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Item struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  *string  `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                         *int     `json:"Buyer"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	BillToParty                                   *int     `json:"BillToParty"`
	BillFromParty                                 *int     `json:"BillFromParty"`
	BillToCountry                                 *string  `json:"BillToCountry"`
	BillFromCountry                               *string  `json:"BillFromCountry"`
	Payer                                         *int     `json:"Payer"`
	Payee                                         *int     `json:"Payee"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller"`
	Product                                       *string  `json:"Product"`
	ProductStandardID                             *string  `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      *string  `json:"BaseUnit"`
	OriginalQuantityInBaseUnit                    *float32 `json:"OriginalQuantityInBaseUnit"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity"`
	CreationDate                                  *string  `json:"CreationDate"`
	CreationTime                                  *string  `json:"CreationTime"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	ItemBillingConfStatus                         *string  `json:"ItemBillingConfStatus"`
	SalesCostGLAccount                            *string  `json:"SalesCostGLAccount"`
	ReceivingGLAccount                            *string  `json:"ReceivingGLAccount"`
	IssuingGoodsMovementType                      *string  `json:"IssuingGoodsMovementType"`
	ReceivingGoodsMovementType                    *string  `json:"ReceivingGoodsMovementType"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryIncompletionStatus                *string  `json:"ItemDeliveryIncompletionStatus"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant"`
	LastChangeDate                                *string  `json:"LastChangeDate"`
	OrderID                                       *int     `json:"OrderID"`
	OrderItem                                     *int     `json:"OrderItem"`
	OrderType                                     *string  `json:"OrderType"`
	ContractType                                  *string  `json:"ContractType"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate"`
	PaymentTerms                                  *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate"`
	Project                                       *string  `json:"Project"`
	ProfitCenter                                  *string  `json:"ProfitCenter"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	StockIsFullyConfirmed                         *bool    `json:"StockIsFullyConfirmed"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
}

type Address struct {
	DeliveryDocument int     `json:"DeliveryDocument"`
	AddressID        int     `json:"AddressID"`
	PostalCode       *string `json:"PostalCode"`
	LocalRegion      *string `json:"LocalRegion"`
	Country          *string `json:"Country"`
	District         *string `json:"District"`
	StreetName       *string `json:"StreetName"`
	CityName         *string `json:"CityName"`
	Building         *string `json:"Building"`
	Floor            *int    `json:"Floor"`
	Room             *int    `json:"Room"`
}

// 項目マッピング変換
type MappingHeader struct {
	Seller                        *int     `json:"Seller"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	DeliverToPlant                *string  `json:"DeliverToPlant"`
	DeliverFromPlant              *string  `json:"DeliverFromPlant"`
	BillFromParty                 *int     `json:"BillFromParty"`
	Payee                         *int     `json:"Payee"`
	IsExportImport                *bool    `json:"IsExportImport"`
	DocumentDate                  *string  `json:"DocumentDate"`
	PlannedGoodsIssueDate         *string  `json:"PlannedGoodsIssueDate"`
	InvoiceDocumentDate           *string  `json:"InvoiceDocumentDate"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	GoodsIssueOrReceiptSlipNumber *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderGrossWeight             *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight               *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit              *string  `json:"HeaderWeightUnit"`
	Incoterms                     *string  `json:"Incoterms"`
}

type MappingItem struct {
	SAPDeliveryDocument              string   `json:"SAPDeliveryDocument"`
	Seller                           *int     `json:"Seller"`
	BillFromParty                    *int     `json:"BillFromParty"`
	Payee                            *int     `json:"Payee"`
	DeliverFromPlantStorageLocation  *string  `json:"DeliverFromPlantStorageLocation"`
	DeliverFromPlantBatch            *string  `json:"DeliverFromPlantBatch"`
	StockConfirmationBusinessPartner *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch      *string  `json:"StockConfirmationPlantBatch"`
	DeliveryDocumentItemTextBySeller *string  `json:"DeliveryDocumentItemTextBySeller"`
	BaseUnit                         *string  `json:"BaseUnit"`
	OriginalQuantityInBaseUnit       *float32 `json:"OriginalQuantityInBaseUnit"`
	DeliveryUnit                     *string  `json:"DeliveryUnit"`
	ActualGoodsIssueDate             *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueQuantity         *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsReceiptQuantity       *float32 `json:"ActualGoodsReceiptQuantity"`
	CreationDate                     *string  `json:"CreationDate"`
	CreationTime                     *string  `json:"CreationTime"`
	ItemGrossWeight                  *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                    *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                   *string  `json:"ItemWeightUnit"`
	ItemIsBillingRelevant            *bool    `json:"ItemIsBillingRelevant"`
	LastChangeDate                   *string  `json:"LastChangeDate"`
}

type MappingPartner struct {
	SAPDeliveryDocument string `json:"SAPDeliveryDocument"`
}

type MappingAddress struct {
	SAPDeliveryDocument string  `json:"SAPDeliveryDocument"`
	PostalCode          *string `json:"PostalCode"`
	LocalRegion         *string `json:"LocalRegion"`
	Country             *string `json:"Country"`
	StreetName          *string `json:"StreetName"`
	CityName            *string `json:"CityName"`
}

// 番号変換
type CodeConversionKey struct {
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionQueryGets struct {
	CodeConversionID  int    `json:"CodeConversionID"`
	SystemConvertTo   string `json:"SystemConvertTo"`
	SystemConvertFrom string `json:"SystemConvertFrom"`
	LabelConvertTo    string `json:"LabelConvertTo"`
	LabelConvertFrom  string `json:"LabelConvertFrom"`
	CodeConvertFrom   string `json:"CodeConvertFrom"`
	CodeConvertTo     string `json:"CodeConvertTo"`
	BusinessPartner   int    `json:"BusinessPartner"`
}

type CodeConversionHeader struct {
	SAPDeliveryDocument string `json:"SAPDeliveryDocument"`
	DeliveryDocument    int    `json:"DeliveryDocument"`
	Buyer               *int   `json:"Buyer"`
	DeliverToParty      *int   `json:"DeliverToParty"`
	BillToParty         *int   `json:"BillToParty"`
	Payer               *int   `json:"Payer"`
}

type CodeConversionItem struct {
	SAPDeliveryDocumentItem string  `json:"SAPDeliveryDocumentItem"`
	DeliveryDocumentItem    int     `json:"DeliveryDocumentItem"`
	Product                 *string `json:"Product"`
}

type CodeConversionPartner struct {
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type CodeConversionAddress struct {
	AddressID int `json:"AddressID"`
}

type ConversionData struct {
	SAPDeliveryDocument     string `json:"SAPDeliveryDocument"`
	DeliveryDocument        int    `json:"DeliveryDocument"`
	SAPDeliveryDocumentItem string `json:"SAPDeliveryDocumentItem"`
	DeliveryDocumentItem    int    `json:"DeliveryDocumentItem"`
}

// 個別処理
type SetCustomer struct {
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
}

type SetCustomerHeader struct {
	CustomerForBillToParty *string `json:"CustomerForBillToParty"`
	CustomerForPayer       *string `json:"CustomerForPayer"`
}

type SetPartnerFunction struct {
	PartnerFunction string `json:"PartnerFunction"`
}
