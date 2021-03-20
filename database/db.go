package database

import (
	"github.com/go-redis/redis"
)

type Task struct {
	Name string
}

type TaskInterface interface {
	CreateTask(task string) error
	GetTasks() ([]string, error)
	DeleteTask(task string) error
}

func InitDb() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func (t Task) CreateTask(task string) error {
	redis := InitDb()
	return redis.LPush(redis.Context(), "Tasks", task).Err()
}

func (t Task) DeleteTask(task string) error {
	redis := InitDb()
	return redis.LRem(redis.Context(), "Tasks", 1, task).Err()
}

func (t Task) GetTasks() ([]string, error) {
	redis := InitDb()
	return redis.LRange(redis.Context(), "Tasks", 0, 100).Result()
}

