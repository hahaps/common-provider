package input

import (
	"github.com/natefinch/pie"
	"net/rpc/jsonrpc"
)

type Credential struct {
	SecretId   string                 `yaml:"secret_id"`
	SecretKey  string                 `yaml:"secret_key"`
	AccountId  string                 `yaml:"account_id"`
	ConnectURL string                 `yaml:"connect_url"`
	Extra      map[string]interface{} `yaml:"extra"`
}

type Replay struct {
	Result []interface{}
	Next   string
	Query  map[string]interface{}
}

type Resource interface {
	Call(params Params, replay *Replay) error
}

type Params struct {
	Timestamp  int64
	Credential Credential
	Args       map[string]interface{}
}

type Version interface {
	Check(version string, matched *bool) error
}

func RunProvider(version Version, resources map[string]Resource) error {
	server := pie.NewProvider()
	if err := server.RegisterName("Version", version); err != nil {
		return err
	}
	for name, resource := range resources {
		if err := server.RegisterName(name, resource); err != nil {
			return err
		}
	}
	server.ServeCodec(jsonrpc.NewServerCodec)
	return nil
}
