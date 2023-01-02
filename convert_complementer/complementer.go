package convert_complementer

import (
	"context"
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Output_Formatter"
	dpfm_api_processing_data_formatter "convert-to-dpfm-delivery-document-from-sap-inbound-delivery/DPFM_API_Processing_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ConvertComplementer struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewConvertComplementer(ctx context.Context, db *database.Mysql, l *logger.Logger) *ConvertComplementer {
	return &ConvertComplementer{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (c *ConvertComplementer) CreateSdc(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(7)

	// Header
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-0. データ連携基盤Delivery Document HeaderとSAP Sales Orderとの項目マッピング変換
		psdc.MappingHeader, e = c.ComplementMappingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Item
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-1. データ連携基盤Delivery Document ItemとSAP Sales Orderとの項目マッピング変換
		psdc.MappingItem, e = c.ComplementMappingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <1-2. 変換元のCustomer(BillToParty)、Customer(Payer)のセット>
		psdc.SetCustomerHeader = c.SetCustomerHeader(sdc, psdc)

		// <1-1. 番号変換＞
		psdc.CodeConversionHeader, e = c.CodeConversionHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// <5-1. 番号変換＞
		psdc.CodeConversionItem, e = c.CodeConversionItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		psdc.ConversionData = c.ConversionData(sdc, psdc)
	}(&wg)

	// Partner
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-1. データ連携基盤Delivery Document PartnerとSAP Sales Order との項目マッピング変換
		psdc.MappingPartner, e = c.ComplementMappingPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <2-2. 変換元のPartnerFunctionのセット>
		psdc.SetPartnerFunction = c.SetPartnerFunction(sdc, psdc)

		// <2-1. 番号変換＞
		psdc.CodeConversionPartner, e = c.CodeConversionPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	// Address
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 5-1. データ連携基盤Delivery Document AddressとSAP Sales Orderとの項目マッピング変換
		psdc.MappingAddress, e = c.ComplementMappingAddress(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// <5-1. 番号変換＞
		psdc.CodeConversionAddress, e = c.CodeConversionAddress(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	c.l.Info(psdc)
	osdc, err = c.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
