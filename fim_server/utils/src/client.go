package src

import (
	"context"
	"fim_server/utils/stores/logs"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type clientServiceInterface interface {
	Yaml(string, interface{})     // 路径|scan结构体
	Mysql(string) *gorm.DB        // ip:端口|密码|库名	127.0.0.1:3306 password db
	Redis(string) *redis.Client   // ip:端口|密码|库名	127.0.0.1:6379 password 0
	Etcd(string) *clientv3.Client // 服务器地址
	Websocket(w http.ResponseWriter, r *http.Request) *websocket.Conn
}
type clientService struct{}

func Client() clientServiceInterface {
	return &clientService{}
}

// Yaml 读取yaml文件的配置
func (clientService) Yaml(path string, scan interface{}) {
	yamlConf, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Fatal("配置文件加载失败", err.Error())
	}
	err = yaml.Unmarshal(yamlConf, scan)
	if err != nil {
		logs.Fatal("配置文件解析失败", err.Error())
	}
}
func (clientService) Mysql(c string) *gorm.DB {
	var conf = strings.Split(c, " ")
	if len(conf) != 3 {
		logs.Fatal("Mysql内部配置错误: ", conf)
	}
	dsn := fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf[1], conf[0], conf[2])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,                                 // 禁止生成实体外键约束
		Logger:                                   logger.Default.LogMode(logger.Error), // 日志等级
	})
	if err != nil {
		logs.Fatal("MySQL连接失败", c)
		return db
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)                // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)               // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 24) // 最大空闲时间
	
	return db
}
func (clientService) Redis(c string) *redis.Client {
	var conf = strings.Split(c, " ")
	if len(conf) != 3 {
		logs.Fatal("Redis内部配置错误: ", conf)
	}
	db, _ := strconv.Atoi(conf[2])
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf[0],
		Password: conf[1], // no password set
		DB:       db,      // use default DB
		PoolSize: 100,     // 连接池大小
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
