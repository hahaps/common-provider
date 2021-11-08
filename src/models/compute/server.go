// @Author Bruce<lixipengs@qq.com>

package compute

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// ServerModel, , Default Cloud Server Model
type ServerModel struct {
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
    // Server name
    Name string
    // Resource create time
    CreateTime string
    // Resource Expire time
    ExpireTime string
    // Resource status
    Status string
    // Pay mode(PayAsYouGo, Subscription)
    PayMode string
    // Auto renew state
    AutoRenew bool
    // Flavor Id
    FlavorId string
    // Flavor Name
    FlavorName string
    // Core Number of vCPUs
    FlavorVCPU int
    // Ram size
    FlavorRam int
    // Flavor Extra info
    FlavorExtra map[string]interface{}
    // Image Id
    ImageId string
    // Image Name
    ImageName string
    // Image OS type
    ImageOsType string
    // Image Extra info
    ImageExtra map[string]interface{}
    // IP addr of primary nic
    PrimaryNicIp string
    // IP protocol of to primary nic
    PrimaryNicIpProtocol string
    // Floating IP bind to primary nic
    PrimaryNicFloatingIp string
    // Floating IP protocol bind to primary nic
    PrimaryNicFloatingIpProtocol string
    // Network of primary nic
    PrimaryNetworkId string
    // Subnet of primary nic
    PrimarySubnetId string
    // Secondary nics info
    SecondaryNics map[string]interface{}
    // Security group info
    SecurityGroups map[string]string
    // Server tags
    Tags string
    // Connection URL string
    Extra map[string]interface{}
    model.BaseModel
}

func (m *ServerModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *ServerModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *ServerModel)GetIndex() string {
    return m.Index
}

func (m *ServerModel)GetChecksum() string {
    return m.Checksum
}

func (m *ServerModel)CheckRequired() (bool, string) {
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

func NewServerModel() *ServerModel {
    m := &ServerModel{}
    m.IndexKeys = "CloudType, AccountId, ProviderId"
    m.ChecksumKeys = "Name, ExpireTime, Status, PayMode, AutoRenew, FlavorId, FlavorName, FlavorVCPU, FlavorRam, FlavorExtra, SystemDiskSize, DataDisks, PrimaryNicFloatingIp, SecondaryNics, SecurityGroups, Tags"
    m.required = []string{"ProviderId", "CloudType", "AccountId", "FlavorId", "FlavorVCPU", "FlavorRam", }

    m.PrimaryNicIpProtocol = "ipv4"
    m.PrimaryNicFloatingIpProtocol = "ipv4"
    return m
}
