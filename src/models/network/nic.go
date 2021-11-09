// @Author Bruce<lixipengs@qq.com>

package network

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// NicModel, , Default Cloud Nic model
type NicModel struct {
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
    // Network ID
    NetworkId string
    // Subnet ID
    SubnetId string
    // Nic Create Time
    CreateTime string
    // Nic name
    Name string
    // Nic type
    Type string
    // Nic CIDR
    CIDR string
    // Nic protocol
    Protocol string
    // Nic Status
    Status string
    // Server Id bind to nic
    BindServer string
    // Network description
    Description string
    // Network tags
    InstanceId string
    Tags string
    // Network Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *NicModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *NicModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *NicModel)GetIndex() string {
    return m.Index
}

func (m *NicModel)GetChecksum() string {
    return m.Checksum
}

func (m *NicModel)CheckRequired() (bool, string) {
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

func NewNicModel() *NicModel {
    m := &NicModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, CIDR, Status, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "NetworkId", "SubnetId", "CIDR", "Protocol", "Status", }

    return m
}
