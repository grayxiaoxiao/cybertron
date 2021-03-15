package tmpl

type StructField struct {
	FieldName string
	FieldType string
	FieldTags string
}

var StructTMPL = `
package {{.packageName}}

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type {{.cybertronName}} struct {
	{{- range .structFields }}
	  {{.FieldName }} {{.FieldType }} {{.FieldTags }}
	{{- end}}
}

func (obj {{.cybertronName}}) Attributes() (attributes []string) {
  attr := reflect.TypeOf(obj)
  num  := attr.NumField()
  for i := 0; i < num; i++ {
      column_name := attr.Field(i).Tag.Get("column_name")
      attributes   = append(attributes, column_name)
  }
  return attributes
}

func (obj {{.cybertronName}}) Insert() (obj_id int64, err error) {
	db         := *sql.DB //mysql_driver.Conn(), here get your sql.DB 
	toJson, _  := obj.ToJson()
	attributes := obj.Attributes()
	columns    := strings.Join(attributes, ", ")
	values     := make([]string, len(attributes))
	for i := 0; i < len(attributes); i++ {
		values = append(values, fmt.Sprintf("%v", toJson[attributes[i]]))
	}
	insertSql := "INSERT INTO {{.tableName}}(" + columns + ") VALUES(" + strings.Join(values, ", ") + ")"
	rows, err := db.QueryContext(context.Background(), insertSql)
	defer rows.Close()
	if err != nil {
		logs.GetLogger("INSERTORDERERROR").Println(err.Error())
	}
	rows.Scan(&obj_id)
	return obj_id, err
}

func (obj {{.cybertronName}}) Destroy() (valid bool, err error) {
	db         := *sql.DB //mysql_driver.Conn(), here get your sql.DB
	destroySql := "DELETE FROM {{.tableName}} WHERE id = ?"
	db.QueryRowContext(context.Background(), destroySql, obj.Id)
	return valid, err
}

func (obj {{.cybertronName}}) ToJson() (res map[string]interface, err error) {
	inrec, err := json.Marshal(obj)
	if err != nil {
		return res, err
	}
	json.Unmarshal(inrec, &res)
	return res, err
}

func (obj {{.cybertronName}}) Update(params map[string]interface{}) (obj_id int64, err error) {
	updateSets := make([]string, len(params))
	for k, v   := range params {
		updateSets = append(updateSets, fmt.Sprintf("%s = '%v'", k , v))
	}
	updateSql := "UPDATE {{.tableName}} SET " + strings.Join(updateSets, ", ") + " WHERE id = ?"
	db        := mysql_driver.Conn()
	_, err     = db.QueryContext(context.Background(), updateSql, Obj.Id)
	return obj.Id, err
}

`