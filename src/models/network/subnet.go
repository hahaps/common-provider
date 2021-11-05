// @Author Bruce<lixipengs@qq.com>

package network

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// SubnetModel, , Default Cloud Subnet model
type SubnetModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    
    // provider id
    ProviderId string
    // Cloud Provider Name
    CloudType string
    // Account ID
    AccountId string
    // Region ID
    RegionId string
    // Subnet name
    Name string
    // Subnet CIDR
    CIDR string
    // Subnet Status
    Status string
    // Subnet description
    Description string
    // Subnet tags
    Tags string
    // Subnet Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *SubnetModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *SubnetModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *SubnetModel)GetIndex() string {
    return m.Index
}

func (m *SubnetModel)GetChecksum() string {
    return m.Checksum
}

func (m *SubnetModel)CheckRequired() (bool, string) {
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

func NewSubnetModel() *SubnetModel {
    m := &SubnetModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, Status, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "CIDR", "Status", }

    return m
}
