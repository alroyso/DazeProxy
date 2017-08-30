package encryption

import (
	"strings"
	"net"
)
type Action interface {
	//Init，代理实例初始化时执行的操作
	//param：配置文件里面填写的EncryptionParam
	//server：此代理实例中为加密模块预留的空间
	Init(param string,	server *interface{})(error)

	//InitUser，用户连接后进行的初始化操作
	//conn：用户的连接套接字
	//client：此用户对象中为加密模块预留的空间
	//server：此代理实例中为加密模块预留的空间
	InitUser(conn net.Conn,	client *interface{},	server *interface{})(error)

	//Encrypt，加密
	//client与server同上
	//data：源数据
	//输出 加密后的数据 与 一个error(若发生了错误)
	Encrypt(client *interface{},	server *interface{},	data []byte) ([]byte,error)

	//Decrypt，解密
	//client与server同上
	//data：加密数据
	//输出 解密后的数据 与 一个error(若发生了错误)
	Decrypt(client *interface{},	server *interface{},	data []byte) ([]byte,error)
}
type regfunc func()(Action)
var encryptionMap map[string]regfunc

func GetEncryption(name string) (regfunc,bool){
	name=strings.ToLower(name)
	d,flag:=encryptionMap[name]
	return d,flag
}

func init(){
	encryptionMap=make(map[string]regfunc)

	//自己开发的加密模块必需在此注册
	encryptionMap["none"]=func()(Action){
		return Action(&none{})
	}

}

//none-无加密
type none struct {
	RegArg string
}
func (this *none)Init(param string,server *interface{})(error){
	return nil
}
func (this *none)InitUser(conn net.Conn,client *interface{},server *interface{})(error){
	return nil
}
func (this *none)Encrypt(client *interface{},server *interface{},data []byte)([]byte,error){
	return data,nil
}
func (this *none)Decrypt(client *interface{},server *interface{},data []byte)([]byte,error){
	return data,nil
}
