package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// 字符转公共key
func ConvertStrToPKCS1PublicKey(s string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(s))

	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// 字符转私有key
func ConvertStrToPKCS1PrivateKey(s string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(s))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
