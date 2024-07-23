package GoCommon_Redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func (redisSetting *RedisSetting) R_lpush(key string, l_value string) int {
	// 调用Redis客户端的LPush方法，将l_value插入到key对应的列表的左侧，并获取列表的新长度
	size, err := redisSetting.RedisClient.LPush(key, l_value).Result()
	if err != nil {
		// 如果插入失败，打印错误信息并返回-1
		fmt.Printf("插入失败%v", err)
		return -1
	}
	// 返回列表的新长度
	return int(size)
}

func (redisSetting *RedisSetting) R_lpop(client *redis.Client, key string) string {
	// 从Redis中取出列表最左边的元素
	data, err := redisSetting.RedisClient.LPop(key).Result()
	if err != nil {
		fmt.Printf("插入失败%v", err)
		return ""
	}
	fmt.Printf("取出的数据为%v", data)
	return data
}
