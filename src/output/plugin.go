package output

import (
	"github.com/natefinch/pie"
	"net/rpc/jsonrpc"
)

type Replay struct {
	Status int32
	Output interface{}
}

type Params struct {
	Timestamp int64
	Resource  string
	Setting   map[string]interface{}
	Input     []map[string]interface{}
	Query     map[string]interface{}
}

type SyncJobInfo struct {
	Index    string
	Resource string
	Type     string
	Value    string
	Status   string
	StartAt  string
	EndAt    string
}

type JobParams struct {
	Setting map[string]interface{}
	SyncJob SyncJobInfo
	Resource string
}

type Provider interface {
	Push(params Params, replay *Replay) (err error)
	UpdateDeleted(params Params, replay *int32) (err error)
	UpdateSyncJob(params JobParams, jobId *string) (err error)
}

type Version interface {
	Check(version string, matched *bool) error
}

func RunProvider(version Version, store Provider) error {
	server := pie.NewProvider()
	if err := server.RegisterName("Version", version); err != nil {
		return err
	}
	if err := server.RegisterName("Store", store); err != nil {
		return err
	}
	server.ServeCodec(jsonrpc.NewServerCodec)
	return nil
}
