package core

import (
	"importer/dbm"
	"strings"
)

// 创建表
func CreateTablePostgresql(distTable string, hds []string) error {

	createSql := "CREATE TABLE IF NOT EXISTS " + distTable + " ( "
	for i, v := range hds {
		v = strings.Replace(v, "\ufeff", "", -1)

		if i == len(hds)-1 {
			createSql += `"` + v + `"` + " text "
		} else {
			createSql += `"` + v + `"` + " text, "
		}
	}
	createSql = createSql + " )"
	res := dbm.DB.Exec(createSql)
	return res.Error
}
