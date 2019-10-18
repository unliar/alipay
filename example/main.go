package main

import (
	"fmt"
	"github.com/unliar/alipay"
	ustring "github.com/unliar/utils/go/string"
)

const AliPayPublicKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkquFFuOlLI4uV5gU4OdU8H+A6GZSx7kXjEtvIxVLQBNdAbfoEW9ZM0e3ECsvgwj2KCStXQk+qDRICVfl+6ebdXk9O/3ut3uonW/cvcMgADwv62LmKCW3gbT+uYPZzprBO5fr+wz6F9WqrOAT4T/oaKlN6WXFA8M9pmlqbqeoR0Y6Ge1Qlzavo5Oc58s/L+nFpp3dhW58xGa4ud6GsnIZ3Sp9HAIQt7ptILReeFyS80httCUobqhTKVOJyvElicCCMn/OMx1KlOJc3HlfH8bt2gqtTQTxStNRtkruodCdcOtc+Y9yepBwhDxPsHGF/4n0hq1kssgpWQb0vj5Y+1wRCQIDAQAB
-----END PUBLIC KEY-----

`
const AppPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAkquFFuOlLI4uV5gU4OdU8H+A6GZSx7kXjEtvIxVLQBNdAbfoEW9ZM0e3ECsvgwj2KCStXQk+qDRICVfl+6ebdXk9O/3ut3uonW/cvcMgADwv62LmKCW3gbT+uYPZzprBO5fr+wz6F9WqrOAT4T/oaKlN6WXFA8M9pmlqbqeoR0Y6Ge1Qlzavo5Oc58s/L+nFpp3dhW58xGa4ud6GsnIZ3Sp9HAIQt7ptILReeFyS80httCUobqhTKVOJyvElicCCMn/OMx1KlOJc3HlfH8bt2gqtTQTxStNRtkruodCdcOtc+Y9yepBwhDxPsHGF/4n0hq1kssgpWQb0vj5Y+1wRCQIDAQABAoIBAFM5yLeITX3O4DUMMyy7l9MwRrGY/ngea2JMm6/IsX6WfjwvYxwhlvgmRZaDKt6sAyIcTdNxH4DKyeWjXvlBzlEGjtyv5vluzesoQpXbVxlLDdX/tKIrZMnRLomOa9V0sxlhx6IZZwCHWeXm8ODsJrDdATzjZrwP2bfIMYDA2hg2wPSDwThf7mZrFQzF1kyt25J+KfoUblv+2+nPFkBbVkD86PVuKerL/XfdHDW25JjL1JdJ3p6YbLg1aUR1Qy8twQcjtwwFcqpNLYdd0E2yNDQhd4UdJlOOiOlh7RHaiC+AfS7tN8/lPB7DuQGCgxg062h8G50wGv4M9U5EgxbqBVECgYEAzxI04sIwkJ0IhlFyk8BO3nxLd64O+QngHstDObUgl2uYOwvcCH3hw8xT6BgSbphyIdRgT57UK+6N7zTNWQ6fIoL0BlPHM40oBX85mr5kaxRyJjFJowk0GA1oZ+4DB4Wa8q1Q+J7dWZB9R14+FBqHYKckMSxPYoRFC6U5VBWN5p8CgYEAtVOhsziq/HrC9kLko+3M5/0txaL0ulQ/49xpkhut7dvBRpciiiIR4QIG8XEJnBwpqkXTuwrfOjKGJrsMQhVL8VD0nzQyty4PCt28ST5stqkjzlpmAPVlpQy80BYw663ZgtYSUTw7lZbQaL4sCJtL9o08IMO2l94lwnhrI9MEr1cCgYEAqhdRe8Z5cACdxP9YN4erAVRmBUvjnqt+/qyGbvuaZucJp87pEcydS2Emtyo20cPFVIaICj70F5Yf3pKn7vR0wTuCSUQ+B9l2O3WzEqo8AD1OOpMX4qUntm4lCeHXeTFFAvxc26xbNDvcvGfsZEUaWMlSkFXOa6UsthElEy2VXw8CgYEApyTnSihzXPbgXlZ9IirjFfnIRoUW5+cfWbBkVD1Vj4thVuub+A69wla3Bbp37EH9mipxqNm1uZS3Gl6TRxsQfOpuA47/LOG1FgQdOrrjRWEWU3H60ulh/8mFBp9eCvGfLb6c9er61cJGbDbYqjpUxHmeMmmWF9m7ns6XnFc421MCgYB2CMmZsiVfjd1YVNSuNCxdEPDDkNmt7mZHW2l/SEmGHn7Oq20ZXrgJHZxuvH0T0ljcoPSbVLwJPRZpSW0pxh8tZNUMayiBVyPqxaXMPoQRqOcvpcrL8UZOCFO1Ja3XEk9nO57nId5ozs12OF6td3BcRmG8Hu38oUXCrrpNELYndw==
-----END RSA PRIVATE KEY-----
`

func main() {

	c := alipay.Client{
		AppID:           "20190992667839325",
		AliPayPublicKey: ustring.ConvertStrToPKCS1PublicKey(AliPayPublicKey),
		AppPrivateKey:   ustring.ConvertStrToPKCS1PrivateKey(AppPrivateKey),
		NotifyURL:       "https://happysooner.com/api/v1/pay/alipay/hook",
		SignType:        "RSA2",
		EndpointURL:     alipay.AlipayTradeEndpointURL,
	}
	res, err := c.TradePreCreate(alipay.BizContentRequestParams{
		OutTradeNo:  "99299dd392ssss392dddd93929iid",
		TotalAmount: "1.21",
		Subject:     "王一鸣首席前端工程师",
	})
	if err != nil {
		fmt.Println("请求失败", err)
		return
	}

	fmt.Println("信息", res.AlipayTradePrecreateResponse)
	if res.AlipayTradePrecreateResponse.Code == "10000" {
		fmt.Println("请求成功")
		fmt.Println("二维码地址", res.AlipayTradePrecreateResponse.QrCode)
		fmt.Println("商家订单", res.AlipayTradePrecreateResponse.OutTradeNo)
		return
	}
	fmt.Println("订单请求失败", res.AlipayTradePrecreateResponse.Code, res.AlipayTradePrecreateResponse.SubCode)
}
