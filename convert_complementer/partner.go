package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingPartner(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingPartner, error) {
	res, err := psdc.ConvertToMappingPartner(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionPartner(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionPartner, error) {
	var data []dpfm_api_processing_formatter.CodeConversionPartner

	for i, partner := range sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner {
		var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
		var args []interface{}

		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "PartnerFunction", "PartnerFunction", (*psdc.SetPartnerFunction)[i].PartnerFunction))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Customer", "BusinessPartner", partner.Customer))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Supplier", "BusinessPartner", partner.Supplier))

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

		datum, err := psdc.ConvertToCodeConversionPartner(dataQueryGets)
		if err != nil {
			return nil, err
		}

		data = append(data, *datum)
	}

	return &data, nil
}

func (c *ConvertComplementer) SetPartnerFunction(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) *[]dpfm_api_processing_formatter.SetPartnerFunction {
	var data []dpfm_api_processing_formatter.SetPartnerFunction

	for _, v := range sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner {
		var datum *dpfm_api_processing_formatter.SetPartnerFunction
		partnerFunction := v.PartnerFunction
		if partnerFunction == "SP" || partnerFunction == "SH" || partnerFunction == "BP" || partnerFunction == "PY" {
			datum = psdc.ConvertToSetPartnerFunction(partnerFunction)
		} else {
			datum = psdc.ConvertToSetPartnerFunction("")
		}
		data = append(data, *datum)
	}

	return &data
}
