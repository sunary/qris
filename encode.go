package qris

import (
	"strconv"
)

func Encode(ti TransferInfo) string {
	content := qrContent(ti)
	return content + hashCrc(content)
}

func qrContent(ti TransferInfo) string {
	version := genFieldData(_VERSION, "01")
	initMethod := genFieldData(_INIT_METHOD, "11")
	amount := ""

	if ti.Amount > 0 {
		initMethod = genFieldData(_INIT_METHOD, "12")
		amount = strconv.Itoa(int(ti.Amount))
	}

	providerReservedDomain := genFieldData(_PROVIDER_RESERVED_DOMAIN, ti.MerchantAccount.ReverseDomain)
	providerGlobalID := genFieldData(_PROVIDER_GUID, ti.MerchantAccount.GlobalID)
	providerID := genFieldData(_PROVIDER_ID, ti.MerchantAccount.ID)
	providerType := genFieldData(_PROVIDER_TYPE, ti.MerchantAccount.Type)
	providerData := genFieldData(_DOMESTIC1, joinString(providerReservedDomain, providerGlobalID, providerID, providerType))

	category := genFieldData(_CATEGORY, "")
	currency := genFieldData(_CURRENCY, "360")
	amountStr := genFieldData(_AMOUNT, amount)

	tipAndFeeType := genFieldData(_TIP_AND_FEE_TYPE, "")
	tipAndFeeAmount := genFieldData(_TIP_AND_FEE_AMOUNT, "")
	tipAndFeePercent := genFieldData(_TIP_AND_FEE_PERCENT, "")
	nation := genFieldData(_NATION, "ID")
	merchantName := genFieldData(_MERCHANT_NAME, ti.MerchantName)
	city := genFieldData(_CITY, ti.MerchantCity)
	zipCode := genFieldData(_ZIP_CODE, "")

	purpose := genFieldData(_PURPOSE_OF_TRANSACTION, ti.Message)
	//joinString(buildNumber, mobileNumber, storeLabel, loyaltyNumber, reference, customerLabel, terminal, purpose, dataRequest)
	additionalData := genFieldData(_ADDITIONAL_DATA, purpose)

	EVMCoContent := ""
	unreservedContent := ""

	return joinString(version, initMethod, providerData, category, currency, amountStr, tipAndFeeType, tipAndFeeAmount, tipAndFeePercent,
		nation, merchantName, city, zipCode, additionalData, EVMCoContent, unreservedContent, _CRC, "04")
}
