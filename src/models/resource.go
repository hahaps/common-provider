// @Author Bruce<lixipengs@qq.com>

package models

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// ResourceModel, , Default Cloud Resource Model
type ResourceModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    
    // provider id
    ProviderId string
    // Cloud Provider Name
    CloudType string
    // Resource name
    Name string
    // Product type
    ProductType string
    // Product code
    ProductCode string
    // Create time
    CreateTime string
    // resource tags
    Tags string
    // Extra Info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *ResourceModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *ResourceModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *ResourceModel)GetIndex() string {
    return m.Index
}

func (m *ResourceModel)GetChecksum() string {
    return m.Checksum
}

func (m *ResourceModel)CheckRequired() (bool, string) {
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

func NewResourceModel() *ResourceModel {
    m := &ResourceModel{}
    m.IndexKeys = "CloudType, ProviderId"
    m.ChecksumKeys = "Name, Tags"
    m.required = []string{"ProviderId", "CloudType", "ProductType", }

    return m
}
