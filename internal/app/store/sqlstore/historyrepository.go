package sqlstore

import (
	"github.com/JustVlad124/EWallet/internal/app/models"
)

type HistoryRepository struct {
	store *Store
}

func (r *HistoryRepository) TransferTo(h *models.History) error {
	return r.store.db.QueryRow(
		"INSERT INTO history (time, \"from\", \"to\", amount) VALUES ($1, $2, $3, $4) RETURNING id",
		h.Time,
		h.From,
		h.To,
		h.Amount,
	).Scan(&h.ID)
}

func (r *HistoryRepository) FindAllTransferByID(id string) ([]models.History, error) {
	rows, err := r.store.db.Query(
		"SELECT time, \"from\", \"to\", amount FROM history WHERE \"from\" = $1 OR \"to\" = $2",
		id,
		id,
	)
	if err != nil {
		return nil, err
	}
	transers := []models.History{}
	defer rows.Close()
	for rows.Next() {
		transer := models.History{}
		err = rows.Scan(&transer.Time, &transer.From, &transer.To, &transer.Amount)
		if err != nil {
			return nil, err
		}
		transers = append(transers, transer)
	}

	return transers, nil
}
