package internal

import (
	"context"
	"errors"

	"github.com/soonio/pupil/app"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	config := app.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(errors.New("redis connect ping failed, err:" + err.Error()))
	} else {
		app.Redis = client
	}
}
