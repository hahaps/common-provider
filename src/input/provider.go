package input


type Credential struct {
	SecretId   string                 `yaml:"secret_id"`
	SecretKey  string                 `yaml:"secret_key"`
	AccountId  string                 `yaml:"account_id"`
	ConnectURL string                 `yaml:"connect_url"`
	Extra      map[string]interface{} `yaml:"extra"`
}

type Resource interface {
	Init(credential Credential) error
	Run(params map[string]interface{}) ([]interface{}, string, map[string]interface{}, error)
}

type Provider interface {
	GetResourceByName(name string) (Resource, error)
}
