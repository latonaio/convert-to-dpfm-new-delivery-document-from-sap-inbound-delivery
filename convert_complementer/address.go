package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingAddress(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingAddress, error) {
	res, err := psdc.ConvertToMappingAddress(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionAddress(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionAddress, error) {
	var data []dpfm_api_processing_formatter.CodeConversionAddress

	for _, partner := range sdc.SAPInboundDeliveryHeader.SAPInboundDeliveryHeaderPartner {
		for _, address := range partner.SAPInboundDeliveryPartnerAddress {
			var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
			var args []interface{}

			dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "AddressID", "AddressID", address.AddressID))

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

			datum, err := psdc.ConvertToCodeConversionAddress(dataQueryGets)
			if err != nil {
				return nil, err
			}

			data = append(data, *datum)
		}
	}

	return &data, nil
}
