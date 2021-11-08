// @Author Bruce<lixipengs@qq.com>

package models

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// RegionModel, , Default Cloud Region Model
type RegionModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    Deleted int64
    // Cloud Region Name
    Name string
    // Cloud Region ID
    Id string
    // Cloud Provider Name
    CloudType string
    // Cloud AccountId
    AccountId string
    // Cloud Region Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *RegionModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *RegionModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *RegionModel)GetIndex() string {
    return m.Index
}

func (m *RegionModel)GetChecksum() string {
    return m.Checksum
}

func (m *RegionModel)CheckRequired() (bool, string) {
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

func NewRegionModel() *RegionModel {
    m := &RegionModel{}
    m.IndexKeys = "CloudType, AccountId, Id"
    m.ChecksumKeys = "Name"
    m.required = []string{"Name", "Id", "CloudType", "AccountId", }

    return m
}
