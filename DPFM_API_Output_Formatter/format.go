package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Processing_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*Header, error) {
	mappingHeader := psdc.MappingHeader
	codeConversionHeader := psdc.CodeConversionHeader

	header := Header{}

	data, err := json.Marshal(mappingHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	header.DeliveryDocument = codeConversionHeader.DeliveryDocument
	header.Buyer = codeConversionHeader.Buyer
	header.DeliverToParty = codeConversionHeader.DeliverToParty
	header.BillToParty = codeConversionHeader.BillToParty
	header.Payer = codeConversionHeader.Payer

	return &header, nil
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Item, error) {
	var items []Item
	mappingItem := psdc.MappingItem
	codeConversionItem := psdc.CodeConversionItem
	conversionData := psdc.ConversionData

	for i := range *mappingItem {
		item := Item{}

		data, err := json.Marshal((*mappingItem)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &item)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.SAPDeliveryDocument == (*mappingItem)[i].SAPDeliveryDocument {
				item.DeliveryDocument = v.DeliveryDocument
				break
			}
		}
		item.DeliveryDocumentItem = (*codeConversionItem)[i].DeliveryDocumentItem
		item.Product = (*codeConversionItem)[i].Product

		items = append(items, item)
	}

	return &items, nil
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Partner, error) {
	var partners []Partner
	mappingPartner := psdc.MappingPartner
	codeConversionPartner := psdc.CodeConversionPartner
	conversionData := psdc.ConversionData

	for i := range *mappingPartner {
		partner := Partner{}

		data, err := json.Marshal((*mappingPartner)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &partner)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.SAPDeliveryDocument == (*mappingPartner)[i].SAPDeliveryDocument {
				partner.DeliveryDocument = v.DeliveryDocument
				break
			}
		}
		partner.PartnerFunction = (*codeConversionPartner)[i].PartnerFunction
		partner.BusinessPartner = (*codeConversionPartner)[i].BusinessPartner

		partners = append(partners, partner)
	}

	return &partners, nil
}

func ConvertToAddress(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.SDC,
) (*[]Address, error) {
	var addresses []Address
	mappingAddress := psdc.MappingAddress
	codeConversionAddress := psdc.CodeConversionAddress
	conversionData := psdc.ConversionData

	for i := range *mappingAddress {
		address := Address{}

		data, err := json.Marshal((*mappingAddress)[i])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &address)
		if err != nil {
			return nil, err
		}

		for _, v := range *conversionData {
			if v.SAPDeliveryDocument == (*mappingAddress)[i].SAPDeliveryDocument {
				address.DeliveryDocument = v.DeliveryDocument
				break
			}
		}
		address.AddressID = (*codeConversionAddress)[i].AddressID

		addresses = append(addresses, address)
	}

	return &addresses, nil
}
