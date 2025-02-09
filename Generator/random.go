package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func getRSAPublicPrivate() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", fmt.Errorf("生成 RSA 密钥失败: %v", err)
	}

	privDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", "", fmt.Errorf("私钥转换为 PKCS#8 失败: %v", err)
	}
	privBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privDER,
	}
	privateKeyPEM := string(pem.EncodeToMemory(privBlock))

	pubDER, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", fmt.Errorf("公钥转换失败: %v", err)
	}
	pubBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	}
	publicKeyPEM := string(pem.EncodeToMemory(pubBlock))

	return publicKeyPEM, privateKeyPEM, nil
}

func getAES(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}
