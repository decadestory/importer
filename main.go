package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"importer/core"
	"importer/dbm"
	"os"

	"gitlab.smartee.cn/chenguorui/goutil/exception"
)

// var dstType = "clickhouse"
// var distDsn = "clickhouse://default:123456@192.168.100.81:9000/default"
// var distTable = "oss_log"
// var batchSize = 5000

// var dstType = "postgres"
// var distDsn = "host=172.16.26.44 user=postgres password=123456 dbname=postgres port=5432"
// var distTable = "oss_log"
// var batchSize = 1000

// var dstType = "mysql"
// var distDsn = "root:123456@tcp(192.168.100.80:3306)/test"
// var distTable = "oss_log"
// var batchSize = 1000

// 实际中应该用更好的变量名
var (
	dstType   string
	distDsn   string
	distTable string
	batchSize int
	fileName  string
)

func init() {
	flag.StringVar(&dstType, "tp", "mysql", "[数据库类型] 支持mysql, postgres, clickhouse")
	flag.StringVar(&distDsn, "dsn", "", "[连接串] gorm连接串")
	flag.StringVar(&distTable, "tb", "", "[表名]")
	flag.IntVar(&batchSize, "bs", 1000, "[批次条数]")
	flag.StringVar(&fileName, "f", "", "[文件名]")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "importer version: 1.0.0 \n Usage: importer -tp [类型] -dsn [连接串] -tb [表名] -bs [条数] -f [文件名] \n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if distDsn == "" || distTable == "" || fileName == "" {
		flag.Usage()
		return
	}

	fmt.Println("dstType:", dstType)
	fmt.Println("distDsn:", distDsn)
	fmt.Println("distTable:", distTable)
	fmt.Println("batchSize:", batchSize)
	fmt.Println("fileName:", fileName)

	dbm.InitDb(dstType, distDsn)

	file, err := os.Open(fileName)
	exception.Errors.CheckErr(err)
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ','
	hds, err := reader.Read()
	exception.Errors.CheckErr(err)
	fmt.Println("columns:", hds)

	//创建表
	switch dstType {
	case "mysql":
		err = core.CreateTableMysql(distTable, hds)
		exception.Errors.CheckErr(err)
		err = core.ImportData(distTable, batchSize, hds, reader)
		exception.Errors.CheckErr(err)
	case "postgres":
		err = core.CreateTablePostgresql(distTable, hds)
		exception.Errors.CheckErr(err)
		err = core.ImportData(distTable, batchSize, hds, reader)
		exception.Errors.CheckErr(err)
	case "clickhouse":
		err = core.CreateTableClickhouse(distTable, hds)
		exception.Errors.CheckErr(err)
		err = core.ImportData(distTable, batchSize, hds, reader)
		exception.Errors.CheckErr(err)
	default:
		fmt.Println("unsupported dstType")
		return
	}

	fmt.Println("import completed !")
}
