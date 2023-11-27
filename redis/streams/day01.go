package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// 研究一下stream 做消息队列 这个数据结构
// 按照那种简易的写法，来完成这个过程。

var (
	streamKey = "my_stream"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func consumer(client *redis.Client, ctx context.Context, groupName string, consumerName string, streamkey string) {

	for {
		result, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamKey, ">"},
			Block:    0,
			Count:    1,
		}).Result()
		if err != nil {

		}
		for _, message := range result {
			for _, x := range message.Messages {
				for _, msg := range x.Values {
					log.Printf("consumer %s\n", msg)
				}
				client.XAck(ctx, streamkey, groupName, x.ID)
			}

		}
	}
}

func producer(client *redis.Client, ctx context.Context, streamkey string) {
	client.XAdd(ctx, &redis.XAddArgs{
		Stream: streamkey,
		Values: "我和你都约好了",
	})
}

func main() {
	client := NewRedisClient()
	ctx := context.Background()
	groupName := "mygroup"
	consumerName := "myconsumer"
	go consumer(client, ctx, groupName, consumerName, streamKey)
	go producer(client, ctx, streamKey)
	time.Sleep(5 * time.Second)

}
