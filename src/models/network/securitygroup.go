// @Author Bruce<lixipengs@qq.com>

package network

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// SecurityModel, , Default Cloud security group model
type SecurityGroupModel struct {
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
    // SecurityGroup group name
    Name string
    // SecurityGroup group description
    Description string
    // SecurityGroup group tags
    Tags string
    // SecurityGroup group Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *SecurityGroupModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *SecurityGroupModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *SecurityGroupModel)GetIndex() string {
    return m.Index
}

func (m *SecurityGroupModel)GetChecksum() string {
    return m.Checksum
}

func (m *SecurityGroupModel)CheckRequired() (bool, string) {
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

func NewSecurityGroupModel() *SecurityGroupModel {
    m := &SecurityGroupModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", }

    return m
}
