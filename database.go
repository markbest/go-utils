package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type DB struct {
	Client *sql.DB
}

//New database client
func NewDB(host string, port int, user string, password string, database string) (*DB, error) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+
		database+"?charset=utf8")
	if err != nil {
		return nil, err
	}
	client := &DB{Client: db}
	return client, nil
}

//Parse request params
func parseRequestParams(action string, conditions map[string]interface{}, values map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	switch action {
	case "insert":
		var column, temp []string
		var value []interface{}
		if len(values) > 0 {
			for k, v := range values {
				temp = append(temp, "?")
				column = append(column, k)
				value = append(value, v)
			}
		}
		params["column"] = strings.Join(column, ", ")
		params["temp"] = strings.Join(temp, ", ")
		params["value"] = value
		return params
	case "update":
		var updateColumn, updateCondition []string
		var value []interface{}
		if len(values) > 0 {
			for k, v := range values {
				updateColumn = append(updateColumn, k+" = ?")
				value = append(value, v)
			}
		}
		if len(conditions) > 0 {
			index := 1
			for k, v := range conditions {
				updateCondition = append(updateCondition, k+" = ?")
				value = append(value, v)
				index++
			}
		}
		params["condition"] = strings.Join(updateCondition, " AND ")
		params["column"] = strings.Join(updateColumn, " AND ")
		params["value"] = value
		return params
	case "delete":
		var condition []string
		var value []interface{}
		if len(conditions) > 0 {
			for k, v := range conditions {
				condition = append(condition, k+" = ?")
				value = append(value, v)
			}
		}
		params["condition"] = strings.Join(condition, " AND ")
		params["value"] = value
		return params
	case "select":
		var condition []string
		var value []interface{}
		if len(conditions) > 0 {
			for k, v := range conditions {
				condition = append(condition, k+" = ?")
				value = append(value, v)
			}
		}
		params["condition"] = strings.Join(condition, " AND ")
		params["value"] = value
		return params
	default:
	}
	return params
}

//Insert action
func (d *DB) Insert(table string, values map[string]interface{}) (sql.Result, error) {
	params := parseRequestParams("insert", nil, values)
	query := "INSERT INTO " + table + "(" + params["column"].(string) + ")" + " VALUES (" + params["temp"].(string) + ")"
	result, err := d.Client.Exec(query, params["value"].([]interface{})...)
	return result, err
}

//Update action
func (d *DB) Update(table string, conditions map[string]interface{}, values map[string]interface{}) (sql.Result, error) {
	params := parseRequestParams("update", conditions, values)
	query := "UPDATE " + table + " set " + params["condition"].(string) + " where " + params["column"].(string)
	result, err := d.Client.Exec(query, params["value"].([]interface{})...)
	return result, err
}

//Delete action
func (d *DB) Delete(table string, conditions map[string]interface{}) (sql.Result, error) {
	params := parseRequestParams("delete", conditions, nil)
	query := "DELETE FROM " + table + " where " + params["condition"].(string)
	result, err := d.Client.Exec(query, params["value"].([]interface{})...)
	return result, err
}

//Select action
func (d *DB) Select(table string, conditions map[string]interface{}) ([]interface{}, error) {
	params := parseRequestParams("select", conditions, nil)
	query := "SELECT * FROM " + table + " where " + params["condition"].(string)
	rows, err := d.Client.Query(query, params["value"].([]interface{})...)
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := make([]interface{}, 0)
	for rows.Next() {
		temp := make(map[string]interface{})
		if err := rows.Scan(scanArgs...); err != nil {
			panic(err)
		}

		for k, v := range columns {
			temp[v] = string(values[k])
		}
		result = append(result, temp)
	}
	return result, err
}

//Close database
func (d *DB) Close() {
	d.Client.Close()
}
