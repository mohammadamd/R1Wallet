package repositories

import "github.com/go-redis/redis/v7"

type redisRepository struct {
	client *redis.Client
}

func (r *redisRepository) SubscribeChannel(channelName string) <-chan *redis.Message {
	p := r.client.Subscribe(channelName)
	return p.Channel()
}
