package GoCommon_Redis

import (
	"testing"

	"github.com/qaqab/GoCommon_DbManager"

	"github.com/qaqab/GoCommon_Redis"
)

func TestRedisHSet(t *testing.T) {
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}
	clientAll.DbManagerClient("redis.task")
	redis_client := clientAll.RedisClient

	redisSetting := GoCommon_Redis.RedisSetting{RedisClient: redis_client, Password: clientAll.RedisSettingData.Password, Addresse: clientAll.RedisSettingData.Addresse, DB: clientAll.RedisSettingData.DB}
	redisSetting.R_hset("sad", "sadasdas", "asdasdasd")
}
