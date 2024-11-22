package core

import (
	"importer/dbm"
	"strings"
)

// 创建表
func CreateTableClickhouse(distTable string, hds []string) error {

	createSql := "CREATE TABLE IF NOT EXISTS " + distTable + " ( "
	for _, v := range hds {
		v = strings.Replace(v, "\ufeff", "", -1)
		createSql += v + " String, "
	}
	createSql = createSql + " ) ENGINE = MergeTree() ORDER BY " + hds[0]
	res := dbm.DB.Exec(createSql)
	return res.Error
}
