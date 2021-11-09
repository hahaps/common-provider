// @Author Bruce<lixipengs@qq.com>

package compute

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// ImageModel, , Default Cloud self defined image
type ImageModel struct {
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
    // Image name
    Name string
    // Image OS type
    // Image Create Time
    CreateTime string
    OSType string
    // Image size
    Size int32
    // Image description
    Description string
    // Image Status
    Status string
    // Image tags
    Tags string
    // Image Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *ImageModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *ImageModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *ImageModel)GetIndex() string {
    return m.Index
}

func (m *ImageModel)GetChecksum() string {
    return m.Checksum
}

func (m *ImageModel)CheckRequired() (bool, string) {
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

func NewImageModel() *ImageModel {
    m := &ImageModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, Status, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "Status", }

    return m
}
