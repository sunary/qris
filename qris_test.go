package qris

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		ti TransferInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "only beneficiary info",
			args: args{
				TransferInfo{
					MerchantData: MerchantAccountDomesticInfo{
						ReverseDomain: "ID.CO.SHOPEE.WWW",
						GlobalID:      "01",
						ID:            "1893600918",
					},
				},
			},
			want: "00020101021126400016ID.CO.SHOPEE.WWW0102010210189360091853033605802ID630492B2",
		},
		{
			name: "with amount",
			args: args{
				TransferInfo{
					MerchantData: MerchantAccountDomesticInfo{
						ReverseDomain: "ID.CO.SHOPEE.WWW",
						GlobalID:      "01",
						ID:            "1893600918",
					},
					Amount: 120000,
				},
			},
			want: "00020101021226400016ID.CO.SHOPEE.WWW01020102101893600918530336054061200005802ID63040EE3",
		},
		{
			name: "with message",
			args: args{
				TransferInfo{
					MerchantData: MerchantAccountDomesticInfo{
						ReverseDomain: "ID.CO.SHOPEE.WWW",
						GlobalID:      "01",
						ID:            "1893600918",
					},
					Message: "gen by sunary/qris",
				},
			},
			want: "00020101021126400016ID.CO.SHOPEE.WWW0102010210189360091853033605802ID62220818gen by sunary/qris6304BB82",
		},
		{
			name: "full info",
			args: args{
				TransferInfo{
					MerchantData: MerchantAccountDomesticInfo{
						ReverseDomain: "ID.CO.SHOPEE.WWW",
						GlobalID:      "01",
						ID:            "1893600918",
					},
					CentralRepository: MerchantAccountDomesticInfo{
						ReverseDomain: "ID.CO.QRIS.WWW",
						GlobalID:      "01",
						ID:            "20250525",
						Type:          "UBE",
					},
					MerchantName: "PT. SHOPPEE INDONESIA",
					MerchantCity: "JAKARTA",
					Amount:       152000,
					Message:      "gen by go-qris",
				},
			},
			want: "00020101021226400016ID.CO.SHOPEE.WWW0102010210189360091851430014ID.CO.QRIS.WWW0102010208202505250303UBE530336054061520005802ID5921PT. SHOPPEE INDONESIA6007JAKARTA62180814gen by go-qris6304058E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.ti); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	equalTransferInfoFn := func(t1, t2 *TransferInfo) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		return t1.MerchantData.ReverseDomain == t2.MerchantData.ReverseDomain &&
			t1.MerchantData.GlobalID == t2.MerchantData.GlobalID &&
			t1.MerchantData.ID == t2.MerchantData.ID &&
			t1.MerchantData.Type == t2.MerchantData.Type &&
			t1.CentralRepository.ReverseDomain == t2.CentralRepository.ReverseDomain &&
			t1.CentralRepository.GlobalID == t2.CentralRepository.GlobalID &&
			t1.CentralRepository.ID == t2.CentralRepository.ID &&
			t1.CentralRepository.Type == t2.CentralRepository.Type &&
			t1.MerchantName == t2.MerchantName &&
			t1.MerchantCity == t2.MerchantCity &&
			t1.Amount == t2.Amount &&
			t1.Message == t2.Message
	}

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *TransferInfo
		wantErr bool
	}{
		{
			name: "wrong crc",
			args: args{
				"00020101021153033605802ID6304xxxx",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing crc",
			args: args{
				"00020101021153033605802ID6304",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct data: only beneficiary info",
			args: args{
				"00020101021126400016ID.CO.SHOPEE.WWW0102010210189360091853033605802ID630492B2",
			},
			want: &TransferInfo{
				MerchantData: MerchantAccountDomesticInfo{
					ReverseDomain: "ID.CO.SHOPEE.WWW",
					GlobalID:      "01",
					ID:            "1893600918",
				},
			},
			wantErr: false,
		},
		{
			name: "correct data: with amount",
			args: args{
				"00020101021226400016ID.CO.SHOPEE.WWW01020102101893600918530336054061200005802ID63040EE3",
			},
			want: &TransferInfo{
				MerchantData: MerchantAccountDomesticInfo{
					ReverseDomain: "ID.CO.SHOPEE.WWW",
					GlobalID:      "01",
					ID:            "1893600918",
				},
				Amount: 120000,
			},
			wantErr: false,
		},
		{
			name: "correct data: with message",
			args: args{
				"00020101021126400016ID.CO.SHOPEE.WWW0102010210189360091853033605802ID62220818gen by sunary/qris6304BB82",
			},
			want: &TransferInfo{
				MerchantData: MerchantAccountDomesticInfo{
					ReverseDomain: "ID.CO.SHOPEE.WWW",
					GlobalID:      "01",
					ID:            "1893600918",
				},
				Message: "gen by sunary/qris",
			},
			wantErr: false,
		},
		{
			name: "correct data: full info",
			args: args{
				"00020101021226400016ID.CO.SHOPEE.WWW0102010210189360091851430014ID.CO.QRIS.WWW0102010208202505250303UBE530336054061520005802ID5921PT. SHOPPEE INDONESIA6007JAKARTA62180814gen by go-qris6304058E",
			},
			want: &TransferInfo{
				MerchantData: MerchantAccountDomesticInfo{
					ReverseDomain: "ID.CO.SHOPEE.WWW",
					GlobalID:      "01",
					ID:            "1893600918",
				},
				CentralRepository: MerchantAccountDomesticInfo{
					ReverseDomain: "ID.CO.QRIS.WWW",
					GlobalID:      "01",
					ID:            "20250525",
					Type:          "UBE",
				},
				MerchantName: "PT. SHOPPEE INDONESIA",
				MerchantCity: "JAKARTA",
				Amount:       152000,
				Message:      "gen by go-qris",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalTransferInfoFn(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
