package util

import (
	"encoding/pem"
	"crypto/x509"
	"os"
	"crypto/rsa"
	"crypto/rand"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"crypto/aes"
)

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
func CheckKeyAndGen()  {
	fmt.Println("[!]开始检查密钥文件")
	_,privateKeyErr:=os.Stat("private.pem")
	_,publicKeyErr:=os.Stat("public.pem")
	if privateKeyErr!=nil || publicKeyErr!=nil{
		fmt.Println("[!]密钥文件不存在，开始生成密钥文件")
		os.Remove("private.pem")
		os.Remove("public.pem")
		GenRsaKey(1024)
		fmt.Println("[√]密钥文件生成成功")
	}
	fmt.Println("[√]密钥文件检查完毕")
}
func CheckLicense(){

}
func GenAESKey(bit int) []byte{
	buf:=make([]byte,bit)
	rand.Read(buf)
	return buf
}
func DecryptRSA(data []byte) ([]byte,error){
	KeyFileBuf,PrivateKeyErr:=ioutil.ReadFile("private.pem")
	if PrivateKeyErr!=nil{
		fmt.Println("[×D]私钥文件丢失！！系统强制退出")
		os.Exit(-1)
	}
	block,_:=pem.Decode(KeyFileBuf)
	PrivateKey,PrivateKeyParseErr:=x509.ParsePKCS1PrivateKey(block.Bytes)
	if PrivateKeyParseErr!=nil{
		fmt.Println("[×D]私钥文件解析错误！！系统强制退出")
		os.Exit(-1)
	}
	return rsa.DecryptPKCS1v15(rand.Reader,PrivateKey,data)
}
func EncryptRSA(data []byte) ([]byte,error){
	KeyFileBuf,PrivateKeyErr:=ioutil.ReadFile("private.pem")
	if PrivateKeyErr!=nil{
		fmt.Println("[×E]私钥文件丢失！！系统强制退出")
		os.Exit(-1)
	}
	block,_:=pem.Decode(KeyFileBuf)
	PrivateKey,PrivateKeyParseErr:=x509.ParsePKCS1PrivateKey(block.Bytes)
	if PrivateKeyParseErr!=nil{
		fmt.Println("[×E]私钥文件解析错误！！系统强制退出",PrivateKeyParseErr.Error())
		os.Exit(-1)
	}
	return rsa.EncryptPKCS1v15(rand.Reader,&PrivateKey.PublicKey,data)
}
func DecryptAES(data []byte,key []byte) ([]byte,error){
	block,CipherErr:=aes.NewCipher(key)
	Decrypter:=cipher.NewCFBDecrypter(block,key[:block.BlockSize()])
	decoded:=make([]byte,len(data))
	Decrypter.XORKeyStream(decoded,data)
	return decoded,CipherErr
}
func EncryptAES(data []byte,key []byte) ([]byte,error){
	block,CipherErr:=aes.NewCipher(key)
	Decrypter:=cipher.NewCFBEncrypter(block,key[:block.BlockSize()])
	encoded:=make([]byte,len(data))
	Decrypter.XORKeyStream(encoded,data)
	return encoded,CipherErr
}