package qris

import (
	"errors"
	"strconv"
)

const paddingCrc = 4

func Decode(s string) (*TransferInfo, error) {
	if !validCrcContent(s) {
		return nil, errors.New("invalid CRC")
	}

	ti := TransferInfo{}
	ti.parseRootContent(s[:len(s)-paddingCrc])

	return &ti, nil
}

func (t *TransferInfo) parseRootContent(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _VERSION:
	case _INIT_METHOD:
	case _DOMESTIC1, _DOMESTIC2:
		t.parseProviderInfo(value)
	case _CATEGORY:
	case _CURRENCY:
	case _AMOUNT:
		t.Amount, _ = strconv.ParseInt(value, 10, 64)
	case _TIP_AND_FEE_TYPE:
	case _TIP_AND_FEE_AMOUNT:
	case _TIP_AND_FEE_PERCENT:
	case _NATION:
	case _MERCHANT_NAME:
		t.MerchantName = value
	case _CITY:
		t.MerchantCity = value
	case _ZIP_CODE:
	case _ADDITIONAL_DATA:
		t.parseAdditionalData(value)
	}

	if len(nextValue) > 4 {
		t.parseRootContent(nextValue)
	}
}

func (t *TransferInfo) parseProviderInfo(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _PROVIDER_RESERVED_DOMAIN:
		t.MerchantAccount.ReverseDomain = value
	case _PROVIDER_GUID:
		t.MerchantAccount.GlobalID = value
	case _PROVIDER_ID:
		t.MerchantAccount.ID = value
	case _PROVIDER_TYPE:
		t.MerchantAccount.Type = value
	}

	if len(nextValue) > 4 {
		t.parseProviderInfo(nextValue)
	}
}

func (t *TransferInfo) parseAdditionalData(s string) {
	id, value, nextValue := slideContent(s)
	switch id {
	case _PURPOSE_OF_TRANSACTION:
		t.Message = value
	case _BILL_NUMBER:
	case _MOBILE_NUMBER:
	case _REFERENCE_LABEL:
	case _STORE_LABEL:
	case _TERMINAL_LABEL:
	}

	if len(nextValue) > 4 {
		t.parseAdditionalData(nextValue)
	}
}
