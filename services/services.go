package services

import (
	"r1wallet/repositories"
	"r1wallet/services/consumer"
	"r1wallet/services/consumer/redis"
	"r1wallet/services/wallet"
	"r1wallet/services/wallet/r1Wallet"
)

type Services struct {
	Consumer consumer.Consumer
	Wallet   wallet.Wallet
}

func NewServices(repository *repositories.Repository) *Services {
	return &Services{
		Consumer: redis.NewRedisConsumer(repository),
		Wallet:   r1Wallet.NewR1Wallet(repository),
	}
}
