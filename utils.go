package alipay

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// 字符转公共key
func ConvertStrToPKCS1PublicKey(s string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(s))
	key, _ := x509.ParsePKIXPublicKey(block.Bytes)
	switch pub := key.(type) {
	case *rsa.PublicKey:
		return pub
	}
	return &rsa.PublicKey{}
}

// 字符转私有key
func ConvertStrToPKCS1PrivateKey(s string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(s))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}
