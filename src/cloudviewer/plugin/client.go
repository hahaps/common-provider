package plugin

import (
	"github.com/natefinch/pie"
	"net/rpc/jsonrpc"
)

type Provider interface {
	Auth(setting map[string]interface{}, replay *bool) (err error)
	Query(qp QueryParam, replay *interface{}) (err error)
}

func RunProvider(store Provider) error {
	server := pie.NewProvider()
	if err := server.RegisterName("Store", store); err != nil {
		return err
	}
	server.ServeCodec(jsonrpc.NewServerCodec)
	return nil
}
