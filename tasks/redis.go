package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type redisBroker struct {
	RedisClient *redis.Client
}

type jsonTaskSerializer struct {
	taskName string        `json:"taskName"`
	args     []interface{} `json:"args"`
}

func NewRedisBroker(host string, port uint, password string) Broker {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected: ", pong)
	return &redisBroker{RedisClient: redisClient}
}

func (tm *redisBroker) SetTask(task Task, delay time.Duration, args ...interface{}) {
	key := fmt.Sprintf("%s_%s_%d", TaskNameStartsWith, task.GetName(), time.Now().Unix())
	value, err := tm.serializeTask(task, args)
	if err != nil {
		panic(err)
	}

	set, err := tm.RedisClient.SetNX(key, string(value), delay).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(set)
}

func (tm *redisBroker) serializeTask(task Task, args ...interface{}) ([]byte, error) {
	taskValues := jsonTaskSerializer{
		task.GetName(),
		args,
	}
	return json.Marshal(taskValues)
}
