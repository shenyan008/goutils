package goutils

import "strings"
import "fmt"

func NamedFormatString(format string, m map[string]interface{}) string {
	for key, val := range m {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%v", val), -1)
	}
	return format
}
