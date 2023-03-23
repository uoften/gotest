package main

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

func main(){
	plaintext := []byte("maipu_Runlian@2012") //明文
	fmt.Println("明文为：", string(plaintext))

	// 传入明文和自己定义的密钥，密钥和IV都为8字节
	cryptText, err := goEncrypt.DesCbcEncrypt(plaintext, []byte("asc@1234"), []byte("iv@12345")) //得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，8字节 如果解密秘钥错误解密后的明文会为空
	newplaintext, err := goEncrypt.DesCbcDecrypt(cryptText, []byte("asc@1234"), []byte("iv@12345")) //解密得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DES的CBC模式解密完：", string(newplaintext))



	hosykey, err := goEncrypt.AesCtrDecrypt([]byte("0ddcf458013a9bccee6ac3c98d6860e7"),[]byte("823b967c52e5e04a3df013f1ed2d614a"))
	fmt.Println("AES的CBC模式解密完：", string(newplaintext))
}