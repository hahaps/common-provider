// @Author Bruce<lixipengs@qq.com>

package network

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// SGRuleModel, , Default Cloud Security Group Rule
type SGRuleModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    
    // Security group Id
    SGId string
    // Cloud Provider Name
    CloudType string
    // Account ID
    AccountId string
    // Region ID
    RegionId string
    // Rule Protocol
    Protocol string
    // Port range
    PortRange string
    // Rule CIDR
    CIDR string
    // Rule deraction
    Deraction string
    // Image Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *SGRuleModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *SGRuleModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *SGRuleModel)GetIndex() string {
    return m.Index
}

func (m *SGRuleModel)GetChecksum() string {
    return m.Checksum
}

func (m *SGRuleModel)CheckRequired() (bool, string) {
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

func NewSGRuleModel() *SGRuleModel {
    m := &SGRuleModel{}
    m.IndexKeys = "CloudType, AccountId, SGId, Description, CIDR, PortRange, Protocol"
    m.ChecksumKeys = "Description, CIDR, PortRange, Protocol"
    m.required = []string{"SGId", "CloudType", "AccountId", "Protocol", "PortRange", "CIDR", }

    return m
}
