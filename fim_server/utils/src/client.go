package src

import (
	"context"
	"fim_server/utils/stores/logs"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type clientServiceInterface interface {
	Mysql(string) *gorm.DB        // root:password@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4&parseTime=True&loc=Local
	Redis(string) *redis.Client   // ip:端口|密码|第几个数据库|连接池大小	 127.0.0.1:80 password 0 100
	Etcd(string) *clientv3.Client // 服务器地址
	Websocket(w http.ResponseWriter, r *http.Request) *websocket.Conn
}
type clientService struct{}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Client() clientServiceInterface {
	return &clientService{}
}
func (clientService) Mysql(c string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // 日志等级
	})
	if err != nil {
		logs.Fatal("MySQL连接失败", c)
		return db
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	
	return db
}
func (clientService) Redis(c string) *redis.Client {
	var opt = strings.Split(c, " ")
	if len(opt) != 4 {
		logs.Fatal("Redis内部配置错误: ", opt)
	}
	db, _ := strconv.Atoi(opt[2])
	size, _ := strconv.Atoi(opt[3])
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt[0],
		Password: opt[1], // no password set
		DB:       db,     // use default DB
		PoolSize: size,   // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logs.Fatal("Redis连接失败: ", err.Error())
	}
	return rdb
}
func (clientService) Etcd(c string) *clientv3.Client {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{c}, // etcd服务器地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Fatal(err)
	}
	return etcd
}
func (clientService) Websocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 鉴权   true 放行 | false 拦截
		},
	}
	// http upgrade websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Fatal("websocket升级失败 ->", err)
		return nil
	}
	return conn
}
