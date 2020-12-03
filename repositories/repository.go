package repositories

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	"r1wallet/models"
)

type Wallet interface {
	GetBalanceByUserID(ID int) (int, error)
	InsertTransaction(ut models.UserTransactionModel) error
}

type Redis interface {
	SubscribeChannel(channelName string) <-chan *redis.Message
}

type Repository struct {
	Wallet Wallet
	Redis  Redis
}

func NewRepository(db *sql.DB, re *redis.Client) *Repository {
	return &Repository{
		Wallet: &r1WalletRepository{db: db},
		Redis:  &redisRepository{client:re},
	}
}
