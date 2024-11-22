### USAGE
```
.\importer.exe -tp postgres -dsn "host=172.16.26.44 user=postgres password=123456 dbname=postgres port=5432" -tb oss_log2 -f data.csv
```


### NOTE
```
importer version: 1.0.0
 Usage: importer -tp [类型] -dsn [连接串] -tb [表名] -bs [条数] -f [文件名]
  -bs int
        [批次条数] (default 1000)
  -dsn string
        [连接串] gorm连接串
  -f string
        [文件名]
  -tb string
        [表名]
  -tp string
        [数据库类型] 支持mysql, postgres, clickhouse (default "mysql")
```
