package GoCommon_Redis

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_DbManager"
	"github.com/qaqab/GoCommon_Redis"
)

func TestRedisClient(t *testing.T) {

	// 创建一个 DbManager.ClientAll 实例
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}

	// 调用 DbManager.ClientAll 的 DbManagerClient 方法，参数为 "redis"
	clientAll.DbManagerClient("redis.task")

	// 获取 Redis 客户端实例
	redis_Client := clientAll.RedisClient

	// 创建一个 ControlRedis.RedisSetting 实例，并使用 clientAll.RedisSettingData 的值进行初始化
	redis_Setting := GoCommon_Redis.RedisSetting{Addresse: clientAll.RedisSettingData.Addresse, Password: clientAll.RedisSettingData.Password, DB: clientAll.RedisSettingData.DB}

	// 将 redis_Client 赋值给 redis_Setting 的 RedisClient 字段
	redis_Setting.RedisClient = redis_Client

	// 打印 redis_Client 的值
	fmt.Println(redis_Client)
}
