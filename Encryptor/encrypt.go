package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"

	"crypto/rand"

	"github.com/wenzhenxi/gorsa"
)

// 公钥加密私钥解密
func rsa_encrypt(data []byte) ([]byte, error) {
	gRsa := gorsa.RSASecurity{}
	gRsa.SetPublicKey(byte_decode2str(PUBKEY[:]))

	rsaData, err := gRsa.PubKeyENCTYPT([]byte(data))
	if err != nil {
		return []byte{}, err
	}
	return rsaData, nil
}

func rsa_decrypt(data []byte) ([]byte, error) {
	gRsa := gorsa.RSASecurity{}

	if err := gRsa.SetPrivateKey(byte_decode2str(PRIKEY[:])); err != nil {
		return []byte{}, err
	}

	rsaData, err := gRsa.PriKeyDECRYPT(data)
	if err != nil {
		return []byte{}, err
	}
	return rsaData, nil
}

func aes_pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// AES加密
func aes_encrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKEY[:])
	if err != nil {
		return nil, err
	}

	msg := aes_pad(text)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], msg)

	return ciphertext, nil
}

func aes_unpad(src []byte) ([]byte, error) {
	padding := src[len(src)-1]
	if int(padding) > len(src) {
		return nil, errors.New("padding size is invalid")
	}
	return src[:len(src)-int(padding)], nil
}

// AES解密
func aes_decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKEY[:])
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return aes_unpad(ciphertext)
}

func byte_decode2str(data []byte) string {
	// 字节转换为 utf-8 字符串
	return string(bytes.SplitN(data[:], []byte{0}, 2)[0])
}
