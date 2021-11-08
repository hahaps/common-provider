package output

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

func (pm *ProviderMaster) CheckVersion(version string)(matched bool, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return matched, err
	}
	pm.client = client
	defer pm.client.Close()
	err = pm.client.Call("Version.Check",
		version, &matched)
	return matched, err
}

func (pm *ProviderMaster) Run(params Params) (replay Replay, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return replay, err
	}
	pm.client = client
	params.Setting = pm.setting
	defer pm.client.Close()
	err = pm.client.Call("Store.Push",
		params, &replay)
	return replay, err
}

func (pm *ProviderMaster) UpdateSyncJob(params JobParams) (jobId string, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return jobId, err
	}
	pm.client = client
	params.Setting = pm.setting
	defer pm.client.Close()
	err = pm.client.Call("Store.UpdateSyncJob",
		params, &jobId)
	return jobId, err
}

func (pm *ProviderMaster) UpdateDeleted(params Params) (replay int32, err error) {
	client, err := pie.StartProviderCodec(
		jsonrpc.NewClientCodec, os.Stderr, pm.Path)
	if err != nil {
		return replay, err
	}
	pm.client = client
	params.Setting = pm.setting
	defer pm.client.Close()
	err = pm.client.Call("Store.UpdateDeleted",
		params, &replay)
	return replay, err
}
