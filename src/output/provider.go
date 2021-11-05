package output

type Provider interface {
	Init(setting map[string]interface{}) error
	Push(rest []interface{}, query map[string]interface{}) error
}
