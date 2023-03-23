package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)
func PKCS5Padding(plainText []byte, blockSize int) []byte{
	padding := blockSize - (len(plainText)%blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func PKCS5UnPadding(plainText []byte)([]byte,error){
	length := len(plainText)
	number:= int(plainText[length-1])
	if number>length{
		return nil,errors.New("padding size error please check the secret key or iv")
	}
	return plainText[:length-number],nil
}

func EcbEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return decrypted, nil
}

func main()  {


	//str:=[]byte("fd7d726beae61bd45c43614611bf953d")

	test, _ := hex.DecodeString("a6ecb4152437d343446108c7e89fafc7")
	res, err := EcbEncrypt(test, []byte("823b967c52e5e04a3df013f1ed2d614a"))
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(strings.Trim(string(res)," "))
}