package xlog

// Fields represents attached fileds of log
type Fields map[string]interface{}

// Merge multi fileds into new Fields instance
func NewFields(fields ...Fields) Fields {
	result := Fields{}
	for _, item := range fields {
		if item == nil {
			continue
		}
		for k, v := range item {
			result[k] = v
		}
	}
	return result
}
