package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"time"
)

type redisBroker struct {
	RedisClient *redis.Client
}

type jsonTaskSerializer struct {
	TaskName string        `json:"TaskName"`
	Uuid     string        `json:"Uuid"`
	Args     []interface{} `json:"Args"`
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

func (rb *redisBroker) AddTask(task Task, delay time.Duration, args ...interface{}) {
	taskId := uuid.Must(uuid.NewV4()).String()

	value, err := rb.serializeTask(task, taskId, args)
	if err != nil {
		panic(err)
	}
	score := float64(time.Now().UTC().Add(delay).Unix())

	redisTaskItem := redis.Z{Score: score, Member: value}
	_, err = rb.RedisClient.ZAdd(WaitingQueueName, redisTaskItem).Result()
	if err != nil {
		panic(err)
	}
}

func (rb redisBroker) HandleWaitingQueue() {
	go func() {
		for range time.Tick(2 * time.Second) {
			fmt.Println("Here in for")
			tasks, err := rb.RedisClient.ZRangeWithScores(
				WaitingQueueName,
				0, 0,
			).Result()
			fmt.Println(tasks)
			if err != nil {
				fmt.Println("error", err)
			}
			if len(tasks) == 0 {
				continue
			}
			taskPayload := tasks[0]
			if int64(taskPayload.Score) < time.Now().UTC().Unix() {
				rb.runTask(taskPayload.Member.(string))
			}
		}
	}()
	fmt.Println("after go")
}

func (rb *redisBroker) runTask(value string) {
	fmt.Println("TASK running method is UP")
	fmt.Println("Payload: ", value)
	rb.deleteTaskFromQueue(value)
}

func (rb *redisBroker) deleteTaskFromQueue(value string) {
	_, err := rb.RedisClient.ZRem(WaitingQueueName, value).Result()
	if err != nil {
		print("Error while deleting")
	}
}

func (rb *redisBroker) serializeTask(task Task, uuid string, args ...interface{}) ([]byte, error) {
	taskPayload := jsonTaskSerializer{
		task.GetName(),
		uuid,
		args,
	}
	return json.Marshal(taskPayload)
}
