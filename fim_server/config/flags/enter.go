package flags

import (
	"fim_server/utils/stores/logs"
	"flag"
)

// Command 根据命令执行不同的函数
func Command() {
	DB := flag.Bool("db", false, "初始化数据库")
	flag.Parse() // 解析命令行参数写入注册的flag里

	if *DB {
		logs.Fatal("->生成数据库表结构", MigrationTable() == nil)
	}
}
