// @Author Bruce<lixipengs@qq.com>

package models

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// MetricModel, , Default Metric Model
type MetricModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    Deleted int64
    // Metric Name
    Name string
    // Metric unit
    Unit string
    // Metric Value
    Value string
    // Cloud Provider Name
    CloudType string
    // Cloud AccountId
    AccountId string
    // Instance ID
    InstanceId string
    // Cloud Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *MetricModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *MetricModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *MetricModel)GetIndex() string {
    return m.Index
}

func (m *MetricModel)GetChecksum() string {
    return m.Checksum
}

func (m *MetricModel)CheckRequired() (bool, string) {
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

func NewMetricModel() *MetricModel {
    m := &MetricModel{}
    m.IndexKeys = "Id"
    m.ChecksumKeys = "Name, Type"
    m.required = []string{"Name", "Unit", "Value", "CloudType", "AccountId", "InstanceId", }

    return m
}
