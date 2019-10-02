package main

import "github.com/unliar/alipay"

func main() {

	c := alipay.Client{
		AppID: "2019092667839325",
		// AliPayPublicKey: alipay.ConvertStrToPKCS1PublicKey(AliPayPublicKey),
		// AppPrivateKey: alipay.ConvertStrToPKCS1PrivateKey(AppPrivateKey),
	}
	c.TradePreCreate()
}
