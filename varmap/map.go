package varmap

import (
	"fmt"
)

func Import(dest map[string]string, src map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range dest {
		result[k] = v
	}
	for k, v := range src {
		key := fmt.Sprintf("var.%s", k)
		result[key] = v
	}
	return result
}
