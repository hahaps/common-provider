// @Author Bruce<lixipengs@qq.com>

package storage

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// DiskModel, , Default Cloud Disk model
type DiskModel struct {
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
    // Disk name
    Name string
    // Disk create time
    CreateTime string
    // Expired time
    ExpiredTime string
    // Disk type
    Type string
    // Category
    Category string
    // Disk size
    Size int32
    // Disk Status
    Status string
    // Disk Attach Status
    Attachments []map[string]interface{}
    // Attached Server Id
    AttachedServer string
    // Disk description
    Description string
    // Disk tags
    Tags string
    // Disk Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *DiskModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *DiskModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *DiskModel)GetIndex() string {
    return m.Index
}

func (m *DiskModel)GetChecksum() string {
    return m.Checksum
}

func (m *DiskModel)CheckRequired() (bool, string) {
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

func NewDiskModel() *DiskModel {
    m := &DiskModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, Size, Status, Attachments, AttachedServer, Description, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "Size", "Status", }

    return m
}
