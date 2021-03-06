package r1Wallet

import (
	"fmt"
	"r1wallet/models"
	"r1wallet/repositories"
)

type R1Wallet struct {
	repository *repositories.Repository
}

func NewR1Wallet(repository *repositories.Repository) *R1Wallet {
	return &R1Wallet{
		repository: repository,
	}
}

func (r *R1Wallet) GetBalance(userID int) (int, error) {
	b, err := r.repository.Wallet.GetBalanceByUserID(userID)
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (r *R1Wallet) Increase(userID, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("negative amount")
	}

	b, err := r.repository.Wallet.GetBalanceByUserID(userID)
	if err != nil {
		return err
	}

	err = r.repository.Wallet.InsertTransaction(models.UserTransactionModel{
		UserID:         userID,
		Amount:         amount,
		CurrentBalance: b + amount,
		Description:    description,
	})

	return err
}

func (r *R1Wallet) Decrease(userID, amount int, description string) error {
	if amount <= 0 {
		return fmt.Errorf("negative amount")
	}

	b, err := r.repository.Wallet.GetBalanceByUserID(userID)
	if err != nil {
		return err
	}

	if amount > b {
		return fmt.Errorf("insufficient credit")
	}

	err = r.repository.Wallet.InsertTransaction(models.UserTransactionModel{
		UserID:         userID,
		Amount:         amount,
		CurrentBalance: b - amount,
		Description:    description,
	})

	return err
}
