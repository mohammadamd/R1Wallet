package redis

import (
	"fmt"
	"r1wallet/repositories"
)

type redisConsumer struct {
	repository *repositories.Repository
}

func NewRedisConsumer(repository *repositories.Repository) *redisConsumer {
	return &redisConsumer{
		repository: repository,
	}
}

func (r *redisConsumer) Consume(messages chan<- string, channelName string) {
	ch := r.repository.Redis.SubscribeChannel(channelName)
	for m := range ch {
		fmt.Println("new Message received: ", m)
		fmt.Println("message payload: ", m.Payload)
		messages <- m.Payload
	}
}
