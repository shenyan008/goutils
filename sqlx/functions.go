package sqlx

import (
	"fmt"
	"reflect"
	"strings"
	utils "github.com/shenyan008/goutils/functions"
	"github.com/jmoiron/sqlx"
)

// generic mysql table insert
func TableInsert(DB *sqlx.DB, table string, object interface{}, ignore_fields []string) error {
	var sql_str string
	var str_list []string
	var val_list []interface{}
	var question_list []string

	thisVal := reflect.ValueOf(object) // Value
	thisType := reflect.TypeOf(object) // Type of LiveStream
	for i := 0; i < thisType.NumField(); i++ {
		field := thisType.Field(i)

		ignore_fields_interface := []interface{}{"Id", "CreateTime", "UpdateTime"}
		for _, v := range ignore_fields {
			ignore_fields_interface = append(ignore_fields_interface, v)
		}
		if !utils.InSlice(field.Name, ignore_fields_interface) {

			dbTag := field.Tag.Get("json") // json名同字段名

			str_list = append(str_list, "`"+dbTag+"`")
			question_list = append(question_list, "?")

			fieldType := field.Type
			val := thisVal.FieldByName(field.Name)
			if field.Type.Kind() == reflect.String {
				val_list = append(val_list, val.String())
			} else if strings.Index(fieldType.String(), "int") >= 0 {
				val_list = append(val_list, int(val.Int()))
			}
		}
	}

	sql_str = strings.Join(str_list, ",")
	question_str := strings.Join(question_list, ",")
	fmt.Println(sql_str)
	sql := "insert into `" + table + "` (" + sql_str + ") values (" + question_str + ")"
	_, err := DB.Exec(sql, val_list...)
	return err
}