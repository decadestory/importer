package core

import (
	"encoding/csv"
	"fmt"
	"importer/dbm"
	"strings"
	"time"

	"gitlab.smartee.cn/chenguorui/goutil/exception"
)

// 导入数据
func ImportData(distTable string, batchSize int, hds []string, reader *csv.Reader) error {

	stime := time.Now()
	var cst time.Time
	handleCnt := 0
	records := []map[string]interface{}{}
	for {
		cst = time.Now()
		line, err := reader.Read()
		if err != nil {
			break
		}

		record := map[string]interface{}{}
		for i, v := range hds {
			v = strings.Replace(v, "\ufeff", "", -1)
			record[v] = line[i]
		}
		records = append(records, record)

		if len(records) == batchSize {
			err := dbm.DB.Table(distTable).CreateInBatches(records, batchSize).Error
			exception.Errors.CheckErr(err)

			handleCnt += len(records)
			fmt.Println("handled count:", handleCnt, "batch used time:", time.Since(cst).Milliseconds(), "ms")
			records = []map[string]interface{}{}
		}
	}

	fmt.Println("total handled count:", handleCnt)
	fmt.Println("total time:", time.Since(stime).Seconds(), "s")

	return nil
}
