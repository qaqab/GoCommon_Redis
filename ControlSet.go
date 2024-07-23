package GoCommon_Redis

import (
	"fmt"
	"time"
)

func (redisSetting *RedisSetting) R_sadd(key string, data string) bool {
	// 使用Redis客户端执行SAdd命令，将data添加到key对应的集合中
	err := redisSetting.RedisClient.SAdd(key, data).Err()
	if err != nil {
		// 如果发生错误，打印错误信息
		fmt.Printf("数据写入错误：%v", err)
		// 返回false表示数据插入失败
		return false
	}
	// 打印数据插入成功的消息
	fmt.Println("数据插入成功")
	// 返回true表示数据插入成功
	return true
}

func (redisSetting *RedisSetting) R_scard(key string) int {
	// 休眠1秒
	time.Sleep(1 * time.Second)

	// 调用RedisClient的SCard方法查询key对应集合的大小，并将结果转换为int类型
	size, err := redisSetting.RedisClient.SCard(key).Result()
	if err != nil {
		// 如果查询出错，则打印错误信息并返回-1
		fmt.Printf("数据查询错误错误：%v", err)
		return -1
	}

	// 返回集合的大小
	return int(size)
}

func (redisSetting *RedisSetting) R_spop(key string) string {
	// 从 Redis 中弹出一个元素，并返回该元素的值
	val, _ := redisSetting.RedisClient.SPop(key).Result()

	// 打印弹出元素的值
	fmt.Println(val)

	// 返回弹出元素的值
	return val
}

func (redisSetting *RedisSetting) R_spopn(key string, count int64) []string {

	// 从 Redis 中使用 SPopN 方法获取指定 key 中 count 个元素，并返回结果
	vals, _ := redisSetting.RedisClient.SPopN(key, count).Result()

	// 打印获取到的元素列表
	fmt.Println(vals)

	// 返回获取到的元素列表
	return vals
}
