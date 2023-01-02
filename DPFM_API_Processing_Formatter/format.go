package dpfm_api_processing_formatter

import (
	"convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// Header
// データ連携基盤とSAP Inbound Deliveryとの項目マッピング
func (psdc *SDC) ConvertToMappingHeader(sdc *dpfm_api_input_reader.SDC) (*MappingHeader, error) {
	data := sdc.SAPInboundDeliveryHeader

	isExportDelivery, err := ParseBoolPtr(data.IsExportDelivery)
	if err != nil {
		return nil, err
	}

	headerGrossWeight, err := ParseFloat32Ptr(data.HeaderGrossWeight)
	if err != nil {
		return nil, err
	}
	headerNetWeight, err := ParseFloat32Ptr(data.HeaderNetWeight)
	if err != nil {
		return nil, err
	}

	res := MappingHeader{
		Seller:                        sdc.BusinessPartnerID,
		DeliverFromParty:              sdc.BusinessPartnerID,
		DeliverToPlant:                data.ReceivingPlant,
		DeliverFromPlant:              data.ShippingPoint,
		BillFromParty:                 sdc.BusinessPartnerID,
		Payee:                         sdc.BusinessPartnerID,
		IsExportImport:                isExportDelivery,
		DocumentDate:                  data.DocumentDate,
		PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
		InvoiceDocumentDate:           data.BillingDocumentDate,
		LastChangeDate:                data.LastChangeDate,
		GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
		HeaderGrossWeight:             headerGrossWeight,
		HeaderNetWeight:               headerNetWeight,
		HeaderWeightUnit:              data.HeaderWeightUnit,
		Incoterms:                     data.IncotermsClassification,
	}

	return &res, nil
}

// Item
// データ連携基盤とSAP Inbound Deliveryとの項目マッピング
func (psdc *SDC) ConvertToMappingItem(sdc *dpfm_api_input_reader.SDC) (*[]MappingItem, error) {
	var res []MappingItem
	data := sdc.SAPInboundDeliveryHeader
	dataItem := sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryItem

	for _, dataItem := range dataItem {
		systemDate := GetSystemDatePtr()
		systemTime := GetSystemTimePtr()

		originalDeliveryQuantity, err := ParseFloat32Ptr(dataItem.OriginalDeliveryQuantity)
		if err != nil {
			return nil, err
		}

		actualDeliveryQuantity, err := ParseFloat32Ptr(dataItem.ActualDeliveryQuantity)
		if err != nil {
			return nil, err
		}

		itemGrossWeight, err := ParseFloat32Ptr(dataItem.ItemGrossWeight)
		if err != nil {
			return nil, err
		}

		itemNetWeight, err := ParseFloat32Ptr(dataItem.ItemNetWeight)
		if err != nil {
			return nil, err
		}

		itemIsBillingRelevant, err := ParseBoolPtr(dataItem.ItemIsBillingRelevant)
		if err != nil {
			return nil, err
		}

		res = append(res, MappingItem{
			SAPDeliveryDocument:              data.DeliveryDocument,
			Seller:                           sdc.BusinessPartnerID,
			BillFromParty:                    sdc.BusinessPartnerID,
			Payee:                            sdc.BusinessPartnerID,
			DeliverFromPlantStorageLocation:  dataItem.StorageLocation,
			DeliverFromPlantBatch:            dataItem.Batch,
			StockConfirmationBusinessPartner: sdc.BusinessPartnerID,
			StockConfirmationPlant:           data.ShippingPoint,
			StockConfirmationPlantBatch:      dataItem.Batch,
			DeliveryDocumentItemTextBySeller: dataItem.DeliveryDocumentItemText,
			BaseUnit:                         dataItem.BaseUnit,
			OriginalQuantityInBaseUnit:       originalDeliveryQuantity,
			DeliveryUnit:                     dataItem.DeliveryQuantityUnit,
			ActualGoodsIssueDate:             data.ActualGoodsMovementDate,
			ActualGoodsIssueQuantity:         actualDeliveryQuantity,
			ActualGoodsReceiptQuantity:       actualDeliveryQuantity,
			CreationDate:                     systemDate,
			CreationTime:                     systemTime,
			ItemGrossWeight:                  itemGrossWeight,
			ItemNetWeight:                    itemNetWeight,
			ItemWeightUnit:                   dataItem.ItemWeightUnit,
			ItemIsBillingRelevant:            itemIsBillingRelevant,
			LastChangeDate:                   systemDate,
		})
	}
	return &res, nil
}

// Partner
func (psdc *SDC) ConvertToMappingPartner(sdc *dpfm_api_input_reader.SDC) (*[]MappingPartner, error) {
	var res []MappingPartner
	data := sdc.SAPInboundDeliveryHeader
	dataPartner := sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner

	for range dataPartner {
		res = append(res, MappingPartner{
			SAPDeliveryDocument: data.DeliveryDocument,
		})
	}

	return &res, nil
}

// Address
// データ連携基盤とSAP Inbound Deliveryとの項目マッピング
func (psdc *SDC) ConvertToMappingAddress(sdc *dpfm_api_input_reader.SDC) (*[]MappingAddress, error) {
	var res []MappingAddress
	data := sdc.SAPInboundDeliveryHeader
	dataPartner := sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner

	for _, dataPartner := range dataPartner {
		dataAddress := dataPartner.SAPInboundDeliveryPartnerAddress
		for _, dataAddress := range dataAddress {
			res = append(res, MappingAddress{
				SAPDeliveryDocument: data.DeliveryDocument,
				PostalCode:          dataAddress.PostalCode,
				LocalRegion:         dataAddress.Region,
				Country:             dataAddress.Country,
				StreetName:          dataAddress.StreetName,
				CityName:            dataAddress.CityName,
			})
		}
	}
	return &res, nil
}

// 番号処理
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	var res []CodeConversionQueryGets

	for i := 0; true; i++ {
		pm := &requests.CodeConversionQueryGets{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionHeader(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionHeader, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionHeader{}

	pm.SAPDeliveryDocument = dataMap["DeliveryDocument"].CodeConvertFrom
	pm.DeliveryDocument, err = ParseInt(dataMap["DeliveryDocument"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.Buyer, err = ParseIntPtr(GetStringPtr(dataMap["Buyer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.DeliverToParty, err = ParseIntPtr(GetStringPtr(dataMap["DeliverToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.BillToParty, err = ParseIntPtr(GetStringPtr(dataMap["BillToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.Payer, err = ParseIntPtr(GetStringPtr(dataMap["Payer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionHeader{
		SAPDeliveryDocument: data.SAPDeliveryDocument,
		DeliveryDocument:    data.DeliveryDocument,
		Buyer:               data.Buyer,
		DeliverToParty:      data.DeliverToParty,
		BillToParty:         data.BillToParty,
		Payer:               data.Payer,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionItem(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItem, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItem{}

	pm.SAPDeliveryDocumentItem = dataMap["DeliveryDocumentItem"].CodeConvertFrom
	pm.DeliveryDocumentItem, err = ParseInt(dataMap["DeliveryDocumentItem"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.Product = GetStringPtr(dataMap["Product"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItem{
		SAPDeliveryDocumentItem: data.SAPDeliveryDocumentItem,
		DeliveryDocumentItem:    data.DeliveryDocumentItem,
		Product:                 data.Product,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionPartner(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionPartner, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionPartner{}

	pm.PartnerFunction = dataMap["PartnerFunction"].CodeConvertTo
	pm.BusinessPartner, err = ParseInt(dataMap["BusinessPartner"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionPartner{
		PartnerFunction: data.PartnerFunction,
		BusinessPartner: data.BusinessPartner,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionAddress(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionAddress, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionAddress{}

	pm.AddressID, err = ParseInt(dataMap["AddressID"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionAddress{
		AddressID: data.AddressID,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConversionData() *[]ConversionData {
	var res []ConversionData

	for _, v := range *psdc.CodeConversionItem {
		pm := &requests.ConversionData{}

		pm.SAPDeliveryDocument = psdc.CodeConversionHeader.SAPDeliveryDocument
		pm.DeliveryDocument = psdc.CodeConversionHeader.DeliveryDocument
		pm.SAPDeliveryDocumentItem = v.SAPDeliveryDocumentItem
		pm.DeliveryDocumentItem = v.DeliveryDocumentItem

		data := pm
		res = append(res, ConversionData{
			SAPDeliveryDocument:     data.SAPDeliveryDocument,
			DeliveryDocument:        data.DeliveryDocument,
			SAPDeliveryDocumentItem: data.SAPDeliveryDocumentItem,
			DeliveryDocumentItem:    data.DeliveryDocumentItem,
		})
	}

	return &res
}

// 個別処理
func (psdc *SDC) ConvertToSetCustomer(headerPartner dpfm_api_input_reader.SAPInboundDeliveryHeaderPartner) *SetCustomer {
	pm := &requests.SetCustomer{}

	pm.PartnerFunction = headerPartner.PartnerFunction
	pm.Customer = headerPartner.Customer

	data := pm
	res := SetCustomer{
		PartnerFunction: data.PartnerFunction,
		Customer:        data.Customer,
	}

	return &res
}

func (psdc *SDC) ConvertToSetCustomerHeader(customer []SetCustomer) *SetCustomerHeader {
	pm := &requests.SetCustomerHeader{}

	customerMap := make(map[string]SetCustomer, len(customer))
	for _, v := range customer {
		customerMap[v.PartnerFunction] = v
	}

	pm.CustomerForBillToParty = customerMap["BL"].Customer
	pm.CustomerForPayer = customerMap["PY"].Customer

	data := pm
	res := SetCustomerHeader{
		CustomerForBillToParty: data.CustomerForBillToParty,
		CustomerForPayer:       data.CustomerForPayer,
	}

	return &res
}

func (psdc *SDC) ConvertToSetPartnerFunction(partnerFunction string) *SetPartnerFunction {
	pm := &requests.SetPartnerFunction{}

	pm.PartnerFunction = partnerFunction

	data := pm
	res := SetPartnerFunction{
		PartnerFunction: data.PartnerFunction,
	}

	return &res
}
