package sap_api_input_reader

type EC_MC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	BillingDocument struct {
		BillingDocument string `json:"document_no"`
		DeliverTo       string `json:"deliver_to"`
		Quantity        string `json:"quantity"`
		PickedQuantity  string `json:"picked_quantity"`
		TotalNetAmount  string `json:"price"`
		Batch           string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
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
	} `json:"production_order"`
	APISchema           string `json:"api_schema"`
	MaterialCode        string `json:"material_code"`
	Plant_Supplier      string `json:"plant/supplier"`
	Stock               string `json:"stock"`
	BillingDocumentType string `json:"document_type"`
	SDDocument          string `json:"document_no"`
	PlannedDate         string `json:"planned_date"`
	BillingDocumentDate string `json:"validated_date"`
	Deleted             bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	SupplierQuotation struct {
		SupplierQuotation              string `json:"SupplierQuotation"`
		CompanyCode                    string `json:"CompanyCode"`
		PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
		PurchasingDocumentType         string `json:"PurchasingDocumentType"`
		Supplier                       string `json:"Supplier"`
		CreatedByUser                  string `json:"CreatedByUser"`
		CreationDate                   string `json:"CreationDate"`
		Language                       string `json:"Language"`
		DocumentCurrency               string `json:"DocumentCurrency"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
		IncotermsVersion               string `json:"IncotermsVersion"`
		IncotermsLocation1             string `json:"IncotermsLocation1"`
		IncotermsLocation2             string `json:"IncotermsLocation2"`
		PaymentTerms                   string `json:"PaymentTerms"`
		CashDiscount1Days              int    `json:"CashDiscount1Days"`
		CashDiscount2Days              int    `json:"CashDiscount2Days"`
		CashDiscount1Percent           int    `json:"CashDiscount1Percent"`
		CashDiscount2Percent           int    `json:"CashDiscount2Percent"`
		NetPaymentDays                 int    `json:"NetPaymentDays"`
		PricingProcedure               string `json:"PricingProcedure"`
		PurchasingOrganization         string `json:"PurchasingOrganization"`
		PurchasingGroup                string `json:"PurchasingGroup"`
		PurchasingDocumentOrderDate    string `json:"PurchasingDocumentOrderDate"`
		AbsoluteExchangeRate           int    `json:"AbsoluteExchangeRate"`
		ExchRateIsIndirectQuotation    bool   `json:"ExchRateIsIndirectQuotation"`
		EffectiveExchangeRate          int    `json:"EffectiveExchangeRate"`
		ExchangeRateIsFixed            bool   `json:"ExchangeRateIsFixed"`
		PurContrValidityStartDate      string `json:"PurContrValidityStartDate"`
		PurContrValidityEndDate        string `json:"PurContrValidityEndDate"`
		IsEndOfPurposeBlocked          string `json:"IsEndOfPurposeBlocked"`
		PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
		RequestForQuotation            string `json:"RequestForQuotation"`
		SupplierQuotationExternalID    string `json:"SupplierQuotationExternalID"`
		QuotationSubmissionDate        string `json:"QuotationSubmissionDate"`
		QuotationLatestSubmissionDate  string `json:"QuotationLatestSubmissionDate"`
		BindingPeriodValidityEndDate   string `json:"BindingPeriodValidityEndDate"`
		QtnLifecycleStatus             string `json:"QtnLifecycleStatus"`
		FollowOnDocumentCategory       string `json:"FollowOnDocumentCategory"`
		PurgDocFollowOnDocumentType    string `json:"PurgDocFollowOnDocumentType"`
		SupplierQuotationItem          struct {
			SupplierQuotation              string `json:"SupplierQuotation"`
			SupplierQuotationItem          string `json:"SupplierQuotationItem"`
			PurchasingDocumentCategory     string `json:"PurchasingDocumentCategory"`
			PurchasingDocumentItemText     string `json:"PurchasingDocumentItemText"`
			Material                       string `json:"Material"`
			ProductTypeCode                string `json:"ProductTypeCode"`
			ManufacturerMaterial           string `json:"ManufacturerMaterial"`
			SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
			ManufacturerPartNmbr           string `json:"ManufacturerPartNmbr"`
			Manufacturer                   string `json:"Manufacturer"`
			MaterialGroup                  string `json:"MaterialGroup"`
			Plant                          string `json:"Plant"`
			ManualDeliveryAddressID        string `json:"ManualDeliveryAddressID"`
			ReferenceDeliveryAddressID     string `json:"ReferenceDeliveryAddressID"`
			AddressID                      string `json:"AddressID"`
			ItemDeliveryAddressID          string `json:"ItemDeliveryAddressID"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
			IncotermsLocation1             string `json:"IncotermsLocation1"`
			IncotermsLocation2             string `json:"IncotermsLocation2"`
			ScheduleLineDeliveryDate       string `json:"ScheduleLineDeliveryDate"`
			ScheduleLineOrderQuantity      int    `json:"ScheduleLineOrderQuantity"`
			AwardedQuantity                int    `json:"AwardedQuantity"`
			PerformancePeriodStartDate     string `json:"PerformancePeriodStartDate"`
			PerformancePeriodEndDate       string `json:"PerformancePeriodEndDate"`
			OrderPriceUnit                 string `json:"OrderPriceUnit"`
			OrderPriceUnitToOrderUnitNmrtr int    `json:"OrderPriceUnitToOrderUnitNmrtr"`
			OrdPriceUnitToOrderUnitDnmntr  int    `json:"OrdPriceUnitToOrderUnitDnmntr"`
			OrderQuantityUnit              string `json:"OrderQuantityUnit"`
			OrderItemQtyToBaseQtyNmrtr     int    `json:"OrderItemQtyToBaseQtyNmrtr"`
			OrderItemQtyToBaseQtyDnmntr    int    `json:"OrderItemQtyToBaseQtyDnmntr"`
			PurgDocPriceDate               string `json:"PurgDocPriceDate"`
			BaseUnit                       string `json:"BaseUnit"`
			NetAmount                      int    `json:"NetAmount"`
			GrossAmount                    int    `json:"GrossAmount"`
			EffectiveAmount                int    `json:"EffectiveAmount"`
			NetPriceAmount                 int    `json:"NetPriceAmount"`
			NetPriceQuantity               int    `json:"NetPriceQuantity"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			PurchaseRequisition            string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
			RequestForQuotation            string `json:"RequestForQuotation"`
			RequestForQuotationItem        string `json:"RequestForQuotationItem"`
			PurchasingInfoRecordUpdateCode string `json:"PurchasingInfoRecordUpdateCode"`
			PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
			PurchasingDocumentItemCategory string `json:"PurchasingDocumentItemCategory"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
		} `json:"SupplierQuotationItem"`
	} `json:"SupplierQuotation"`
	APISchema             string   `json:"api_schema"`
	Accepter              []string `json:"accepter"`
	SupplierQuotationNo   string   `json:"supplier_quotation_no"`
	Deleted               bool     `json:"deleted"`
}
