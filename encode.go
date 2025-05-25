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

	providerReservedDomain := genFieldData(_PROVIDER_RESERVED_DOMAIN, ti.MerchantData.ReverseDomain)
	providerGlobalID := genFieldData(_PROVIDER_GUID, ti.MerchantData.GlobalID)
	providerID := genFieldData(_PROVIDER_ID, ti.MerchantData.ID)
	providerType := genFieldData(_PROVIDER_TYPE, ti.MerchantData.Type)
	providerData := genFieldData(_DOMESTIC_PAYMENT1, joinString(providerReservedDomain, providerGlobalID, providerID, providerType))

	crReservedDomain := genFieldData(_PROVIDER_RESERVED_DOMAIN, ti.CentralRepository.ReverseDomain)
	crGlobalID := genFieldData(_PROVIDER_GUID, ti.CentralRepository.GlobalID)
	crID := genFieldData(_PROVIDER_ID, ti.CentralRepository.ID)
	crType := genFieldData(_PROVIDER_TYPE, ti.CentralRepository.Type)
	centralRepositoryData := genFieldData(_DOMESTIC_CENTRAL_REPOSITORY, joinString(crReservedDomain, crGlobalID, crID, crType))

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

	return joinString(version, initMethod, providerData, centralRepositoryData, category, currency, amountStr, tipAndFeeType, tipAndFeeAmount, tipAndFeePercent,
		nation, merchantName, city, zipCode, additionalData, EVMCoContent, unreservedContent, _CRC, "04")
}
