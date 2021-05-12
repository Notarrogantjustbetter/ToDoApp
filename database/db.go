package database

import (
	"github.com/go-redis/redis"
)

var client *redis.Client


func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}


func CreateTask(task string) error {
	return client.LPush("Tasks", task).Err()
}

func DeleteTask(task string) error {
	return client.LRem("Tasks", 1, task).Err()
}

func GetTasks()([]string, error) {
	return client.LRange("Tasks", 0, 100).Result()
}


