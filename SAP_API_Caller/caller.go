package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-supplier-quotation-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetSupplierQuotation(supplierQuotation, supplierQuotationItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(supplierQuotation)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(supplierQuotation, supplierQuotationItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(supplierQuotation string) {
	headerData, err := c.callWarehouseSrvAPIRequirementHeader("SupplierQuotation", supplierQuotation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)
}

func (c *SAPAPICaller) callWarehouseSrvAPIRequirementHeader(api, supplierQuotation string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_supplierquotation_2/srvd_a2x/sap/supplierquotation/0001", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, supplierQuotation)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(supplierQuotation, supplierQuotationItem string) {
	data, err := c.callWarehouseSrvAPIRequirementItem("SupplierQuotation", supplierQuotation, supplierQuotationItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callWarehouseSrvAPIRequirementItem(api, supplierQuotation, supplierQuotationItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "api_supplierquotation_2/srvd_a2x/sap/supplierquotation/0001", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, supplierQuotation, supplierQuotationItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, supplierQuotation string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SupplierQuotation eq '%s'", supplierQuotation)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, supplierQuotation, supplierQuotationItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SupplierQuotation eq '%s' and SupplierQuotationItem eq '%s'", supplierQuotation, supplierQuotationItem)
	return params
}
