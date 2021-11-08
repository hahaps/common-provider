// @Author Bruce<lixipengs@qq.com>

package models

import (
    "reflect"

	"github.com/hahaps/common-provider/src/common/model"
)

// InstanceBillModel, , Default Cloud InstanceBill Model
type InstanceBillModel struct {
    IndexKeys string
    ChecksumKeys string
    Index string
    Checksum string
    required []string
    Deleted int64
    // Billing cycle
    BillingCycle string
    // Billing date
    BillingDate string
    // Cloud Provider Name
    CloudType string
    // Cloud AccountId
    AccountId string
    // Cloud Resource ID
    InstanceId string
    // Cloud Resource ID
    InstanceName string
    // Subscription Type
    SubscriptionType string
    // Item action, such as New, ReNew, etc
    ItemAction string
    // Product Name
    ProductName string
    // Product Code
    ProductCode string
    // Resource Tags
    Tags string
    // Resource Region
    Region string
    // Official price
    PretaxGrossAmount float64
    // Pretax amount
    PretaxAmount float64
    // Deducated amount
    DeductionAmount float64
    // Instance Bill Extra info
    Extra map[string]interface{}
    model.BaseModel
}

func (m *InstanceBillModel)SetIndex() {
    m.Index = model.GetValFromKeys(m, m.IndexKeys)
}

func (m *InstanceBillModel)SetChecksum() {
    val := model.GetValFromKeys(m, m.ChecksumKeys)
    m.Checksum = model.Checksum(&val)
}

func (m *InstanceBillModel)GetIndex() string {
    return m.Index
}

func (m *InstanceBillModel)GetChecksum() string {
    return m.Checksum
}

func (m *InstanceBillModel)CheckRequired() (bool, string) {
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

func NewInstanceBillModel() *InstanceBillModel {
    m := &InstanceBillModel{}
    m.IndexKeys = "CloudType, AccountId, ProductCode, InstanceId"
    m.ChecksumKeys = "Name, Tags, PretaxGrossAmount, PretaxAmount, DeductionAmount"
    m.required = []string{"BillingCycle", "CloudType", "AccountId", "InstanceId", "PretaxGrossAmount", "PretaxAmount", "DeductionAmount", }

    return m
}
