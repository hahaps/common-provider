package input

type Resource interface {
	Run(args ...interface{}) ([]interface{}, string, map[string]interface{}, error)
}

type Provider interface {
	GetResourceByName(name string) (Resource, error)
}
