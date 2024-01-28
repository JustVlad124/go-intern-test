package store

import "github.com/JustVlad124/EWallet/internal/app/models"

type WalletRepository struct {
	stort *Store
}

func (r *WalletRepository) Create(w *models.Wallet) (*models.Wallet, error) {
	if err := r.stort.db.QueryRow(
		"INSERT INTO wallet (balance) VALUES (DEFAULT) RETURNING id",
	).Scan(&w.ID); err != nil {
		return nil, err
	}

	return w, nil
}

func (r *WalletRepository) FindByID(id string) (*models.Wallet, error) {
	w := &models.Wallet{}

	if err := r.stort.db.QueryRow(
		"SELECT id, balance FROM wallet WHERE id = $1",
		id,
	).Scan(
		&w.ID,
		&w.Balance,
	); err != nil {
		return nil, err
	}

	return w, nil
}
