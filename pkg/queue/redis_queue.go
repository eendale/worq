package queue

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/eendale/worq/pkg/job"
)

type RedisQueue struct {
	client *redis.Client
	name   string
}

func NewRedisQueue(client *redis.Client, name string) *RedisQueue {
	return &RedisQueue{client: client, name: name}
}

func (rq *RedisQueue) Enqueue(j job.Job) error {
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}
	return rq.client.RPush(context.Background(), rq.name, data).Err()
}

func (rq *RedisQueue) Dequeue() (job.Job, error) {
	data, err := rq.client.LPop(context.Background(), rq.name).Result()
	if err != nil {
		return job.Job{}, err
	}
	var j job.Job
	err = json.Unmarshal([]byte(data), &j)
	return j, err
}
