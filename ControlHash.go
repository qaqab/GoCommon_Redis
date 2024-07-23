package GoCommon_Redis

import (
	"fmt"
)

func (redisSetting *RedisSetting) R_hset(key string, h_key string, h_value string) bool {
	// 调用 Redis 客户端的 HSet 方法，将 h_key 和 h_value 存入指定的 key 对应的哈希表中
	ok, err := redisSetting.RedisClient.HSet(key, h_key, h_value).Result()

	// 打印操作是否成功的标志
	fmt.Println(ok)
	// 打印操作过程中出现的错误
	fmt.Println(err)

	// 返回操作是否成功的标志
	return ok
}
func (redisSetting *RedisSetting) R_hget(key string, h_key string) string {
	// 调用 RedisClient 的 HGet 方法，获取指定哈希表中的字段值
	value, err := redisSetting.RedisClient.HGet(key, h_key).Result()
	// 打印获取到的字段值
	fmt.Println(value)
	// 打印错误信息（如果有）
	fmt.Println(err)
	// 返回获取到的字段值
	return value
}

func (redisSetting *RedisSetting) R_hgetall(key string) {
	// 从 Redis 中获取指定 key 的所有字段和值
	data, err := redisSetting.RedisClient.HGetAll(key).Result()
	if err != nil {
		// 如果获取失败，抛出异常
		panic(err)
	}
	// 遍历字段和值，并打印到控制台
	for field, val := range data {
		fmt.Println(field, val)
	}
}
