package qris

type TransferInfo struct {
	MerchantData      MerchantAccountDomesticInfo
	CentralRepository MerchantAccountDomesticInfo
	MerchantName      string
	MerchantCity      string
	Amount            int64
	Message           string
}

type MerchantAccountDomesticInfo struct {
	ReverseDomain string
	GlobalID      string
	ID            string
	Type          string
}
