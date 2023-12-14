package util

import (
	"testing"
)

func TestGetAddressFromSignTransaction(t *testing.T) {
	cases := []struct {
		name, data, signed, userAddr string
	}{
		{"阿博personal_sign", "Example `personal_sign` message", "0xfd361b933045e5b4d0c96243ed06df008eeed7526d79c7702f046bdce26c9df52f8c80733bc9c8e346f9421909e832954fbbf1a4d992ce497801282044a19cf41c", "0xc17c970f31850bbc34e732b09ab3983c34d4f9cf"},
		{"opensea", "Welcome to OpenSea!\n\nClick to sign in and accept the OpenSea Terms of Service: https://opensea.io/tos\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x55808567522ee593a7e107a5a9e31c270388fb37\n\nNonce:\n56fd5698-5afb-4755-b3df-c68b2381d145", "0x3a852d40e31f5495ba6281daf4d6fe6ac5f0de40441ef75d2545a90d7e535f55740cc102d358624d06aa397e960bb58c2252632bac5c1fc3595dc32fa682d7bf1c", "0x55808567522EE593A7E107A5A9E31C270388FB37"},
	}

	for _, c := range cases {
		if got := VerifySig(c.userAddr, c.signed, []byte(c.data)); !got {
			t.Fatalf("%s 测试失败", c.name)
		}
	}
}
