package store

import "github.com/JustVlad124/EWallet/internal/app/models"

type WalletRepository interface {
	Create(*models.Wallet) error
	FindByID(string) (*models.Wallet, error)
	Update(*models.Wallet, map[string]interface{}) error
}

type HistoryRepository interface {
	TransferTo(*models.History) error
	FindAllTransferByID(string) ([]models.History, error)
}
