package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Headerの処理
func (c *ConvertComplementer) ComplementMappingHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.MappingHeader, error) {
	res, err := psdc.ConvertToMappingHeader(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*dpfm_api_processing_formatter.CodeConversionHeader, error) {
	var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
	var args []interface{}

	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "DeliveryDocument", "DeliveryDocument", sdc.SAPInboundDeliveryHeader.DeliveryDocument))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "SoldToParty", "Buyer", sdc.SAPInboundDeliveryHeader.SoldToParty))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "ShipToParty", "DeliverToParty", sdc.SAPInboundDeliveryHeader.ShipToParty))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Customer", "BillToParty", psdc.SetCustomerHeader.CustomerForBillToParty))
	dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Customer", "Payer", psdc.SetCustomerHeader.CustomerForPayer))

	repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?)"
	for _, v := range dataKey {
		args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
	}

	rows, err := c.db.Query(
		`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
		CodeConvertFrom, CodeConvertTo, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
		WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCodeConversionHeader(dataQueryGets)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *ConvertComplementer) SetCustomerHeader(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) *dpfm_api_processing_formatter.SetCustomerHeader {
	var dataCustomer []dpfm_api_processing_formatter.SetCustomer

	for _, v := range sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner {
		partnerFunction := v.PartnerFunction
		if partnerFunction == "BL" || partnerFunction == "PY" {
			datum := psdc.ConvertToSetCustomer(v)
			dataCustomer = append(dataCustomer, *datum)
		}
	}

	data := psdc.ConvertToSetCustomerHeader(dataCustomer)

	return data
}
