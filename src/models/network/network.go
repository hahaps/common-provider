// @Author Bruce<lixipengs@qq.com>

package network

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// NetworkModel, , Default Cloud Network model
type NetworkModel struct {
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
    // Network name
    Name string
    // Network type
    Type string
    // Network CIDR
    CIDR string
    // Network Status
    Status string
    // Network description
    Description string
    // Network tags
    Tags string
    // Network Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *NetworkModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *NetworkModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *NetworkModel)GetIndex() string {
    return m.Index
}

func (m *NetworkModel)GetChecksum() string {
    return m.Checksum
}

func (m *NetworkModel)CheckRequired() (bool, string) {
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

func NewNetworkModel() *NetworkModel {
    m := &NetworkModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, Status, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "CIDR", "Status", }

    return m
}
