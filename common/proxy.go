package common

import (
	"github.com/crabkun/DazeProxy/obscure"
	"github.com/crabkun/DazeProxy/encryption"
)

type S_config struct{
	Debug bool
	Proxy []S_proxy
	DatabaseDriver string
	DatabaseConnectionString string
	NoAuth bool
}
type S_proxy struct{
	Port string

	//加密方式与参数
	Encryption string
	EncryptionParam string

	//伪装方式与参数
	Obscure string
	ObscureParam string

	//服务器所属组
	Group string

	//加密与伪装的接口
	Ob obscure.Action `json:"-"`
	E encryption.Action `json:"-"`
	EReserved interface{} `json:"-"`
	ObReserved interface{} `json:"-"`
	Config S_config `json:"-"`
}