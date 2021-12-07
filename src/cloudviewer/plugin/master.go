package plugin

import (
	"errors"
	"github.com/natefinch/pie"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"runtime"
)

type ProviderMaster struct {
	Path    string
	setting map[string]interface{}
	client  *rpc.Client
}

type QueryParam struct {
	Setting       map[string]interface{}
	QueryString   string
}

func (pm *ProviderMaster) Init(setting map[string]interface{}) error {
	if runtime.GOOS == "windows" {
		pm.Path = pm.Path + ".exe"
	}
	fInfo, err := os.Stat(pm.Path)
	if err != nil {
		return err
	}
	if fInfo.IsDir() {
		return errors.New("must be file, not dir")
	}
	pm.setting = setting
	return nil
}

func (pm *ProviderMaster)Auth() (replay bool, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return replay, err
	}
	pm.client = client
	defer pm.client.Close()
	err = pm.client.Call("Store.Auth",
		pm.setting, &replay)
	return replay, err
}

func (pm *ProviderMaster) Query(qs string) (replay interface{}, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return replay, err
	}
	pm.client = client
	qp := QueryParam{
		Setting: pm.setting,
		QueryString: qs,
	}
	defer pm.client.Close()
	err = pm.client.Call("Store.Query",
		qp, &replay)
	return replay, err
}

