// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gdb

import (
	"strings"

	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/encoding/gbinary"

	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
)

// convertValue automatically checks and converts field value from database type
// to golang variable type.
func (c *Core) convertValue(fieldValue interface{}, fieldType string) interface{} {
	// If there's no type retrieved, it returns the <fieldValue> directly
	// to use its original data type, as <fieldValue> is type of interface{}.
	if fieldType == "" {
		return fieldValue
	}
	t, _ := gregex.ReplaceString(`\(.+\)`, "", fieldType)
	t = strings.ToLower(t)
	switch t {
	case
		"binary",
		"varbinary",
		"blob",
		"tinyblob",
		"mediumblob",
		"longblob":
		return gconv.Bytes(fieldValue)

	case
		"int",
		"tinyint",
		"small_int",
		"smallint",
		"medium_int",
		"mediumint",
		"serial":
		if gstr.ContainsI(fieldType, "unsigned") {
			gconv.Uint(gconv.String(fieldValue))
		}
		return gconv.Int(gconv.String(fieldValue))

	case
		"big_int",
		"bigint",
		"bigserial":
		if gstr.ContainsI(fieldType, "unsigned") {
			gconv.Uint64(gconv.String(fieldValue))
		}
		return gconv.Int64(gconv.String(fieldValue))

	case "real":
		return gconv.Float32(gconv.String(fieldValue))

	case
		"float",
		"double",
		"decimal",
		"money",
		"numeric",
		"smallmoney":
		return gconv.Float64(gconv.String(fieldValue))

	case "bit":
		s := gconv.String(fieldValue)
		// mssql is true|false string.
		if strings.EqualFold(s, "true") {
			return 1
		}
		if strings.EqualFold(s, "false") {
			return 0
		}
		return gbinary.BeDecodeToInt64(gconv.Bytes(fieldValue))

	case "bool":
		return gconv.Bool(fieldValue)

	case "date":
		t, _ := gtime.StrToTime(gconv.String(fieldValue))
		return t.Format("Y-m-d")

	case
		"datetime",
		"timestamp":
		t, _ := gtime.StrToTime(gconv.String(fieldValue))
		return t.String()

	default:
		// Auto detect field type, using key match.
		switch {
		case strings.Contains(t, "text") || strings.Contains(t, "char") || strings.Contains(t, "character"):
			return gconv.String(fieldValue)

		case strings.Contains(t, "float") || strings.Contains(t, "double") || strings.Contains(t, "numeric"):
			return gconv.Float64(gconv.String(fieldValue))

		case strings.Contains(t, "bool"):
			return gconv.Bool(gconv.String(fieldValue))

		case strings.Contains(t, "binary") || strings.Contains(t, "blob"):
			return fieldValue

		case strings.Contains(t, "int"):
			return gconv.Int(gconv.String(fieldValue))

		case strings.Contains(t, "time"):
			s := gconv.String(fieldValue)
			t, err := gtime.StrToTime(s)
			if err != nil {
				return s
			}
			return t.String()

		case strings.Contains(t, "date"):
			s := gconv.String(fieldValue)
			t, err := gtime.StrToTime(s)
			if err != nil {
				return s
			}
			return t.Format("Y-m-d")

		default:
			return gconv.String(fieldValue)
		}
	}
}

// filterFields removes all key-value pairs which are not the field of given table.
func (c *Core) filterFields(schema, table string, data map[string]interface{}) map[string]interface{} {
	// It must use data copy here to avoid its changing the origin data map.
	newDataMap := make(map[string]interface{}, len(data))
	if fields, err := c.DB.TableFields(table, schema); err == nil {
		for k, v := range data {
			if _, ok := fields[k]; ok {
				newDataMap[k] = v
			}
		}
	}
	return newDataMap
}
