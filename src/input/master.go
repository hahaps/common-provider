package input

import (
	"errors"
	"fmt"
	"github.com/natefinch/pie"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"runtime"
)

type ProviderMaster struct {
	PluginPath string
	credential Credential
	client     *rpc.Client
}

func (pm *ProviderMaster) Init(credential Credential) error {
	if runtime.GOOS == "windows" {
		pm.PluginPath = pm.PluginPath + ".exe"
	}
	fInfo, err := os.Stat(pm.PluginPath)
	if err != nil {
		return err
	}
	if fInfo.IsDir() {
		return errors.New("must be file, not dir")
	}
	pm.credential = credential
	return nil
}

func (pm *ProviderMaster) CheckVersion(version string)(matched bool, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.PluginPath)
	if err != nil {
		return false, err
	}
	pm.client = client
	defer pm.client.Close()
	err = pm.client.Call("Version.Check",
		version, &matched)
	return matched, err
}

func (pm *ProviderMaster) Run(resource string, args map[string]interface{}) (replay Replay, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.PluginPath)
	if err != nil {
		return replay, err
	}
	pm.client = client
	params := Params{
		Credential: pm.credential,
		Args: args,
	}
	defer pm.client.Close()
	err = pm.client.Call(fmt.Sprintf("%v.Call", resource),
		params, &replay)
	return replay, err
}
