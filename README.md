# qris

Lightweight encode/decode of QRIS (QR Code Indonesian Standard).


## Sample

```go
package main

import (
	"github.com/sunary/qris"
)

func main() {
	println("encode", qris.Encode(qris.TransferInfo{
		MerchantData: qris.MerchantAccountDomesticInfo{
			ReverseDomain: "ID.CO.SHOPEE.WWW",
			GlobalID:      "01",
			ID:            "15863724",
		},
		MerchantName: "PT. SHOPPEE INDONESIA",
		MerchantCity: "JAKARTA",
		Amount:       144000,
		Message:      "gen by go-qris",
	}))

	info, err := qris.Decode("00020101021226380016ID.CO.SHOPEE.WWW010201020815863724530336054061440005802ID5921PT. SHOPPEE INDONESIA6007JAKARTA62180814gen by go-qris63049635")
	if err != nil {
		println("err", err.Error())
		return
	}

	println("merchant reverse domain", info.MerchantData.ReverseDomain)
	println("merchant global id", info.MerchantData.GlobalID)
	println("merchant id", info.MerchantData.ID)
	println("merchant type", info.MerchantData.Type)
	println("merchant name", info.MerchantName)
	println("merchant city", info.MerchantCity)
	println("amount", info.Amount)
	println("message", info.Message)
}
```

Run online: [go.dev/play](https://go.dev/play/p/RlnMhfEcwMW)