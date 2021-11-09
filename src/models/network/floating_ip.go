// @Author Bruce<lixipengs@qq.com>

package network

import (
	"reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// NetworkModel, , Default Cloud Network model
type FloatingIpModel struct {
	IndexKeys string
	ChecksumKeys string
	Index string
	Checksum string
	required []string
	Deleted int64
	// provider id
	ProviderId string
	// Cloud Provider Name
	CloudType string
	// Account ID
	AccountId string
	// Region ID
	RegionId string
	// FIP protocol
	Protocol string
	// FIP Create Time
	CreateTime string
	// FIP expired time
	ExpiredTime string
	// FIP bandwidth
	Bandwidth string
	// FIP CIDR
	IpAddr string
	// FIP Status
	Status string
	// Bind Resource Type
	BindResourceType string
	// Bind Resource Id
	BindResourceId string
	// FIP description
	Description string
	// FIP tags
	Tags string
	// FIP Extra info
	Extra map[string]interface{}
	model.BaseModel
}

func (m *FloatingIpModel)SetIndex() {
	m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *FloatingIpModel)SetChecksum() {
	val := model.GetValFromKeys(m, m.ChecksumKeys)
	m.Checksum = model.Checksum(&val)
}

func (m *FloatingIpModel)GetIndex() string {
	return m.Index
}

func (m *FloatingIpModel)GetChecksum() string {
	return m.Checksum
}

func (m *FloatingIpModel)CheckRequired() (bool, string) {
	ref := reflect.ValueOf(m).Elem()
	for _, key := range m.required {
		if ref.Kind() == reflect.String {
			if ref.FieldByName(key).String() == "" {
				return false, key
			}
		}
	}
	return true, ""
}

func NewFloatingIpModel() *FloatingIpModel {
	m := &FloatingIpModel{}
	m.IndexKeys = "CloudType, AccountId, ProviderId"
	m.ChecksumKeys = "Status, BindResourceType, BindResourceId"
	m.required = []string{"ProviderId", "CloudType", "AccountId", "IpAddr", "Status", }

	return m
}
