package functions

import (
	"reflect"
	"strconv"
	"strings"
)
import "fmt"

func NamedFormatString(format string, m map[string]interface{}) string {
	for key, val := range m {
		format = strings.Replace(format, "%{"+key+"}s", fmt.Sprintf("%v", val), -1)
	}
	return format
}

// generic version
func InSlice(v interface{}, l []interface{}) bool {
	if v == nil || l == nil || len(l) == 0 {
		return false
	}
	vKind := reflect.ValueOf(v).Kind()
	if vKind == reflect.Struct {
		return false
	}

	lElemKind := reflect.ValueOf(l[0]).Kind()
	if vKind != lElemKind {
		return false
	}

	for _, elem := range l {
		if v == elem {
			return true
		}
	}

	return false
}

func IntArrToString(ints []int, sep string) string {
	ints_str := ""
	for idx, i := range ints {
		if idx == 0 {
			ints_str += strconv.Itoa(i)
		} else {
			ints_str += sep + strconv.Itoa(i)
		}
	}
	return ints_str
}

func StringToIntArr(s string, sep string) []int {
	int_strs := strings.Split(s, sep)
	ints := []int{}
	for _, int_str := range int_strs {
		i, _ := strconv.Atoi(int_str)
		ints = append(ints, i)
	}
	return ints
}
