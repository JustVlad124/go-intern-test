package sqlstore

import "github.com/JustVlad124/EWallet/internal/app/models"

type WalletRepository struct {
	stort *Store
}

func (r *WalletRepository) Create(w *models.Wallet) error {
	return r.stort.db.QueryRow(
		"INSERT INTO wallet (balance) VALUES (DEFAULT) RETURNING id",
	).Scan(&w.ID)
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

// need to make more flexible
func (r *WalletRepository) Update(w *models.Wallet, params map[string]interface{}) error {
	return r.stort.db.QueryRow(
		"UPDATE wallet SET balance = $1 WHERE id = $2 RETURNING id",
		params["balance"],
		w.ID,
	).Scan(&w.ID)
}
