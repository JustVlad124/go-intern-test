package store

import "github.com/JustVlad124/EWallet/internal/app/models"

type HistoryRepository struct {
	store *Store
}

func (r *HistoryRepository) Create(h *models.History) (*models.History, error) {
	return nil, nil
}

func (r *HistoryRepository) TransferTo(h *models.History) (*models.History, error) {
	return nil, nil
}

func (r *HistoryRepository) FindAllTransferByID(id string) (*models.History, error) {
	return nil, nil
}
