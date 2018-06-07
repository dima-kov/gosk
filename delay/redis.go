package delay

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"time"
)

type redisBroker struct {
	redisClient *redis.Client
	queueName   string
}

type jsonTaskSerializer struct {
	TaskName string        `json:"TaskName"`
	Uuid     string        `json:"Uuid"`
	Args     []interface{} `json:"Args"`
}

func NewRedisBroker(host string, port uint, password, queueName string) *redisBroker {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})
	broker := redisBroker{redisClient: redisClient, queueName: queueName}
	broker.checkConnection()
	return &broker
}

func (rb *redisBroker) AddTask(task DelayTask, delay time.Duration, args ...interface{}) (int64, error) {
	taskId := uuid.Must(uuid.NewV4()).String()

	score := float64(time.Now().UTC().Add(delay).Unix())
	serializedTask, err := rb.serializeTask(task, taskId, args)
	if err != nil {
		panic(err)
	}

	redisTaskItem := redis.Z{Score: score, Member: serializedTask}
	return rb.redisClient.ZAdd(rb.queueName, redisTaskItem).Result()
}

func (rb redisBroker) HandleWaitingQueue() {
	go func() {
		for range time.Tick(2 * time.Second) {
			tasks, err := rb.redisClient.ZRangeWithScores(rb.queueName, 0, 0).Result()
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
}

func (rb *redisBroker) runTask(value string) {
	fmt.Println("TASK running method is UP")
	fmt.Println("Payload: ", value)
	rb.deleteTaskFromQueue(value)
}

func (rb *redisBroker) deleteTaskFromQueue(value string) {
	fmt.Println("Delete task from queue: ", value)
	_, err := rb.redisClient.ZRem(rb.queueName, value).Result()
	if err != nil {
		print("Error while deleting")
	}
}

func (rb *redisBroker) checkConnection() {
	pong, err := rb.redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis connected: PING -", pong)
}

func (rb *redisBroker) serializeTask(task DelayTask, uuid string, args ...interface{}) ([]byte, error) {
	taskPayload := jsonTaskSerializer{
		task.GetName(),
		uuid,
		args,
	}
	return json.Marshal(taskPayload)
}
