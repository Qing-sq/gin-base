package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)
var rdb *redis.Client

func initClient()(err error)  {
	rdb = redis.NewClient(&redis.Options{
		Addr:fmt.Sprintf("%s:%d",viper.GetString("host"),viper.GetInt("port")),
		Password: "",
		DB: viper.GetInt("db"),	//use default db
		PoolSize: viper.GetInt("poolSize"), // redis连接池大小

	})
	_,err = rdb.Ping().Result()
	if err != nil{
		return err
	}
	return nil
}

func Main()(err error)  {
	if err := initClient();err != nil{
		fmt.Println("init redis fail")
		return err
	}
	//程序退出释放相关资源
	defer rdb.Close()
	return
}