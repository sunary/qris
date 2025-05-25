package qris

const (
	// headers
	_VERSION                     = "00"
	_INIT_METHOD                 = "01"
	_DOMESTIC_PAYMENT1           = "26"
	_DOMESTIC_PAYMENT2           = "45"
	_DOMESTIC_CENTRAL_REPOSITORY = "51"
	_CATEGORY                    = "52"
	_CURRENCY                    = "53" // defined by ISO 4217, Indonesian Rupiah (IDR) is 360
	_AMOUNT                      = "54"
	_TIP_AND_FEE_TYPE            = "55"
	_TIP_AND_FEE_AMOUNT          = "56"
	_TIP_AND_FEE_PERCENT         = "57"
	_NATION                      = "58" // defined by ISO 3166-1 alpha_2, Indonesia is ID
	_MERCHANT_NAME               = "59"
	_CITY                        = "60"
	_ZIP_CODE                    = "61"
	_ADDITIONAL_DATA             = "62"
	_CRC                         = "63"

	// provider headers
	_PROVIDER_RESERVED_DOMAIN = "00"
	_PROVIDER_GUID            = "01"
	_PROVIDER_ID              = "02"
	_PROVIDER_TYPE            = "03"

	// additional headers
	_BILL_NUMBER                      = "01" // BillNumber
	_MOBILE_NUMBER                    = "02" // MobileNumber
	_STORE_LABEL                      = "03" // StoreLabel
	_LOYALTY_NUMBER                   = "04" // LoyaltyNumber
	_REFERENCE_LABEL                  = "05" // ReferenceLabel
	_CUSTOMER_LABEL                   = "06" // CustomerLabel
	_TERMINAL_LABEL                   = "07" // TerminalLabel
	_PURPOSE_OF_TRANSACTION           = "08" // PurposeOfTransaction
	_ADDITIONAL_CONSUMER_DATA_REQUEST = "09" // AdditionalConsumerDataRequest
	_MERCHANT_TAX_ID                  = "10" // MerchantTaxID
	_MERCHANT_CHANNEL                 = "11" // MerchantChannel
)
